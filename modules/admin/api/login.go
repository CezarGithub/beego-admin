package api

import (
	"errors"
	"quince/internal/jwt"
	"quince/modules/admin/controllers"
	"quince/modules/admin/form"
	"quince/modules/admin/services"

	"github.com/beego/beego/v2/core/logs"
)

type ApiController struct {
	controllers.BaseController
}

// CheckLogin Login authentication
// @params u , p , m = user ,password . MAC address
// Header [X-API-Key] contains JWT key -except for login method
// Header [X-Requested-With] = XMLHttpRequest
// POST method only
func (ac *ApiController) CheckLogin() {
	//Data validation
	user := ac.GetString("u", "none")
	pass := ac.GetString("p", "none")
	macAddr := ac.GetString("m", "")

	loginForm := form.LoginForm{}
	loginForm.Username = user
	loginForm.Password = pass
	loginForm.Captcha = macAddr
	//After the basic verification is passed, perform user verification
	adminUserService := services.NewAdminUserService()
	loginUser, err := adminUserService.CheckLogin(loginForm, ac.Ctx)
	if err != nil {
		logs.Warn("Login attempt : %s IP : %s Err: %v", user, ac.Ctx.Request.RemoteAddr, err.Error()) //log -  file
	} else {
		logs.Info("Login succeed : %s IP : %s", loginUser.LoginName, ac.Ctx.Request.RemoteAddr)
	}

	if err != nil {
		e := errors.New("access_denied")
		ac.ResponseErrorWithMessage(e, ac.Ctx)
		return
	}

	//Login logging
	adminLogService := services.NewAdminLogService()
	adminLogService.LoginLog(loginUser.Id, ac.Ctx) //database
	//ac.I18n.Lang = loginUser.Language
	logs.Info("API logged user %s lang:%s : %s ", loginUser.LoginName, loginUser.Language, ac.Translate("login.hi"))
	//ac.ResponseSuccess(ac.Ctx)

	msg := "OK"
	url := ac.Ctx.Request.URL
	data, _ := jwt.GenerateToken(loginUser.LoginName, loginUser.Id,loginUser.Roles) // "JWT_TOKEN"
	header := make(map[string]string)
	header["Authorization"] = "X-API-Key"
	ac.ResponseSuccessWithDetailed(msg, url.Host, data, 0, header, ac.Ctx)

}
