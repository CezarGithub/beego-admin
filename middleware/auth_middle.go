package middleware

import (
	"fmt"
	"quince/global"
	"quince/global/response"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"quince/utils"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	context2 "github.com/beego/beego/v2/server/web/context"
)

// Provide no need to check url's
func AuthException() map[string]interface{} {
	//No verification required
	authExcept := map[string]interface{}{
		"admin/auth/login":           0,
		"admin/auth/check_login":     1,
		"admin/auth/logout":          2,
		"admin/auth/captcha":         3,
		"admin/editor/server":        4,
		"admin/auth/refresh_captcha": 5,
		"admin/api/login":            6,
	}
	return authExcept
}

// AuthMiddle Middleware - n
func AuthMiddle() {

	authExcept := AuthException()
	//Login authentication middleware filter
	var filterLogin = func(ctx *context2.Context) {
		url := strings.TrimLeft(ctx.Input.URL(), "/")

		//JWT zone
		token := ctx.Input.Header("Authorization")
		if token != "" { //we have a JWT auth request
			if token != "X-API-Key" { //malformed request
				logs.Warn("Authorization header malformed %s - %s", url, token)
				response.ErrorWithMessage("access_denied", ctx)
				return
			}
			claims := JWT(ctx)
			adminUserService := services.NewAdminUserService()
			urlCheck := adminUserService.APIUrlCheck(url, claims.UserId, claims.Role)
			if !urlCheck {
				response.ErrorWithMessage("access_denied", ctx)
			}
			return
		}

		//Login verification required
		if !isAuthExceptUrl(strings.ToLower(url), authExcept) {
			//Verify that you are logged in
			loginUser, isLogin := isLogin(ctx)
			if !isLogin {
				response.ErrorWithMessageAndUrl("Not logged in ", "/admin/auth/login", (*context2.Context)(ctx))
				return
			}

			if !utils.KeyInMap(url, authExcept) { //check if is excepted route
				//Verify, whether you have permission to access
				adminUserService := services.NewAdminUserService()
				menuCheck := adminUserService.MenuCheck(url, authExcept, loginUser)
				urlCheck := adminUserService.UrlCheck(url, authExcept, loginUser)
				if loginUser.Id != 1 && !menuCheck && !urlCheck {
					errorBackURL := global.URL_CURRENT
					if ctx.Request.Method == "GET" {
						errorBackURL = ""
					}
					response.ErrorWithMessageAndUrl("admin.permission_required", errorBackURL, (*context2.Context)(ctx))
					return
				}
			}
		}

		checkAuth, _ := strconv.Atoi(ctx.Request.PostForm.Get("check_auth"))

		if checkAuth == 1 {
			response.Success((*context2.Context)(ctx))
			return
		}

	}

	//beego.InsertFilter("/admin/*", beego.BeforeRouter, filterLogin)
	web.InsertFilter("/*", web.BeforeRouter, filterLogin)
}

// Determine whether authentication is not required to log in
func isAuthExceptUrl(url string, m map[string]interface{}) bool {
	urlArr := strings.Split(url, "/")
	if len(urlArr) > 3 {
		url = fmt.Sprintf("%s/%s/%s", urlArr[0], urlArr[1], urlArr[2])
	}
	_, ok := m[url]
	return ok
}

// Whether to log in
func isLogin(ctx *context2.Context) (*models.LoginUser, bool) {
	loginUser, ok := ctx.Input.Session(global.LOGIN_USER).(models.LoginUser)
	if !ok {
		loginUserIDStr := ctx.GetCookie(global.LOGIN_USER_ID)
		loginUserIDSign := ctx.GetCookie(global.LOGIN_USER_ID_SIGN)

		if loginUserIDStr != "" && loginUserIDSign != "" {
			loginUserID, _ := strconv.ParseInt(loginUserIDStr, 10, 64)
			adminUserService := services.NewAdminUserService()
			loginUserPointer := adminUserService.GetAdminUserById(loginUserID)

			if loginUserPointer != nil && loginUserPointer.GetSignStrByAdminUser((*context2.Context)(ctx)) == loginUserIDSign {
				ctx.Output.Session(global.LOGIN_USER, *loginUserPointer)
				return loginUserPointer, true
			}
		}
		return nil, false
	}

	return &loginUser, true
}
