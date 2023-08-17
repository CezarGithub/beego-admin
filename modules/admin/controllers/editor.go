package controllers

import (
	"net/url"
	"quince/modules/admin/services"
	"strings"
)

// EditorController struct
type EditorController struct {
	BaseController
}

// Prepare
func (ec *EditorController) Prepare() {
	//Cancel _xsrf verification
	ec.EnableXSRF = false
}

// Server
func (ec *EditorController) Server() {
	var result map[string]interface{}
	ueditorService := services.NewUeditorService()
	action := ec.GetString("action")
	switch action {
	case "config":
		result = ueditorService.GetConfig()
	case "uploadimage":
		result = ueditorService.UploadImage(ec.Ctx)
	case "uploadscrawl":
		//Upload graffiti
		//If the value with + sign is not processed, it will be converted to a space
		ec.Ctx.Request.URL.RawQuery = strings.ReplaceAll(ec.Ctx.Request.URL.RawQuery, "+", "%2b")
		values, _ := url.ParseQuery(ec.Ctx.Request.URL.RawQuery)
		result = ueditorService.UploadScrawl(values)
	case "uploadvideo":
		result = ueditorService.UploadVideo(ec.Ctx)
	case "uploadfile":
		result = ueditorService.UploadFile(ec.Ctx)
	case "listimage":
		values, _ := ec.Input()
		result = ueditorService.ListImage(values)
	case "listfile":
		values, _ := ec.Input()
		result = ueditorService.ListFile(values)
	case "catchimage":
		result = ueditorService.CatchImage(ec.Ctx)
	default:
		result = map[string]interface{}{
			"state": ec.Translate("error.address_request"),
		}
	}
	ec.Data["json"] = result
	ec.ServeJSON()

}
