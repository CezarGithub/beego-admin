package controllers

import (
	"errors"
	"fmt"
	"quince/internal/toolbar"
	"quince/modules/admin/controllers"
	"quince/modules/master/form"
	"quince/modules/master/models"
	"quince/modules/master/services"
	"quince/modules/master/transactions"

	"github.com/beego/beego/v2/server/web"
)

// CompanyController struct
type CompanyController struct {
	controllers.BaseController
}

// NestPrepare - Init controller
func (mc *CompanyController) NestPrepare() {

}

func (mc *CompanyController) Index() {
	companyService := services.NewCompanyService()
	data := companyService.GetAll(*mc.GetQueryParams())
	mc.Data["data"] = data
	add := toolbar.Html("company.add").Add(web.URLFor("CompanyController.Add"))
	edit := toolbar.Html("company.edit").Edit(web.URLFor("CompanyController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	mc.AddButtons(add, edit)
	//mc.ViewPath = "modules"
	//mc.Layout = "admin/views/public/base.html"
	mc.TplName = "master/views/company/index.html"
}
func (mc *CompanyController) Add() {
	item := models.Company{}
	item.Country = &models.Country{}
	item.Group = &models.Group{}
	mc.show(&item)
}
func (mc *CompanyController) Edit() {
	id, _ := mc.GetInt64("id", -1)
	companyService := services.NewCompanyService()
	item := companyService.GetCompanyById(id)
	if item == nil {
		mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	mc.show(item)
}
func (mc *CompanyController) show(item *models.Company) {
	groupService := services.NewGroupService()
	mc.Data["group_list"] = groupService.GetAll(nil)
	mc.Data["data"] = item
	mc.DataSearchBox["searchBoxCountry"] = &models.Country{}
	save := toolbar.Ajax("company.submit").Submit("companyForm", web.URLFor("CompanyController.Update"))
	smtp := toolbar.Modal("company.smtp").Edit(web.URLFor("SMTPController.Edit"))
	smtp.AddQueryParams("company_id", fmt.Sprintf("%d", item.Id))
	mc.AddButtons(save, smtp)
	mc.TplName = "master/views/company/edit.html"
}
func (mc *CompanyController) Search() {
	companyService := services.NewCompanyService()
	data := companyService.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}
func (mc *CompanyController) Update() {
	companyForm := form.CompanyForm{}
	if err := mc.ParseForm(&companyForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := companyForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
		//mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	t1 := transactions.CompanyUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
