package controllers

import (
	"errors"
	"quince/global"
	"quince/internal/toolbar"
	"quince/modules/admin/form"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"quince/modules/admin/transactions/admin_menu"
	"strings"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/server/web"
)

// AdminMenuController struct.
type AdminMenuController struct {
	BaseController
}

// Index Menu home.
func (amc *AdminMenuController) Index() {
	var adminTreeService services.AdminTreeService
	m := adminTreeService.AdminMenuTree()
	m = strings.ReplaceAll(m, "[EDIT]", amc.Translate("app.edit"))
	m = strings.ReplaceAll(m, "[DELETE]", amc.Translate("app.delete"))
	m = strings.ReplaceAll(m, "[DELETE_QUESTION]", amc.Translate("app.delete_question"))
	amc.Data["data"] = m
	add := toolbar.Html("admin_menu.add").Add(web.URLFor("AdminMenuController.Add"))
	toggle := toolbar.Ajax("admin_menu.toggle").ToggleSelectedRows(web.URLFor("AdminMenuController.Toggle"))
	amc.AddButtons(add, toggle)
	logs.Info("AdminMenuController.Index refactoring required !")

	amc.TplName = "admin/views/admin_menu/index.html"
}

// Add menu.
func (amc *AdminMenuController) Add() {

	adminMenu := models.AdminMenu{}
	var adminTreeService services.AdminTreeService
	parentID, _ := amc.GetInt64("parent_id", 0)
	parents := adminTreeService.Menu(parentID, 0)
	amc.Data["data"] = adminMenu
	amc.Data["parents"] = parents
	amc.Data["log_method"] = new(models.AdminMenu).GetLogMethod()
	save := toolbar.Ajax("admin_menu.submit").Submit("admin_menuForm", web.URLFor("AdminMenuController.Update"))
	amc.AddButtons(save)

	amc.TplName = "admin/views/admin_menu/edit.html"
}

// Edit menu.
func (amc *AdminMenuController) Edit() {
	id, _ := amc.GetInt64("id", -1)
	if id <= 0 {
		amc.ResponseErrorWithMessage(errors.New("error.param_error"), amc.Ctx)
	}

	var (
		adminTreeService services.AdminTreeService
	)
	adminMenuService := services.NewAdminMenuService()
	adminMenu := adminMenuService.GetAdminMenuById(id)
	if adminMenu == nil {
		amc.ResponseErrorWithMessage(errors.New("error.info_not_found"), amc.Ctx)
	}
	var parentID int64
	if adminMenu != nil {
		parentID = adminMenu.ParentId
	}
	parents := adminTreeService.Menu(parentID, 0)

	amc.Data["data"] = adminMenu

	amc.Data["parents"] = parents
	amc.Data["log_method"] = new(models.AdminMenu).GetLogMethod()
	save := toolbar.Ajax("admin_menu.submit").Submit("admin_menuForm", web.URLFor("AdminMenuController.Update"))
	amc.AddButtons(save)

	amc.TplName = "admin/views/admin_menu/edit.html"
}

// Update menu.
func (amc *AdminMenuController) Update() {
	adminMenuForm := form.AdminMenuForm{}

	if err := amc.ParseForm(&adminMenuForm); err != nil {
		amc.ResponseErrorWithMessage(err, amc.Ctx)
	}
	c, err := adminMenuForm.Validate()
	if err != nil {
		amc.ResponseErrorWithMessage(err, amc.Ctx)
	}

	t1 := admin_menu.AdminMenuUpdate(c)
	amc.Transaction.Add(t1)
	err = amc.Transaction.Execute()

	if err == nil {
		amc.ResponseSuccess(amc.Ctx)
	} else {
		amc.ResponseErrorWithMessage(err, amc.Ctx)
	}
}

// Del .
func (amc *AdminMenuController) Delete() {
	idArr := amc.GetSelectedIDs()
	t1 := admin_menu.AdminMenuDelete(idArr)
	amc.Transaction.Add(t1)
	err := amc.Transaction.Execute()
	if err == nil {
		amc.ResponseSuccess(amc.Ctx)
	} else {
		amc.ResponseErrorWithMessage(err, amc.Ctx)
	}
}
func (amc *AdminMenuController) Toggle() {
	idArr := amc.GetSelectedIDs()
	t1 := admin_menu.AdminMenuToggleStatus(idArr)
	amc.Transaction.Add(t1)
	err := amc.Transaction.Execute()

	if err == nil {
		amc.ResponseSuccess(amc.Ctx)
	} else {
		amc.ResponseErrorWithMessage(err, amc.Ctx)
	}

}
func (amc *AdminMenuController) LeftMenu() {
	adminMenuService := services.NewAdminMenuService()
	loginUser, _ := amc.GetSession(global.LOGIN_USER).(models.LoginUser)
	json := adminMenuService.GetLeftMenu(amc.Ctx.Request.RequestURI, loginUser, loginUser.Language)
	amc.Data["json"] = json
	amc.ServeJSON()
}
