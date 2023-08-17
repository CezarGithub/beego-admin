package routers

import (
	"quince/modules/crm/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	myModule := web.NewNamespace("/crm",
		//index page
		web.NSRouter("/index", &controllers.CrmController{}, "get:Index"),
	)
	web.AddNamespace(myModule)
}
