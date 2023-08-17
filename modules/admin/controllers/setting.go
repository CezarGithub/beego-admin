package controllers

import (
	"encoding/json"
	"errors"
	"quince/global"
	"quince/internal/toolbar"
	"quince/modules/admin/services"
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

// SettingController struct
type SettingController struct {
	BaseController
}

// Admin set up
func (sc *SettingController) Admin() {
	sc.show(1)
}

// show Display single configuration information
func (sc *SettingController) show(id int64) {
	var settingService services.SettingService
	data := settingService.Show(id, sc.i18n.Lang)
	sc.Data["data_config"] = data
	save := toolbar.Ajax("setting.submit").Submit("settingsForm", web.URLFor("SettingController.Update"))
	sc.AddButtons(save)
	sc.TplName = "admin/views/setting/show.html"
}

// Update Settings Center-Update Settings
func (sc *SettingController) Update() {
	input, _ := sc.Input()
	id := input.Get("id")

	if id == "" {
		sc.ResponseErrorWithMessage(errors.New("error.param_error"), sc.Ctx)
	}

	var settingService services.SettingService
	idInt, _ := strconv.ParseInt(id, 10, 64)
	setting := settingService.GetSettingInfoById(idInt)

	if setting == nil {
		sc.ResponseErrorWithMessage(errors.New("error.update_settings"), sc.Ctx)
	} else {
		err := json.Unmarshal([]byte(setting.Content), &setting.ContentStrut)
		if err != nil {
			sc.ResponseErrorWithMessage(errors.New("error.update_settings"), sc.Ctx)
		}
	}
	for key, value := range setting.ContentStrut {
		switch value.Type {
		case "image", "file":
			//Single file upload
			attachmentService := services.NewAttachmentService()
			attachmentInfo, err := attachmentService.Upload(sc.Ctx, value.Field, loginUser.Id, 0)
			if err == nil && attachmentInfo != nil {
				//Picture uploaded successfully
				setting.ContentStrut[key].Content = attachmentInfo.Url
			}
		case "multi_file", "multi_image":
			//Multiple file upload
			attachmentService := services.NewAttachmentService()
			attachments, err := attachmentService.UploadMulti(sc.Ctx, value.Field, loginUser.Id, 0)
			if err == nil && attachments != nil {
				var urls []string
				for _, atta := range attachments {
					urls = append(urls, atta.Url)
				}
				if len(urls) > 0 {
					urlsByte, err := json.Marshal(&urls)
					if err == nil {
						setting.ContentStrut[key].Content = string(urlsByte)
					}
				}
			}
		default:
			input, _ := sc.Input()
			setting.ContentStrut[key].Content = input.Get(value.Field)
		}
	}

	//Modify content
	contentStrutByte, err := json.Marshal(&setting.ContentStrut)
	if err == nil {
		affectRow := settingService.UpdateSettingInfoToContent(idInt, string(contentStrutByte))
		if affectRow > 0 {
			//Update global configuration
			settingService.LoadOrUpdateGlobalBaseConfig(setting)
			sc.ResponseSuccessWithMessageAndUrl("ok.message", global.URL_RELOAD, sc.Ctx)
		} else {
			sc.ResponseErrorWithMessage(errors.New("error.nothing_to_update"), sc.Ctx)
		}
	} else {
		sc.ResponseErrorWithMessage(err, sc.Ctx)
	}

}
