package controllers

import (
	"errors"
	"net/http"
	"os"
	"quince/global"
	"quince/modules/admin/form"
	"quince/modules/admin/services"
	"quince/utils"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	"quince/internal/captcha"
)

// AuthController struct
type AuthController struct {
	BaseController
}

// Login interface
func (ac *AuthController) Login() {
	//Load login configuration information
	var settingService services.SettingService
	data := settingService.Show(1, ac.i18n.Lang)
	for _, setting := range data {
		settingService.LoadOrUpdateGlobalBaseConfig(setting)
	}

	//Get login configuration information
	loginConfig := struct {
		Token      string
		Captcha    string
		Background string
	}{
		Token:   global.BA_CONFIG.Login.Token,
		Captcha: global.BA_CONFIG.Login.Captcha,
	}
	//Login background image
	if _, err := os.Stat(strings.TrimLeft(global.BA_CONFIG.Login.Background, "/")); err != nil {
		global.BA_CONFIG.Login.Background = "/static/admin/images/login-default-bg.jpg"
	}
	loginConfig.Background = global.BA_CONFIG.Login.Background

	var password string
	var username string
	runMode, _ := web.AppConfig.String("runmode")
	if runMode == "dev" { //fill login fields in dev mode
		password = "super_admin"
		username = "super_admin"
	}
	//The login interface only needs the name field
	admin := map[string]interface{}{
		"name":     global.BA_CONFIG.Base.Name,
		"title":    "Login",
		"password": password,
		"username": username,
	}

	ac.Data["login_config"] = loginConfig
	//Login Verification Code
	ac.Data["captcha"] = utils.GetCaptcha()
	ac.Data["admin"] = admin

	ac.TplName = "admin/views/auth/login.html"
}

// Logout sign out
func (ac *AuthController) Logout() {
	ac.DelSession(global.LOGIN_USER)
	ac.Ctx.SetCookie(global.LOGIN_USER_ID, "", -1)
	ac.Ctx.SetCookie(global.LOGIN_USER_ID_SIGN, "", -1)
	ac.Redirect("/admin/auth/login", http.StatusFound)
}

// CheckLogin Login authentication
func (ac *AuthController) CheckLogin() {
	//Data validation

	loginForm := form.LoginForm{}

	if err := ac.ParseForm(&loginForm); err != nil {
		ac.ResponseErrorWithMessage(err, ac.Ctx)
	}

	//See if you need to verify the verification code
	isCaptcha, _ := strconv.Atoi(global.BA_CONFIG.Login.Captcha)
	if isCaptcha > 0 {
		//valid.Required(loginForm.Captcha, "captcha").Message(ac.Translate("login.code_please"))
		if ok := captcha.VerifyString(loginForm.CaptchaId, loginForm.Captcha); !ok {
			ac.ResponseErrorWithMessage(errors.New("login.code_error"), ac.Ctx)
		}
	}

	//After the basic verification is passed, perform user verification
	adminUserService := services.NewAdminUserService()
	loginUser, err := adminUserService.CheckLogin(loginForm, ac.Ctx)
	if err != nil {
		logs.Warn("Login attempt : %s IP : %s Err: %v", loginForm.Username, ac.Ctx.Request.RemoteAddr, err.Error()) //log -  file
	} else {
		logs.Info("Login succeed : %s IP : %s", loginUser.LoginName, ac.Ctx.Request.RemoteAddr)
	}

	if err != nil {
		ac.ResponseErrorWithMessage(err, ac.Ctx)
		return
	}

	//Login logging
	adminLogService := services.NewAdminLogService()
	adminLogService.LoginLog(loginUser.Id, ac.Ctx) //database
	//ac.I18n.Lang = loginUser.Language
	ac.Log.Info("Logged user %s lang:%s : %s ", loginUser.LoginName, loginUser.Language, ac.Translate("login.hi"))

	redirect, _ := ac.GetSession("redirect").(string)
	if redirect != "" {
		ac.ResponseSuccessWithMessageAndUrl("login.success", redirect, ac.Ctx)
	} else {
		ac.ResponseSuccessWithMessageAndUrl("login.success", "/admin/index/index", ac.Ctx)
	}
}

// RefreshCaptcha Refresh Code
func (ac *AuthController) RefreshCaptcha() {
	captchaID := ac.GetString("captchaId")
	res := map[string]interface{}{
		"isNew": false,
	}
	if captchaID == "" {
		res["msg"] = ac.Translate("error.param_error")
	}

	isReload := captcha.Reload(captchaID)
	if isReload {
		res["captchaId"] = captchaID
	} else {
		res["isNew"] = true
		res["captcha"] = utils.GetCaptcha()
	}

	ac.Data["json"] = res

	ac.ServeJSON()
}
