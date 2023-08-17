package session

import "github.com/beego/beego/v2/server/web"

func init() {
	//session Expiration time, the default value is 3600 seconds
	web.BConfig.WebConfig.Session.SessionGCMaxLifetime = 7200

	//session The time when the client’s cookie is stored by default, the default value is 3600 seconds
	web.BConfig.WebConfig.Session.SessionCookieLifeTime = 7200
}
