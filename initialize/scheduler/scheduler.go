package scheduler

import (
	"quince/internal/tasks"
)

var scheduler *tasks.Scheduler
var jobs []*cronJob

// func add() {

// 	id, err := scheduler.Add(&tasks.Task{
// 		Interval: time.Duration(10 * time.Second), // run every 1 hour
// 		RunOnce:  false,
// 		TaskFunc: activateCronJobs,
// 	})
// 	logs.Info(" Task manager scheduler : %s ", id)
// 	if err != nil {
// 		logs.Info("Task manager scheduler error...", err.Error)
// 	}
// }

func NewSchedulerManager() *tasks.Scheduler {
	scheduler = tasks.New()
	cronManager() //start cron manager
	return scheduler
}
func Instance() *tasks.Scheduler {
	return scheduler
}
func NewCronJob() *cronJob {
	c := new(cronJob)
	return c
}
func AddCronJob(j *cronJob) {
	jobs = append(jobs, j)
}
func CronJobs() []*cronJob {
	return jobs
}
