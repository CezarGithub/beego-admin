package transaction

import (
	"context"
	"encoding/json"
	"quince/global"
	"quince/internal/models"
	"quince/utils/encrypter"
	"sync"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Queue holds name, list of jobs and context with cancel.
type queue struct {
	name   string
	user   int64
	jobs   chan Operation
	bunch  []Operation
	ctx    context.Context
	cancel context.CancelFunc
}

// Operation - interface for all transactionable methods
type Operation interface {
	Run(txOrm orm.TxOrmer) error
	Description() string
	GetModel() models.IModel
}

// Worker responsible for queue serving.
type worker struct {
	Queue *queue
}

// NewQueue instantiates new queue.
// func NewQueue(name string, user int64) *queue {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	return &queue{
// 		jobs:   make(chan Operation),
// 		name:   name,
// 		ctx:    ctx,
// 		cancel: cancel,
// 		user:   user,
// 	}
// }
// func (q *queue) AddOperation(b Operation) {
// 	b.SetUser(q.user)
// 	q.bunch = append(q.bunch, b)
// }

// ready - prepare transaction - move operations from list to chan
func (q *queue) ready() {
	q.addTransactions()
}

// addTransactions adds jobs to the queue and cancels channel.
// func (q *Queue) addTransactions(jobs []*Transaction) {
func (q *queue) addTransactions() {
	var wg sync.WaitGroup
	jobs := q.bunch
	wg.Add(len(jobs))

	//for _, job := range jobs {
	for _, job := range q.bunch {
		//job := b.(Operation)
		go func(job Operation) {
			q.jobs <- job
			wg.Done()
		}(job)
	}

	go func() {
		wg.Wait()
		q.cancel()
	}()

}

// NewWorker initialises new Worker.
func newWorker(queue *queue) *worker {
	queue.ready()
	return &worker{
		Queue: queue,
	}
}

// DoWork processes jobs from the queue (jobs channel).
func (w *worker) doWork(txOrm orm.TxOrmer) error {
	for {
		select {
		// if context was canceled.
		case <-w.Queue.ctx.Done():
			logs.Info("WORK DONE in queue %s: %s!", w.Queue.name, w.Queue.ctx.Err())
			return nil
		// if job received.
		case job := <-w.Queue.jobs:
			err := job.Run(txOrm)
			model := job.GetModel()
			jn, _ := json.Marshal(model) //save model to database logs
			cryptData := encrypter.Encrypt(jn, []byte(global.BA_CONFIG.Other.LogAesKey))
			if err != nil {
				logs.Error(err.Error())
				txOrm.Rollback()

				//continue
				return err
			}
			sql := "INSERT INTO admin_log_tx (table_name,table_row_id,data,title,create_user,create_time,update_user,update_time) VALUES (?,?,?,?,?,?,?,?)"
			_, err = txOrm.Raw(sql, model.TableName(), model.GetID(), cryptData, job.Description(), w.Queue.user, time.Now(), w.Queue.user, time.Now()).Exec()
			if err != nil {
				logs.Error(err.Error())
				txOrm.Rollback()

				//continue
				return err
			}
		}
	}
}
