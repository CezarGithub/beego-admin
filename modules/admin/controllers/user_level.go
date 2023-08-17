package controllers

import (
	"errors"
	"quince/internal/toolbar"
	"quince/modules/admin/form"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"quince/modules/admin/transactions/user_level"
	"quince/utils/exceloffice"
	"quince/utils/template"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web"
)

// UserLevelController struct
type UserLevelController struct {
	BaseController
}

// Index
func (ulc *UserLevelController) Index() {
	userLevelService := services.NewUserLevelService()
	data, pagination := userLevelService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	ulc.Data["data"] = data
	ulc.Data["paginate"] = pagination

	add := toolbar.Html("user_level.add").Add(web.URLFor("UserLevelController.Add"))
	toggle := toolbar.Ajax("user_level.toggle").ToggleSelectedRows(web.URLFor("UserLevelController.Toggle"))
	delete := toolbar.Ajax("user_level.delete").DeleteSelectedRows(web.URLFor("UserLevelController.Delete"))
	export :=toolbar.Ajax("user_level.export").Export(web.URLFor("UserLevelController.Export"))
	edit := toolbar.Html("user_level.edit").Edit(web.URLFor("UserLevelController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	ulc.AddButtons(edit,add, toggle, delete, export)
	ulc.TplName = "admin/views/user_level/index.html"
}

// Export
func (ulc *UserLevelController) Export() {
	exportData := ulc.GetString("export_data")
	if exportData == "1" {
		userLevelService := services.NewUserLevelService()
		data := userLevelService.GetExportData(gQueryParams)
		header := []string{"ID", ulc.Translate("app.name"), ulc.Translate("app.info"), ulc.Translate("app.enabled"), ulc.Translate("app.creation_time")}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.FormatInt(item.Id, 10),
				item.Name,
				item.Description,
			}
			if item.Status == 1 {
				record = append(record, ulc.Translate("app.yes"))
			} else {
				record = append(record, ulc.Translate("app.no"))
			}
			record = append(record, template.UnixTimeForFormat(int(item.CreateTime.Unix())))
			body = append(body, record)
		}
		ulc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "user_level-"+time.Now().Format("2006-01-02-15-04-05"), "", "", ulc.Ctx.ResponseWriter)
	}

	ulc.ResponseError(ulc.Ctx)
}

// Add
func (ulc *UserLevelController) Add() {
	userLevel := models.UserLevel{}
	ulc.Data["data"] = userLevel
	save := toolbar.Ajax("user_level.submit").Submit("user_levelForm",web.URLFor("UserLevelController.Update"))
	ulc.AddButtons(save)
	ulc.TplName = "admin/views/user_level/edit.html"
}

// Edit
func (ulc *UserLevelController) Edit() {
	id, _ := ulc.GetInt64("id", -1)
	if id <= 0 {
		ulc.ResponseErrorWithMessage(errors.New("error.param_error"), ulc.Ctx)
	}

	userLevelService := services.NewUserLevelService()

	userLevel := userLevelService.GetUserLevelById(id)
	if userLevel == nil {
		ulc.ResponseErrorWithMessage(errors.New("error.info_not_found"), ulc.Ctx)
	}

	ulc.Data["data"] = userLevel
	save := toolbar.Ajax("user_level.submit").Submit("user_levelForm",web.URLFor("UserLevelController.Update"))
	ulc.AddButtons(save)
	ulc.TplName = "admin/views/user_level/edit.html"
}

// Update
func (ulc *UserLevelController) Update() {
	var userLevelForm form.UserLevelForm

	if err := ulc.ParseForm(&userLevelForm); err != nil {
		ulc.ResponseErrorWithMessage(err, ulc.Ctx)
	}

	if userLevelForm.Id <= 0 {
		ulc.ResponseErrorWithMessage(errors.New("error.param_error"), ulc.Ctx)
	}

	_, _, err := ulc.GetFile("img")
	if err == nil {
		//Process image upload
		attachmentService := services.NewAttachmentService()
		attachmentInfo, err := attachmentService.Upload(ulc.Ctx, "img", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			ulc.ResponseErrorWithMessage(err, ulc.Ctx)
		} else {
			userLevelForm.Img = attachmentInfo.Url
		}
	}

	c, err := userLevelForm.Validate()
	if err != nil {
		ulc.ResponseErrorWithMessage(err, ulc.Ctx)
	}

	t1 := user_level.UserLevelUpdate(c)
	ulc.Transaction.Add(t1)
	err = ulc.Transaction.Execute()

	if err == nil {
		ulc.ResponseSuccess(ulc.Ctx)
	} else {
		ulc.ResponseErrorWithMessage(err, ulc.Ctx)
	}

}

func (auc *UserLevelController) Toggle() {

	idArr := auc.GetSelectedIDs()
	t1 := user_level.UserLevelToggleStatus(idArr)
	auc.Transaction.Add(t1)
	err := auc.Transaction.Execute()

	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}

}

// Del
func (auc *UserLevelController) Del() {
	idArr := auc.GetSelectedIDs()
	t1 := user_level.UserLevelDelete(idArr)
	auc.Transaction.Add(t1)
	err := auc.Transaction.Execute()

	if err == nil {
		auc.ResponseSuccess(auc.Ctx)
	} else {
		auc.ResponseErrorWithMessage(err, auc.Ctx)
	}
}
