package main

import (
	"fmt"
	_ "quince/initialize/conf"
	"quince/initialize/database"
	_ "quince/initialize/funcmap"
	_ "quince/initialize/i18n"
	"quince/initialize/module"
	"quince/initialize/router"
	"quince/initialize/scheduler"
	_ "quince/initialize/session"
	_ "quince/initialize/xsrf"
	"quince/middleware"
	_ "quince/modules"

	e "quince/internal/error"
	"quince/internal/i18n"

	_ "quince/utils/template"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const (
	APP_VER = "0.1.1.0227"
)

func main() {
	//beeog.Info(beego.AppName, APP_VER)
	//default beego system logs ; app logs moved to base controller
	f, _ := web.AppConfig.String("log::system")
	logs.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(`{"filename":"log/%s","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true,"rotate":true}`, f))
	web.BConfig.WebConfig.ViewsPath, _ = web.AppConfig.String("modulesFolder")

	web.BConfig.RecoverPanic = true
	web.BConfig.RecoverFunc = middleware.RecoverPanic
	
	//update database
	database.DbSync()
	router.DBSync()
	module.DBSync()

	//load languages files
	languages()
	//tasks scheduler
	sheduler := scheduler.NewSchedulerManager()
	defer sheduler.Stop()
	logs.Info("Total registered cron jobs : %d", len(scheduler.CronJobs()))
	logs.Info("Total active cron jobs : %d", len(sheduler.Tasks()))
	logs.Info("System started")
	logs.Info("TODO -> recovery panic screen with home button")
	//error handler
	web.ErrorHandler("404", e.PageNotFound)
	web.ErrorHandler("500", e.ServerError)
	web.ErrorHandler("403", e.AccessDenied)

	//debug sql
	runMode, _ := web.AppConfig.String("runmode")
	if runMode == "dev" {
		orm.Debug = true
	}
	//Start beego
	web.Run()
	//beego.RunWithMiddleWares("", middleware.I18nMiddleware)
}

// Load languages files
func languages() {
	logs.Info("Application language files")
	langs := i18n.ApplicationLanguages()
	for _, lang := range langs {
		logs.Info("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "static/i18n/"+"locale_"+lang+".ini"); err != nil {
			logs.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}
