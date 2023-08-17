package scheduler

import (
	"quince/internal/tasks"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func cronManager() {
	job := NewCronJob()
	job.Name = "cron_jobs_manager"
	job.Description = "Application cron jobs manager"
	job.Module = "admin"
	job.Func = activateCronJobs
	AddCronJob(job)

	id, err := Instance().Add(&tasks.Task{
		Interval: time.Duration(30 * time.Minute), // run every 1/2 hour
		RunOnce:  false,
		TaskFunc: activateCronJobs,
	})
	logs.Info(" Task manager scheduler started : %s ", id)
	if err != nil {
		logs.Info("Task manager scheduler error...", err.Error)
	}

}

// Activate scheduled jobs for today
func activateCronJobs() error {
	var err error
	t := time.Now()
	today := Interval{}
	today.DayOfTheMonth = t.Day()
	today.DayOfWeek = int(t.Weekday())
	today.MonthOfTheYear = int(t.Month())
	today.Hour = t.Hour()
	today.Minute = t.Minute()
	logs.Info("Cron manager run  : %+v", today)

	// for _, item := range jobs {
	// 	_, err = scheduler.Instance().Add(item.Task)
	// 	if err != nil {
	// 		logs.Error("Task scheduler error :%s %s", item.Name, err.Error())
	// 	}
	// }
	return err
}
