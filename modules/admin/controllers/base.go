package controllers

import (
	"errors"
	"fmt"
	"net/url"
	"quince/global"
	"quince/global/response"
	im "quince/internal/models"
	c "quince/internal/toolbar/component"
	"quince/internal/transaction"
	"quince/internal/validation"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"quince/utils"
	"strconv"
	"strings"
	"time"

	"quince/internal/i18n"
	"quince/middleware"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)

// NestPreparer Define the sub-controller initialization method
type NestPreparer interface {
	NestPrepare()
}

//public
// type BaseController struct {
// 	baseController
// }

// override "quince/global/response" to add tranlation
type globalResponse struct {
}

// baseController struct
type BaseController struct {
	web.Controller
	Log           *logs.BeeLogger
	i18n          i18n.Locale
	Buttons       []c.IButton
	Transaction   *transaction.Batch
	Response      *globalResponse
	DataSearchBox map[string]im.IModel //contain models for search box html component
}

var (
	//Background variables
	admin map[string]interface{}
	//Current user
	loginUser models.LoginUser
	//parameter
	gQueryParams url.Values
	//Base layout for all views
	//to override base layout just add myController.Layout = "whatever/new_layout.html" in controller function
	base_layout        = "admin/views/public/base.html"
	modal_layer_layout = "admin/views/public/modal-layer.html"
)

func (bc *BaseController) NestPrepare() {

}

// Prepare Parent controller initialization
func (bc *BaseController) Prepare() {

	//Init Logger
	bc.Log = logs.NewLogger()
	f, _ := web.AppConfig.String("log::controller")
	bc.Log.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(`{"filename":"log/%s","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true,"rotate":true}`, f))
	//Access url
	requestURL := strings.ToLower(strings.TrimLeft(bc.Ctx.Input.URL(), "/"))
	if !strings.HasPrefix(requestURL, "https") { //XSRF works only with HTTPS
		bc.EnableXSRF = false
		logs.Info("HTTP only protocol . XSRF disabled .")
	}
	//init data searchbox
	bc.DataSearchBox = make(map[string]im.IModel)

	//OK Message
	succes := web.NewFlash()
	succes.Success("  ") //leave empty
	succes.Store(&bc.Controller)

	//query parameter
	//Only used when the first page list is paged
	if bc.Ctx.Input.IsGet() {
		gQueryParams, _ = url.ParseQuery(bc.Ctx.Request.URL.RawQuery)
		gQueryParams.Set("queryParamUrl", bc.Ctx.Input.URL())
		if len(gQueryParams) > 0 {
			for k, val := range gQueryParams {
				v, ok := strconv.Atoi(val[0])
				if ok == nil {
					bc.Data[k] = v
				} else {
					bc.Data[k] = val[0]
				}
			}
		}
	}

	//Login user
	var isOk bool
	loginUser, isOk = bc.GetSession(global.LOGIN_USER).(models.LoginUser)
	bc.i18n.Lang = loginUser.Language
	//add i18n template funcmap {{i18n "bl.bla"}}
	web.AddFuncMap("i18n", bc.Translate)
	//Init transactional queue list
	bc.Transaction = transaction.NewTransaction(loginUser.LoginName, loginUser.Id)

	//Basic variable
	runMode, _ := web.AppConfig.String("runmode")
	if runMode == "dev" {
		bc.Data["debug"] = true
	} else {
		bc.Data["debug"] = false
	}
	bc.Data["cookie_prefix"] = ""

	//Number of previews per page
	perPageStr := bc.Ctx.GetCookie("admin_per_page")
	var perPage int
	if perPageStr == "" {
		perPage = 10
	} else {
		perPage, _ = strconv.Atoi(perPageStr)
	}
	if perPage >= 100 {
		perPage = 100
	}

	//Log
	adminMenuService := services.NewAdminMenuService()
	adminMenu := adminMenuService.GetAdminMenuByUrl(requestURL)
	title := ""
	if adminMenu != nil {
		title = adminMenu.Name
		if strings.EqualFold(adminMenu.LogMethod, bc.Ctx.Input.Method()) {
			adminLogService := services.NewAdminLogService()
			adminLogService.CreateAdminLog(&loginUser, adminMenu, requestURL, bc.Ctx)
		}
	}
	//"request_type" == "layer_open"
	request := gQueryParams.Get("request_type")
	isModal := false
	if len(request) > 0 {
		if request == "layer_open" { //modal open
			isModal = true
		}
	}
	//Left menu
	//menu := "" //DEPRECATED
	if (requestURL != "admin/auth/login") && !isModal && !(bc.Ctx.Input.Header("X-PJAX") == "true") && isOk {
		//var adminTreeService services.AdminTreeService//DEPRECATED
		//menu = adminTreeService.GetLeftMenu(requestURL, loginUser, bc.i18n.Lang)//DEPRECATED
		// add base layout for any other page except login and ajax requests
		bc.Layout = base_layout
	}
	if isModal {
		bc.Layout = modal_layer_layout // load JS and CSS files
	}
	currentTime := time.Now()
	admin = map[string]interface{}{
		"pjax": bc.Ctx.Input.Header("X-PJAX") == "true",
		"user": &loginUser,
		//"menu":            menu, //deprecated
		"name":            global.BA_CONFIG.Base.Name,
		"author":          global.BA_CONFIG.Base.Author,
		"version":         global.BA_CONFIG.Base.Version,
		"short_name":      global.BA_CONFIG.Base.ShortName,
		"link":            global.BA_CONFIG.Base.Link,
		"per_page":        perPage,
		"per_page_config": []int{10, 20, 30, 50, 100},
		"title":           title,
		"date":            currentTime.Format("2006-Jan-02"),
		"jsLang":          i18n.JavascriptLanguage(loginUser.Language),
	}
	bc.Data["admin"] = admin
	if loginUser.Language != "" {
		bc.i18n.Lang = loginUser.Language
	} else {
		//set language to default
		// langs, _ := web.AppConfig.String("lang::alpha4")
		// lang := strings.Split(langs, "|")
		langs := i18n.ApplicationLanguages()
		bc.i18n.Lang = langs[0]
	}

	//Ajax header unified setting _xsrf
	bc.Data["xsrf_token"] = bc.XSRFToken()

	//Determine whether the sub-controller has an initialization method
	if app, ok := bc.AppController.(NestPreparer); ok {
		app.NestPrepare() //execute init controller's method

	}

	//workaround to have flash messages on redirect
	//see global/response/response.go -> func Result(...)
	beegoFlash := bc.Ctx.GetCookie("BEEGO_FLASH")
	beegoFlash = strings.ReplaceAll(beegoFlash, "%23", "")
	beegoFlash = strings.ReplaceAll(beegoFlash, "%00", "")
	beegoFlash = strings.ReplaceAll(beegoFlash, "+", " ")
	beegoFlash = strings.TrimSpace(beegoFlash)
	v := strings.Split(beegoFlash, web.BConfig.WebConfig.FlashSeparator)
	if len(v) == 2 {
		fl := web.NewFlash()
		if v[0] == "error" {
			fl.Error(v[1])
		} else {
			fl.Notice(v[1])
		}
		fl.Store(&bc.Controller)
		bc.Ctx.SetCookie("BEEGO_FLASH", "", -1, "/")
	}
}
func (bc *BaseController) GetAdminMap(key string) interface{} {
	return admin[key]
}
func (bc *BaseController) GetQueryParams() *url.Values {
	return &gQueryParams
}
func (bc *BaseController) GetSelectedIDs() []int64 {
	idStr := bc.GetString("id")
	ids := make([]int64, 0)
	var idArr []int64

	if idStr == "" {
		bc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		bc.ResponseErrorWithMessage(errors.New("admin.select_user"), bc.Ctx)
	}

	return idArr
}

// before render add variables
func (bc *BaseController) Render() error {
	//add buttons to template
	bc.Data["buttons"] = bc.Buttons
	//add search boxes data to template
	for key, item := range bc.DataSearchBox {
		header := strings.Join(bc.translateStruct(item), "|")
		fields := strings.Join(item.SearchField(), "|")
		value := fmt.Sprintf("%s\t%s", header, fields)
		bc.Data[key] = value
		fmt.Printf("%s - %s", key, value)
	}

	return bc.Controller.Render()
}

// on router finished - is executed after html response
func (bc *BaseController) Finish() {

}
func (bc *BaseController) AddButtons(params ...c.IButton) {
	for _, t := range params {
		//t.Lang = bc.i18n.Lang // attach language
		bc.checkToolbarAccess(t) // check access rights
		t.Translate(bc.i18n.Lang)
		bc.Buttons = append(bc.Buttons, t)
	}

}

func (bc *BaseController) checkToolbarAccess(p c.IButton) {
	adminUserService := services.NewAdminUserService()
	authExcept := middleware.AuthException()
	loginUser, _ = bc.GetSession(global.LOGIN_USER).(models.LoginUser)

	path := p.GetUrl()
	baseUrl, err := url.Parse(path)
	if err != nil {
		logs.Warn("Malformed button Url %s", baseUrl.Path)
		return
	}
	urlCheck := adminUserService.UrlCheck(baseUrl.Path, authExcept, &loginUser)
	if !urlCheck {
		p.SetDeniedAccess()
	}

}

// func (bc *BaseController) ToolbarButtons(params ...toolbar.IButton) {
// 	adminUserService := services.NewAdminUserService()
// 	authExcept := middleware.AuthException()
// 	loginUser, _ = bc.GetSession(global.LOGIN_USER).(models.LoginUser)

// 	for _, p := range params {
// 		path := p.GetUrl()
// 		baseUrl, err := url.Parse(path)
// 		if err != nil {
// 			logs.Warn("Malformed button Url %s", baseUrl.Path)
// 			return
// 		}
// 		urlCheck := adminUserService.UrlCheck(baseUrl.Path, authExcept, &loginUser)
// 		if urlCheck {
// 			switch p.GetToolbar() {
// 			case toolbar.Top:
// 				bc.Toolbar.Top.Buttons = append(bc.Toolbar.Top.Buttons, p)
// 			case toolbar.TopMultiSelect:
// 				bc.Toolbar.Top.MultiSelect = append(bc.Toolbar.Top.MultiSelect, p)
// 			case toolbar.Form:
// 				bc.Toolbar.Form.Buttons = append(bc.Toolbar.Form.Buttons, p)
// 			case toolbar.FormMultiSelect:
// 				bc.Toolbar.Form.MultiSelect = append(bc.Toolbar.Form.MultiSelect, p)
// 			case toolbar.GridRow:
// 				bc.Toolbar.GridRow.Buttons = append(bc.Toolbar.GridRow.Buttons, p)
// 			case toolbar.GridRowMultiSelect:
// 				bc.Toolbar.GridRow.MultiSelect = append(bc.Toolbar.GridRow.MultiSelect, p)
// 			case toolbar.Independent:
// 				bc.Toolbar.Auth.Buttons = append(bc.Toolbar.Auth.Buttons, p)
// 			default:
// 				logs.Warn("ToolbarButtons baseController:266 unknown type %v ", p)
// 			}
// 		}
// 	}
// }

// Translate - overwrite I18n.Tr from beego plugin
func (bc *BaseController) Translate(format string, args ...interface{}) string {
	var item string
	b := true
	bracketsFound := false
	br := false
	for b { //check for words in curly brackets {{}}
		item, br = utils.GetStringInBetweenTwoString(format, "{{", "}}")
		if len(item) > 0 {
			t := bc.i18n.Tr(item, args)
			token := fmt.Sprintf("{{%s}}", item)
			format = strings.Replace(format, token, t, -1) //replace with translated text
			bracketsFound = true
		}
		b = br
	}
	if bracketsFound {
		return format
	} else { //search for translation
		return bc.i18n.Tr(format, args)
	}
}

// Array translate string array  to target language.
// format param could be :
// - []string
// - a pointer to a struct IModel interface with i18n tags as columns names
// func (bc *BaseController) translateArray(values []string, args ...interface{}) []string {
// 	var result []string
// 	for _, v := range values {
// 		result = append(result, bc.Translate(v, args...))
// 	}
// 	return result

// }
func (bc *BaseController) translateErrorMap(err error, args ...interface{}) error {
	result := validation.Errors{}
	values := err.(validation.Errors)
	for k, v := range values {
		result[bc.Translate(k, args...)] = errors.New(bc.Translate(v.Error(), args...))
	}
	return result

}
func (bc *BaseController) translateStruct(format interface{}, args ...interface{}) []string {
	return bc.i18n.Struct(format, args)
}

// func (bc *BaseController) ExecuteTransaction() error {
// 	o := orm.NewOrm()
// 	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {

// 		worker := transaction.NewWorker(bc.Transaction.List)
// 		b := worker.DoWork(txOrm)
// 		if !b {
// 			return errors.New("transaction failed ")
// 		}
// 		return nil
// 	})
// 	return err
// }

// response area
// Success, normal return
func (bc *BaseController) ResponseSuccess(ctx *context2.Context) {
	response.Success(ctx)
}

// SuccessWithMessage  return custom information
func (bc *BaseController) ResponseSuccessWithMessage(msg string, ctx *context2.Context) {
	m := bc.Translate(msg)
	response.SuccessWithMessage(m, ctx)

}

// SuccessWithMessageAndUrl Success, return custom information and url
func (bc *BaseController) ResponseSuccessWithMessageAndUrl(msg string, url string, ctx *context2.Context) {
	m := bc.Translate(msg)
	response.SuccessWithMessageAndUrl(m, url, ctx)
}

// SuccessWithDetailed  return all custom information
func (bc *BaseController) ResponseSuccessWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *context2.Context) {
	m := bc.Translate(msg)
	response.SuccessWithDetailed(m, url, data, wait, header, ctx)
}

// SuccessWithDetailed  return all custom information
func (bc *BaseController) ResponseSuccessWithData(msg string, data interface{}, ctx *context2.Context) {
	m := bc.Translate(msg)
	response.SuccessWithData(m, data, ctx)
}

// Error Failure, normal return
func (bc *BaseController) ResponseError(ctx *context2.Context) {
	response.Error(ctx)
}

// ErrorWithMessage Failure, return custom information
func (bc *BaseController) ResponseErrorWithMessage(err error, ctx *context2.Context) {
	var msg string
	v, ok := err.(validation.Errors) //if type assertion fail =ok=false
	if ok {
		m := bc.translateErrorMap(v)
		msg = m.Error()
	} else {
		msg = bc.Translate(err.Error())
	}

	response.ErrorWithMessage(msg, ctx)
}

// ErrorWithMessageAndUrl Failure, return custom information and url
func (bc *BaseController) ResponseErrorWithMessageAndUrl(err error, url string, ctx *context2.Context) {
	var msg string
	v, ok := err.(validation.Errors) //if type assertion fail =ok=false
	if ok {                          //is validation error type
		m := bc.translateErrorMap(v)
		msg = m.Error()
	} else {
		msg = bc.Translate(err.Error())
	}
	response.ErrorWithMessageAndUrl(msg, url, ctx)
}

// ErrorWithDetailed Fail, return all custom information
func (bc *BaseController) ResponseErrorWithDetailed(err error, url string, data interface{}, wait int, header map[string]string, ctx *context2.Context) {
	var msg string
	v, ok := err.(validation.Errors) //if type assertion fail =ok=false
	if ok {
		m := bc.translateErrorMap(v)
		msg = m.Error()
	} else {
		msg = bc.Translate(err.Error())
	}
	response.ErrorWithDetailed(msg, url, data, wait, header, ctx)
}
