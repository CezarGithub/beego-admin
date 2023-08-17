package transaction

import (
	"context"
	"github.com/beego/beego/v2/client/orm"
)

type Batch struct {
	List *queue
}

// NewQueue instantiates new queue.
func NewTransaction(name string, user int64) *Batch {
	ctx, cancel := context.WithCancel(context.Background())
	t := new(Batch)
	t.List = &queue{
		jobs:   make(chan Operation),
		name:   name,
		ctx:    ctx,
		cancel: cancel,
		user:   user,
	}
	return t
}

// Add operation in transaction list
func (b *Batch) Add(o Operation) {
	o.GetModel().SetUser(b.List.user)
	b.List.bunch = append(b.List.bunch, o)
}

// Execute tranaction
func (b *Batch) Execute() error {
	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

		worker := newWorker(b.List)
		e := worker.doWork(txOrm)
		if e!=nil {
			return e;//errors.New("error.failed")
		}
		return nil
	})
	return err
}
