package controllers

import (
	"encoding/base64"
	"errors"
	"quince/internal/toolbar"
	"quince/modules/admin/form"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"quince/modules/admin/transactions/admin_user"
	ut "quince/modules/admin/transactions/user"
	"quince/utils"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

// AdminUserController struct
type AdminUserController struct {
	BaseController
}

// Index User Management-Home
func (auc *AdminUserController) Index() {
	adminUserService := services.NewAdminUserService()
	data, pagination := adminUserService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	auc.Data["data"] = data
	auc.Data["paginate"] = pagination
	add := toolbar.Html("admin_user.add").Add(web.URLFor("AdminUserController.Add"))
	toggle := toolbar.Ajax("admin_user.toggle").ToggleSelectedRows(web.URLFor("AdminUserController.Toggle"))
	edit := toolbar.Html("admin_user.edit").Edit(web.URLFor("AdminUserController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	auc.AddButtons(edit,add,toggle)
	auc.TplName = "admin/views/admin_user/index.html"
}

// Add User management-add interface
func (auc *AdminUserController) Add() {
	languageService := services.NewLanguageService()
	langs := languageService.GetApplicationLanguages()
	adminRoleService := services.NewAdminRoleService()
	roles := adminRoleService.GetAllData()
	auc.Data["languages"] = langs
	auc.Data["roles"] = roles
	auc.Data["data"] = models.NewLoginUser()
	save := toolbar.Ajax("admin_user.submit").Submit("admin_userForm", web.URLFor("AdminUserController.Update"))
	auc.AddButtons(save)
	auc.DataSearchBox["searchBoxUser"] = &models.User{}
	auc.TplName = "admin/views/admin_user/edit.html"
}

// Edit System Management-User Management-Modify Interface
func (auc *AdminUserController) Edit() {
	id, _ := auc.GetInt64("id", -1)
	if id <= 0 {
		auc.ResponseErrorWithMessage(errors.New("error.param_error"), auc.Ctx)
	}

	adminRoleService := services.NewAdminRoleService()
	adminUserService := services.NewAdminUserService()
	languageService := services.NewLanguageService()
	adminUser := adminUserService.GetAdminUserById(id)
	if adminUser == nil {
		auc.ResponseErrorWithMessage(errors.New("admin.info_not_found"), auc.Ctx)
	}

	roles := adminRoleService.GetAllData()
	langs := languageService.GetApplicationLanguages()
	auc.Data["roles"] = roles
	auc.Data["languages"] = langs
	auc.Data["data"] = adminUser
	if adminUser != nil {
		auc.Data["role_arr"] = strings.Split(adminUser.Roles, ",")
	} else {
		auc.Data["role_arr"] = ""
	}
	save := toolbar.Ajax("admin_user.submit").Submit("admin_userForm", web.URLFor("AdminUserController.Update"))
	auc.AddButtons(save)
	auc.DataSearchBox["searchBoxUser"] = &models.User{}
	auc.TplName = "admin/views/admin_user/edit.html"
}

// Update System Management-User Management-Modification
func (auc *AdminUserController) Update() {
	var adminUserForm form.AdminUserForm
	if err := auc.ParseForm(&adminUserForm); err != nil {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}
	roles := make([]string, 0)
	auc.Ctx.Input.Bind(&roles, "role")
	adminUserForm.Roles = strings.Join(roles, ",")
	c, err := adminUserForm.Validate()
	if err != nil {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}

	t1 := admin_user.LoginUserUpdate(c)
	auc.Transaction.Add(t1)
	err = auc.Transaction.Execute()

	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}

}
func (auc *AdminUserController) Toggle() {

	idArr := auc.GetSelectedIDs()
	t1 := admin_user.LoginUserToggleStatus(idArr)
	auc.Transaction.Add(t1)
	err := auc.Transaction.Execute()

	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}

}

// Del
func (auc *AdminUserController) Del() {
	idArr := auc.GetSelectedIDs()
	t1 := admin_user.LoginUserDelete(idArr)
	auc.Transaction.Add(t1)
	err := auc.Transaction.Execute()

	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}
}

// Profile System Management-Personal Information
func (auc *AdminUserController) Profile() {

	auc.TplName = "admin/views/admin_user/profile.html"
}

func (auc *AdminUserController) UpdateNickName() {
	//id, err := auc.GetInt("id")
	nickname := strings.TrimSpace(auc.GetString("nickname"))

	if nickname == "" { //} || err != nil {
		auc.ResponseErrorWithMessage(errors.New("error.param_error"), auc.Ctx)
	}
	loginUser := auc.GetAdminMap("user").(models.LoginUser) //models.LoginUser{Base: models.Base{Id: id}}
	user := loginUser.User
	user.Nickname = nickname
	t1 := ut.UserUpdate(user)
	auc.Transaction.Add(t1)
	err := auc.Transaction.Execute()
	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}
}
func (auc *AdminUserController) UpdatePassword() {
	id, err := auc.GetInt64("id")
	password := auc.GetString("password")
	newPassword := auc.GetString("new_password")
	reNewPassword := auc.GetString("renew_password")

	if err != nil || password == "" || newPassword == "" || reNewPassword == "" {
		auc.ResponseErrorWithMessage(errors.New("error.param_error"), auc.Ctx)
	}

	if newPassword != reNewPassword {
		auc.ResponseErrorWithMessage(errors.New("admin.password_inconsistent"), auc.Ctx)
	}

	if password == newPassword { //nothing changed
		auc.ResponseSuccessWithMessage("ok.message", auc.Ctx)
	}

	loginUserPassword, err := base64.StdEncoding.DecodeString(loginUser.Password)

	if err != nil {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}

	if !utils.PasswordVerify(password, string(loginUserPassword)) {
		auc.ResponseErrorWithMessage(errors.New("admin.password_wrong"), auc.Ctx)
	}
	user := models.LoginUser{Base: models.Base{Id: id}}
	o := orm.NewOrm()
	if o.Read(&user) == nil {
		user.Password = newPassword
	} else {
		auc.ResponseErrorWithMessage(errors.New("error.id_error"), auc.Ctx)
	}
	t1 := admin_user.LoginUserUpdate(&user)
	auc.Transaction.Add(t1)
	err = auc.Transaction.Execute()

	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}
}

func (auc *AdminUserController) UpdateAvatar() {
	avatar := strings.TrimSpace(auc.GetString("nickname"))

	if avatar == "" { //} || err != nil {
		auc.ResponseErrorWithMessage(errors.New("error.param_error"), auc.Ctx)
	}
	loginUser := auc.GetAdminMap("user").(models.LoginUser) //models.LoginUser{Base: models.Base{Id: id}}
	user := loginUser.User
	user.Avatar = avatar
	t1 := ut.UserUpdate(user)
	auc.Transaction.Add(t1)
	err := auc.Transaction.Execute()
	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}
}
