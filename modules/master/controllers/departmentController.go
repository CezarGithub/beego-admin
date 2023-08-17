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

// DepartmentController struct
type DepartmentController struct {
	controllers.BaseController
}

// NestPrepare - Init controller
func (mc *DepartmentController) NestPrepare() {}

func (mc *DepartmentController) Index() {
	departmentService := services.NewDepartmentService()
	data := departmentService.GetAll(*mc.GetQueryParams())
	mc.Data["data"] = data
	add := toolbar.Html("department.add").Add(web.URLFor("DepartmentController.Add"))
	edit := toolbar.Html("department.edit").Edit(web.URLFor("DepartmentController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	mc.AddButtons(add, edit)
	mc.TplName = "master/views/department/index.html"
}
func (mc *DepartmentController) Add() {
	item := models.NewDepartment()
	mc.show(&item)
}
func (mc *DepartmentController) Edit() {
	id, _ := mc.GetInt64("id", -1)
	departmentService := services.NewDepartmentService()
	item := departmentService.GetDepartmentById(id)
	if item == nil {
		mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	mc.show(item)
}
func (mc *DepartmentController) show(item *models.Department) {
	mc.Data["data"] = item
	save := toolbar.Ajax("department.submit").Submit("departmentForm", web.URLFor("DepartmentController.Update"))
	mc.AddButtons(save)
	service := services.NewCompanyService()
	company_list := service.GetAll(*mc.GetQueryParams())
	mc.Data["company_list"] = company_list
	mc.TplName = "master/views/department/edit.html"
}
func (mc *DepartmentController) Search() {
	departmentService := services.NewDepartmentService()
	data := departmentService.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}
func (mc *DepartmentController) Update() {
	departmentForm := form.DepartmentForm{}
	if err := mc.ParseForm(&departmentForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := departmentForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	t1 := transactions.DepartmentUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
