package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"quince/internal/captcha"
)

// CaptchaController struct
type CaptchaController struct {
	web.Controller
}

// CaptchaId 获取CaptchaId
func (cc *CaptchaController) CaptchaId() {
	captchaID := captcha.NewLen(6)
	cc.Ctx.WriteString(captchaID)
}
