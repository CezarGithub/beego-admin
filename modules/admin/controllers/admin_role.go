package controllers

import (
	"errors"
	"fmt"

	"quince/internal/toolbar"
	"quince/modules/admin/form"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"quince/modules/admin/transactions/admin_rbac"
	"quince/modules/admin/transactions/admin_role"
	"quince/utils"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

// AdminRoleController struct.
type AdminRoleController struct {
	BaseController
}

// Index Role Management Home.
func (arc *AdminRoleController) Index() {
	adminRoleService := services.NewAdminRoleService()
	data, pagination := adminRoleService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	arc.Data["data"] = data
	arc.Data["paginate"] = pagination
	add := toolbar.Html("admin_role.add").Add(web.URLFor("AdminRoleController.Add"))
	delete := toolbar.Ajax("admin_role.delete").DeleteSelectedRows(web.URLFor("AdminRoleController.Delete"))
	toggle := toolbar.Ajax("admin_role.toggle").ToggleSelectedRows(web.URLFor("AdminRoleController.Toggle"))
	edit := toolbar.Html("admin_role.edit").Edit(web.URLFor("AdminRoleController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	actions := toolbar.Html("user.actions_rights").New("admin.user.actions_rights", "fa fa-users", web.URLFor("AdminRoleController.Routes"))
	actions.AddQueryParams("id", "{{.Item.Id}}")
	menu := toolbar.Html("user.menu_rights").New("admin.user.menu_rights", "fa fa-bars", web.URLFor("AdminRoleController.Menus"))
	menu.AddQueryParams("id", "{{.Item.Id}}")

	arc.AddButtons(edit, actions, menu, add, delete, toggle)

	arc.TplName = "admin/views/admin_role/index.html"
}

// Add Role management-add interface.
func (arc *AdminRoleController) Add() {
	adminRole := models.AdminRole{}
	arc.Data["data"] = adminRole
	save := toolbar.Ajax("admin_role.submit").Submit("admin_roleForm", web.URLFor("AdminRoleController.Update"))
	arc.AddButtons(save)
	arc.TplName = "admin/views/admin_role/edit.html"
}

// Edit Menu Management-Role Management-Modify Interface.
func (arc *AdminRoleController) Edit() {
	id, _ := arc.GetInt64("id", -1)
	if id <= 0 {
		arc.ResponseErrorWithMessage(errors.New("error.param_error"), arc.Ctx)
	}
	adminRoleService := services.NewAdminRoleService()
	adminRole := adminRoleService.GetAdminRoleById(id)
	if adminRole == nil {
		arc.ResponseErrorWithMessage(errors.New("error.info_not_found"), arc.Ctx)
	}
	arc.Data["data"] = adminRole
	save := toolbar.Ajax("admin_role.submit").Submit("admin_roleForm", web.URLFor("AdminRoleController.Update"))
	arc.AddButtons(save)
	arc.TplName = "admin/views/admin_role/edit.html"
}

// Update Menu Management-Role Management-Modify.
func (arc *AdminRoleController) Update() {
	var adminRoleForm form.AdminRoleForm
	if err := arc.ParseForm(&adminRoleForm); err != nil {
		arc.ResponseErrorWithMessage(err, arc.Ctx)
	}

	c, err := adminRoleForm.Validate()
	if err != nil {
		arc.ResponseErrorWithMessage(err, arc.Ctx)
	}

	t1 := admin_role.AdminRoleUpdate(c)
	arc.Transaction.Add(t1)
	err = arc.Transaction.Execute()

	if err == nil {
		arc.ResponseSuccess(arc.Ctx)
	} else {
		arc.ResponseErrorWithMessage(err, arc.Ctx)
	}
}

// Deldelete.
func (arc *AdminRoleController) Delete() {
	idArr := arc.GetSelectedIDs()
	t1 := admin_role.AdminRoleDelete(idArr)
	arc.Transaction.Add(t1)
	err := arc.Transaction.Execute()

	if err == nil {
		arc.ResponseSuccess(arc.Ctx)
	} else {
		arc.ResponseErrorWithMessage(err, arc.Ctx)
	}
}
func (arc *AdminRoleController) AccessRbac() {
	var err error
	id, _ := arc.GetInt64("idRole", -1)
	if id <= 0 {
		arc.ResponseErrorWithMessage(errors.New("error.param_error"), arc.Ctx)
	}

	idArr := arc.GetSelectedIDs()
	adminRbacService := services.NewAdminRbacService()
	adminRouteService := services.NewAdminRouteService()
	adminRoleService := services.NewAdminRoleService()
	role := adminRoleService.GetAdminRoleById(id)
	for i, val := range idArr {
		b := adminRbacService.IsExistUrl(id, val)
		route := adminRouteService.GetAdminRouteById(val)
		if b == nil { //not found - set ID to -id
			rbac := models.AdminRbac{Role: role, Route: route, Status: 1}
			t := admin_rbac.AdminRbacUpdate(&rbac)
			idArr[i] = -1
			arc.Transaction.Add(t)
		} else {
			idArr[i] = b.Id // replace with Rbac id
		}
	}
	t1 := admin_rbac.AdminRbacDelete(idArr)
	arc.Transaction.Add(t1)
	err = arc.Transaction.Execute()

	if err == nil {
		arc.ResponseSuccess(arc.Ctx)
	} else {
		arc.ResponseErrorWithMessage(err, arc.Ctx)
	}
}
func (arc *AdminRoleController) Toggle() {
	idArr := arc.GetSelectedIDs()
	t1 := admin_role.AdminRoleToggleStatus(idArr)
	arc.Transaction.Add(t1)
	err := arc.Transaction.Execute()

	if err == nil {
		arc.ResponseSuccess(arc.Ctx)
	} else {
		arc.ResponseErrorWithMessage(err, arc.Ctx)
	}

}

// Access Menu Management-Role Management-Role Authorization Interface.
func (arc *AdminRoleController) Menus() {
	id, _ := arc.GetInt64("id", -1)
	if id <= 0 {
		arc.ResponseErrorWithMessage(errors.New("error.param_error"), arc.Ctx)
	}

	var (
		adminTreeService services.AdminTreeService
	)
	adminRoleService := services.NewAdminRoleService()
	adminMenuService := services.NewAdminMenuService()
	data := adminRoleService.GetAdminRoleById(id)
	menu := adminMenuService.AllMenu()

	menuMap := make(map[int64]orm.Params)

	for _, adminMenu := range menu {
		id := adminMenu.Id
		if menuMap[id] == nil {
			menuMap[id] = make(orm.Params)
		}
		menuMap[id]["Id"] = id
		menuMap[id]["ParentId"] = adminMenu.ParentId
		menuMap[id]["Name"] = adminMenu.Name
		menuMap[id]["Url"] = adminMenu.Url
		menuMap[id]["Icon"] = adminMenu.Icon
		menuMap[id]["IsShow"] = adminMenu.IsShow
		menuMap[id]["SortId"] = adminMenu.SortId
		menuMap[id]["LogMethod"] = adminMenu.LogMethod
	}

	html := adminTreeService.AuthorizeHtml(menuMap, strings.Split(data.Url, ","))

	arc.Data["data"] = data
	arc.Data["html"] = html
	save := toolbar.Ajax("admin_role.acces").Submit("accessForm", web.URLFor("AdminRoleController.AccessOperate"))
	arc.AddButtons(save)
	arc.TplName = "admin/views/admin_role/access.html"
}

func (arc *AdminRoleController) Routes() {
	id, _ := arc.GetInt64("id", -1)
	if id <= 0 {
		arc.ResponseErrorWithMessage(errors.New("error.param_error"), arc.Ctx)
	}
	adminRoleService := services.NewAdminRoleService()
	roleName := adminRoleService.GetAdminRoleById(id)
	adminRouteService := services.NewAdminRouteService()
	data, pagination := adminRouteService.GetRoleRoutesById(id, admin["per_page"].(int), gQueryParams)

	arc.Data["data"] = data
	arc.Data["paginate"] = pagination
	arc.Data["role"] = roleName

	toggle := toolbar.Ajax("admin_roles.routes").ToggleSelectedRows(web.URLFor("AdminRoleController.AccessRbac") + fmt.Sprintf("?idRole=%d", id))
	arc.AddButtons(toggle)
	arc.TplName = "admin/views/admin_role/routes.html"
}

// AccessOperate Menu Management-Role Management-Role Authorization.
func (arc *AdminRoleController) AccessOperate() {
	id, _ := arc.GetInt64("id", -1)
	if id < 0 {
		arc.ResponseErrorWithMessage(errors.New("error.param_error"), arc.Ctx)
	}

	url := make([]string, 0)
	arc.Ctx.Input.Bind(&url, "url")

	if len(url) == 0 {
		arc.ResponseErrorWithMessage(errors.New("admin.select_menu"), arc.Ctx)
	}

	if !utils.InArrayForString(url, "1") {
		arc.ResponseErrorWithMessage(errors.New("admin.permission_required"), arc.Ctx)
	}

	c := &models.AdminRole{Base: models.Base{Id: id}, Url: strings.Join(url, ",")}
	t1 := admin_role.AdminRoleUpdate(c)
	arc.Transaction.Add(t1)
	err := arc.Transaction.Execute()

	if err == nil {
		arc.ResponseSuccess(arc.Ctx)
	} else {
		arc.ResponseErrorWithMessage(err, arc.Ctx)
	}

}
