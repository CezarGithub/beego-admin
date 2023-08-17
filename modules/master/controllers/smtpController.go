package controllers

import (
	"errors"
	"quince/internal/toolbar"
	"quince/modules/admin/controllers"
	"quince/modules/master/form"
	"quince/modules/master/models"
	"quince/modules/master/services"
	"quince/modules/master/transactions"

	"github.com/beego/beego/v2/server/web"
)

// SMTPController struct
type SMTPController struct {
	controllers.BaseController
}

func (mc *SMTPController) Index() {
	company_id, _ := mc.GetInt64("cid", -1)
	smtpService := services.NewSMTPService()
	smtpService.IModel = models.NewSMTP(company_id)
	data, pagination := smtpService.GetPaginateData(mc.GetAdminMap("per_page").(int), *mc.GetQueryParams())
	mc.Data["paginate"] = pagination
	mc.Data["data"] = data
	add := toolbar.Html("smtp.add").Add(web.URLFor("SMTPController.Add"))
	edit := toolbar.Html("smtp.edit").Edit(web.URLFor("SMTPController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	mc.AddButtons(add, edit)
	mc.TplName = "master/views/smtp/index.html"
}
func (mc *SMTPController) Add() {
	item := models.SMTP{}
	item.Company = &models.Company{}
	mc.show(&item)
}
func (mc *SMTPController) Edit() {
	id, _ := mc.GetInt64("id", -1)
	companyId, _ := mc.GetInt64("company_id", -1)
	smtpService := services.NewSMTPService()
	item := smtpService.GetSMTPById(id)
	if item == nil {
		item = smtpService.GetSMTPByCompanyId(companyId)
	}
	if item == nil {
		mc.ResponseErrorWithMessageAndUrl(errors.New("error.info_not_found"), web.URLFor("SMTPController.Index"), mc.Ctx)
		return
	}
	mc.show(item)
}
func (mc *SMTPController) show(item *models.SMTP) {
	companyService := services.NewCompanyService()
	list := companyService.GetAll(nil)
	mc.Data["companies_list"] = list
	mc.Data["data"] = item
	save := toolbar.Ajax("smtp.submit").Submit("smtpForm", web.URLFor("SMTPController.Update"))
	mc.AddButtons(save)
	mc.TplName = "master/views/smtp/edit.html"

}
func (mc *SMTPController) Update() {
	smtpForm := form.SMTPForm{}
	if err := mc.ParseForm(&smtpForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := smtpForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
	t1 := transactions.SmtpUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
func (mc *SMTPController) Search() {
	smtpService := services.NewSMTPService()
	data := smtpService.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}
