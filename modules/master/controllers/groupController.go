package controllers

import (
	"errors"
	"quince/modules/admin/controllers"

	"quince/modules/master/form"
	"quince/modules/master/models"
	"quince/modules/master/services"
	"quince/modules/master/transactions"

	"quince/internal/toolbar"

	"github.com/beego/beego/v2/server/web"
)

// GroupController struct
type GroupController struct {
	controllers.BaseController
}

func (mc *GroupController) Index() {
	groupService := services.NewGroupService()
	data, pagination := groupService.GetPaginateData(mc.GetAdminMap("per_page").(int), *mc.GetQueryParams())
	mc.Data["paginate"] = pagination
	mc.Data["data"] = data
	add := toolbar.Html("group.add").Add(web.URLFor("GroupController.Add"))
	edit := toolbar.Html("group.edit").Edit(web.URLFor("GroupController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	mc.AddButtons(add, edit)
	mc.TplName = "master/views/group/index.html"
}

func (mc *GroupController) Add() {
	item := new(models.Group)
	mc.show(item)
}
func (mc *GroupController) Edit() {
	id, _ := mc.GetInt("id", -1)
	groupService := services.NewGroupService()
	item := groupService.GetGroupById(id)
	if item == nil {
		mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	mc.show(item)
}
func (mc *GroupController) show(item *models.Group) {

	save := toolbar.Ajax("group.submit").Submit("groupForm", web.URLFor("GroupController.Update"))
	mc.AddButtons(save)
	mc.Data["data"] = item
	mc.TplName = "master/views/group/edit.html"
}
func (mc *GroupController) Search() {
	groupService := services.NewGroupService()
	data := groupService.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}

func (mc *GroupController) Update() {
	groupForm := form.GroupForm{}
	if err := mc.ParseForm(&groupForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := groupForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
		//mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	t1 := transactions.GroupUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
