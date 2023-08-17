package response

import (
	"net/http"
	uri "net/url"
	"quince/global"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

// Response Response parameter structure
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Url  string      `json:"url"`
	Wait int         `json:"wait"`
}

// Result Return result auxiliary function
func result(code int, msg string, data interface{}, url string, wait int, header map[string]string, ctx *context.Context) {
	if ctx.Input.IsPost() && ctx.Input.Context.Input.Header("X-Requested-With") == "XMLHttpRequest" { //.Request.Header["X-Requested-With"] == "XMLHttpRequest"
		if (ctx.Input.Query("request_type") == "layer_open") && url == global.URL_BACK { //is request from modal
			jsonResult(code, msg, data, global.URL_CLOSE_REFRESH, wait, header, ctx) //close popup and refresh parent
		} else {
			jsonResult(code, msg, data, url, wait, header, ctx) //ajax request
		}
	} else {
		flashType := "error"
		if code == 1 {
			flashType = "success"
		}
		flashValue := "\x00" + flashType + "\x23" + web.BConfig.WebConfig.FlashSeparator + "\x23" + msg + "\x00"
		ctx.SetCookie(web.BConfig.WebConfig.FlashName, uri.QueryEscape(flashValue))
		if url == "" {
			url = ctx.Request.Referer()
			if url == "" {
				url = "/admin/index/index"
			}
		}

		ctx.Redirect(http.StatusFound, url)
	}
}
func jsonResult(code int, msg string, data interface{}, url string, wait int, header map[string]string, ctx *context.Context) {
	if ctx.Input.IsPost() {
		result := Response{
			Code: code,
			Msg:  msg,
			Data: data,
			Url:  url,
			Wait: wait,
		}

		if len(header) > 0 {
			for k, v := range header {
				ctx.Output.Header(k, v)
			}
		}

		ctx.Output.JSON(result, false, false)

		//Controller Usage in this.StopRun()
		panic(web.ErrAbort)
	}

	if url == "" {
		url = ctx.Request.Referer()
		if url == "" {
			url = "/admin/index/index"
		}
	}

	ctx.Redirect(http.StatusFound, url)
}

// Success, normal return
func Success(ctx *context.Context) {
	result(SUCCESS, "OK", "", global.URL_BACK, 0, map[string]string{}, ctx)
}

// SuccessWithMessage  return custom information
func SuccessWithMessage(msg string, ctx *context.Context) {
	result(SUCCESS, msg, "", global.URL_BACK, 0, map[string]string{}, ctx)
}

// SuccessWithMessage  return custom information
func SuccessWithData(msg string, data interface{}, ctx *context.Context) {
	//d,_:=json.Marshal(data)
	result(SUCCESS, msg, data, global.URL_BACK, 0, map[string]string{}, ctx)
}

// SuccessWithMessageAndUrl Success, return custom information and url
func SuccessWithMessageAndUrl(msg string, url string, ctx *context.Context) {
	result(SUCCESS, msg, "", url, 0, map[string]string{}, ctx)
}

// SuccessWithDetailed  return all custom information
func SuccessWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *context.Context) {
	result(SUCCESS, msg, data, url, wait, header, ctx)
}

// Error Failure, normal return
func Error(ctx *context.Context) {
	logs.Error("Error : %v", global.URL_CURRENT)
	result(ERROR, "Error", "", global.URL_CURRENT, 0, map[string]string{}, ctx)
}

// ErrorWithMessage Failure, return custom information
func ErrorWithMessage(msg string, ctx *context.Context) {
	logs.Error("Error : %s - %v", msg, global.URL_CURRENT)
	result(ERROR, msg, "", global.URL_CURRENT, 0, map[string]string{}, ctx)
}

// ErrorWithMessageAndUrl Failure, return custom information and url
func ErrorWithMessageAndUrl(msg string, url string, ctx *context.Context) {
	logs.Error("Error : %s - %s", msg, url)
	result(ERROR, msg, "", url, 0, map[string]string{}, ctx)
}

// ErrorWithDetailed Fail, return all custom information
func ErrorWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *context.Context) {
	logs.Error("Error : %s - %v - %s", msg, data, url)
	result(ERROR, msg, data, url, wait, header, ctx)
}
