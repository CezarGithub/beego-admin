package controllers

import "github.com/beego/beego/v2/server/web"

// MainController struct
type MainController struct {
	web.Controller
}

// Get
func (c *MainController) Get() {
	c.Data["Website"] = "kodis.ro"
	c.Data["Email"] = "sales@kodis.ro"
	c.TplName = "admin/views/index.tpl"
}
