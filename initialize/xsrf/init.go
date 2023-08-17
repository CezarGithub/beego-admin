package xsrf

import (
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// XSRF works with HTTPS protocol. Check admin/controller/base.go:76
func init() {
	xsrfEnable := true
	xsrfKey := web.AppConfig.DefaultString("xsrf::key", "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o")
	ex := web.AppConfig.DefaultString("xsrf::expire", "3600")
	b := web.AppConfig.DefaultString("xsrf::enable", "true")
	expire, err := strconv.Atoi(ex)
	if b != "true" {
		xsrfEnable = false
	}
	if err != nil {
		expire = 3600
	}
	logs.Info("XSRF enabled : %v", xsrfEnable)
	web.BConfig.WebConfig.EnableXSRF = xsrfEnable
	web.BConfig.WebConfig.XSRFKey = xsrfKey
	web.BConfig.WebConfig.XSRFExpire = expire
}
