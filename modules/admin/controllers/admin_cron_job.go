package controllers

import (
	"errors"
	"fmt"
	"quince/initialize/scheduler"

	"quince/internal/toolbar"
	"quince/modules/admin/form"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	acj "quince/modules/admin/transactions/admin_cron_job"

	"github.com/beego/beego/v2/server/web"
)

// AdminCronJobController struct.
type AdminCronJobController struct {
	BaseController
}

// Index index.
func (alc *AdminCronJobController) Index() {

	adminCronJobService := services.NewAdminCronJobService()
	data, pagination := adminCronJobService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	alc.Data["data"] = data
	alc.Data["paginate"] = pagination
	add := toolbar.Html("cron_job.add").Add(web.URLFor("AdminCronJobController.Add"))
	toggle := toolbar.Ajax("cron_job.toggle").ToggleSelectedRows(web.URLFor("AdminCronJobController.Toggle"))
	delete := toolbar.Ajax("cron_job.delete").DeleteSelectedRows(web.URLFor("AdminCronJobController.Delete"))
	edit := toolbar.Html("cron_job.edit").Edit(web.URLFor("AdminCronJobController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	alc.AddButtons(edit, add, toggle, delete)
	alc.TplName = "admin/views/admin_cron_job/index.html"
}
func (mc *AdminCronJobController) Add() {
	item := models.AdminCronJob{}
	//item.Company=&models.Company{}
	mc.show(&item)
}
func (mc *AdminCronJobController) Edit() {
	id, _ := mc.GetInt("id", -1)
	service := services.NewAdminCronJobService()
	item := service.GetAdminCronJobById(id)
	if item == nil {
		mc.ResponseErrorWithMessageAndUrl(errors.New("error.info_not_found"), web.URLFor("SMTPController.Index"), mc.Ctx)
		return
	} else {
		for _, u := range item.Intervals {
			u.Description = mc.intervalDescription(*u)
		}
	}

	mc.show(item)
}
func (mc *AdminCronJobController) show(item *models.AdminCronJob) {

	mc.DataSearchBox["searchBoxCron"] = &models.AdminCronJob{}
	mc.Data["data"] = item
	save := toolbar.Ajax("cron_job.submit").Submit("cron_jobForm", web.URLFor("AdminCronJobController.Update"))
	edit := toolbar.Modal("cron_job.details").Edit(web.URLFor("AdminCronJobController.Edit"))
	add := toolbar.Modal("cron_job.add_details").Edit(web.URLFor("AdminCronJobController.Add"))
	edit.AddQueryParams("id", fmt.Sprintf("%d", item.Id))
	mc.AddButtons(save, edit,add)
	mc.TplName = "admin/views/admin_cron_job/edit.html"
}

func (mc *AdminCronJobController) intervalDescription(u models.AdminCronJobInterval) string {
	var msg, month, dayOfMonth, dayOfWeek, hour, repeat string

	if u.Interval.MonthOfTheYear > -1 {
		month = mc.i18n.Tr(fmt.Sprintf("app.month.%d", u.Interval.MonthOfTheYear))
	} else {
		month = mc.i18n.Tr("admin.cron.each") + mc.i18n.Tr("admin.cron.month")
	}
	if u.Interval.DayOfWeek > -1 {
		dayOfWeek = mc.i18n.Tr(fmt.Sprintf("app.day.%d", u.Interval.DayOfWeek))
	}
	if u.Interval.DayOfTheMonth > -1 {
		dayOfMonth = fmt.Sprintf("%d", u.Interval.DayOfWeek)
	}
	if u.Interval.Hour > -1 {
		if u.Interval.Minute > -1 {
			hour = fmt.Sprintf("%02d:%02d", u.Interval.Hour, u.Interval.Minute)
		} else {
			hour = fmt.Sprintf("%02d:%02d", u.Interval.Hour, 0)
		}
	} else {
		if u.Interval.Minute > -1 {
			hour = mc.i18n.Tr("admin.cron.each") + mc.i18n.Tr("admin.cron.hour") + fmt.Sprintf(":%02d", u.Interval.Minute)
		}
	}
	if u.RunOnce == 1 {
		repeat = mc.i18n.Tr("admin.cron.runOnce") + " -> " + hour
	} else {
		repeat = mc.i18n.Tr("admin.cron.repeat") + " -> " + hour
	}
	msg = fmt.Sprintf("%s, %s, %s, %s", month, dayOfMonth, dayOfWeek, repeat)
	return msg
}

func (auc *AdminCronJobController) Toggle() {

	idArr := auc.GetSelectedIDs()
	t1 := acj.AdminCronJobToggleStatus(idArr)
	auc.Transaction.Add(t1)
	err := auc.Transaction.Execute()

	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}

}
func (mc *AdminCronJobController) Update() {
	cForm := form.AdminCronJobForm{}
	if err := mc.ParseForm(&cForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := cForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
	t1 := acj.AdminCronJobUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
func (mc *AdminCronJobController) Search() {
	var data []*models.AdminCronJob
	jobs := scheduler.CronJobs()
	for i, j := range jobs {
		t := models.AdminCronJob{}
		t.SetID(int64(i))
		t.Name = j.Name
		t.Description = j.Description
		t.Module = j.Module
		data = append(data, &t)
	}
	mc.Data["json"] = &data
	mc.ServeJSON()
}

// Del
func (mc *AdminCronJobController) Delete() {
	idArr := mc.GetSelectedIDs()
	t1 := acj.AdminCronJobDelete(idArr)
	mc.Transaction.Add(t1)
	err := mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
