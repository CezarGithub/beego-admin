package cron

import (
	"quince/initialize/scheduler"

	"github.com/beego/beego/v2/core/logs"
)

func init() {
	job := scheduler.NewCronJob()
	job.Name = "bnr_exchange_rates"
	job.Description = "BNR curs valutar"
	job.Module = "master"
	job.Func = theJob

	scheduler.AddCronJob(job)

}
func theJob() error {
	logs.Info("BNR cron job ")
	return nil
}
