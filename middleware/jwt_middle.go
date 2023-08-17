package middleware

import (
	"quince/global/response"
	"quince/internal/jwt"

	"github.com/beego/beego/v2/core/logs"
	context "github.com/beego/beego/v2/server/web/context"
)

var JWT = func(ctx *context.Context) *jwt.Claims {
	token := ctx.Request.Header.Get("X-API-Key")
	logs.Info("JWT request token : %s", token)
	if token == "" {
		logs.Info("JWT token missing")
		response.ErrorWithMessage("access_denied", ctx)
		//b = false
	} else {
		claims, err := jwt.ParseToken(token)
		if err != nil {
			info := err.Error()
			logs.Info(info)
			response.ErrorWithMessage("access_denied", ctx)
		} else {
			logs.Info("JWT claims %v", claims)
			return claims
		}
	}
	return nil
}
