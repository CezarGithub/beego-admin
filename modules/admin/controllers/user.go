package controllers

import (
	"errors"
	"quince/internal/toolbar"
	"quince/modules/admin/form"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	ut "quince/modules/admin/transactions/user"
	"quince/utils/exceloffice"
	"quince/utils/template"
	"time"

	"github.com/beego/beego/v2/server/web"
)

// UserController struct
type UserController struct {
	BaseController
}

// Index User level list page
func (uc *UserController) Index() {
	userService := services.NewUserService()
	userLevelService := services.NewUserLevelService()

	//Get user level
	userLevel := userLevelService.GetUserLevel()
	userLevelMap := make(map[int64]string)
	for _, item := range userLevel {
		userLevelMap[item.Id] = item.Name
	}

	data, pagination := userService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["user_level_map"] = userLevelMap
	add := toolbar.Html("user.add").Add(web.URLFor("UserController.Add"))
	toggle := toolbar.Ajax("user.toggle").ToggleSelectedRows(web.URLFor("UserController.Toggle"))
	delete := toolbar.Ajax("user.delete").DeleteSelectedRows(web.URLFor("UserController.Delete"))
	export := toolbar.Ajax("user.export").Export(web.URLFor("UserController.Export"))

	edit := toolbar.Html("user.edit").Edit(web.URLFor("UserController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	uc.AddButtons(edit, add, toggle, delete, export)
	uc.TplName = "admin/views/user/index.html"
}

// Export
func (uc *UserController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		userService := services.NewUserService()
		userLevelService := services.NewUserLevelService()
		userLevel := userLevelService.GetUserLevel()
		userLevelMap := make(map[int64]string)
		for _, item := range userLevel {
			userLevelMap[item.Id] = item.Name
		}

		data := userService.GetExportData(gQueryParams)
		header := []string{"ID", uc.Translate("app.avatar"), uc.Translate("app.user_level"), uc.Translate("login.username"), uc.Translate("app.phone"), uc.Translate("app.nickname"), uc.Translate("app.enabled"), uc.Translate("app.creation_time")}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				(string)(item.Id),
				item.Avatar,
			}
			userLevelName, ok := userLevelMap[item.UserLevel.Id]
			if ok {
				record = append(record, userLevelName)
			}
			record = append(record, item.Username)
			record = append(record, item.Mobile)
			record = append(record, item.Nickname)

			if item.Status == 1 {
				record = append(record, uc.Translate("app.yes"))
			} else {
				record = append(record, uc.Translate("app.no"))
			}
			record = append(record, template.UnixTimeForFormat(int(item.CreateTime.Unix())))
			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "user-"+time.Now().Format("2006-01-02-15-04-05"), "", "", uc.Ctx.ResponseWriter)
	}

	uc.ResponseError(uc.Ctx)
}

// Add user
func (uc *UserController) Add() {
	user := models.User{}
	user.UserLevel = &models.UserLevel{}
	uc.Data["data"] = user
	userLevelService := services.NewUserLevelService()
	//
	userLevel := userLevelService.GetUserLevel()

	uc.Data["user_level_list"] = userLevel
	save := toolbar.Ajax("user.submit").Submit("userForm", web.URLFor("UserController.Update"))
	uc.AddButtons(save)
	uc.TplName = "admin/views/user/edit.html"
}

// Edit user
func (uc *UserController) Edit() {
	id, _ := uc.GetInt64("id", -1)
	if id <= 0 {
		uc.ResponseErrorWithMessage(errors.New("error.param_error"), uc.Ctx)
	}

	userService := services.NewUserService()

	user := userService.GetUserById(id)
	if user == nil {
		uc.ResponseErrorWithMessage(errors.New("error.info_not_found"), uc.Ctx)
	}

	//Get user level
	userLevelService := services.NewUserLevelService()
	userLevel := userLevelService.GetUserLevel()

	uc.Data["user_level_list"] = userLevel
	uc.Data["data"] = user
	save := toolbar.Ajax("user.submit").Submit("userForm", web.URLFor("UserController.Update"))
	uc.AddButtons(save)
	uc.TplName = "admin/views/user/edit.html"
}

// Update user
func (uc *UserController) Update() {
	var userForm form.UserForm
	if err := uc.ParseForm(&userForm); err != nil {
		uc.ResponseErrorWithMessage(err, uc.Ctx)
	}

	if userForm.Id <= 0 {
		uc.ResponseErrorWithMessage(errors.New("error.param_error"), uc.Ctx)
	}

	_, _, err := uc.GetFile("avatar")
	if err == nil {
		//Process image upload
		attachmentService := services.NewAttachmentService()
		attachmentInfo, err := attachmentService.Upload(uc.Ctx, "avatar", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			uc.ResponseErrorWithMessage(err, uc.Ctx)
		} else {
			userForm.Avatar = attachmentInfo.Url
		}
	}
	c, err := userForm.Validate()
	if err != nil {
		uc.ResponseErrorWithMessage(err, uc.Ctx)
	}

	t1 := ut.UserUpdate(c)
	uc.Transaction.Add(t1)
	err = uc.Transaction.Execute()

	if err == nil {
		uc.ResponseSuccess(uc.Ctx)
	} else {
		uc.ResponseErrorWithMessage(err, uc.Ctx)
	}

}

// Del
func (uc *UserController) Del() {
	idArr := uc.GetSelectedIDs()
	t1 := ut.UserDelete(idArr)
	uc.Transaction.Add(t1)
	err := uc.Transaction.Execute()

	if err == nil {
		uc.ResponseSuccess(uc.Ctx)
	} else {
		uc.ResponseErrorWithMessage(err, uc.Ctx)
	}
}
func (auc *UserController) Search() {
	userService := services.NewUserService()
	data := userService.GetAll(*auc.GetQueryParams())
	auc.Data["json"] = &data
	auc.ServeJSON()
}

func (uc *UserController) Toggle() {

	idArr := uc.GetSelectedIDs()
	t1 := ut.UserToggleStatus(idArr)
	uc.Transaction.Add(t1)
	err := uc.Transaction.Execute()

	if err == nil {
		uc.ResponseSuccess(uc.Ctx)
	} else {
		uc.ResponseErrorWithMessage(err, uc.Ctx)
	}

}
