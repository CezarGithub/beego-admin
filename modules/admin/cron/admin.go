package cron

import (
	"quince/initialize/scheduler"

	"github.com/beego/beego/v2/core/logs"
)

func init() {
	job := scheduler.NewCronJob()
	job.Name = "demo_cron_job"
	job.Description = "Admin demo cron job"
	job.Module = "admin"
	job.Func = theJob
	scheduler.AddCronJob(job)

}
func theJob() error {
	logs.Info("Admin demo job handled")
	return nil
}
