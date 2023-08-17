package controllers

import (
	"quince/modules/admin/controllers"
)

// MasterController struct
type MasterController struct {
	controllers.BaseController
}

func (mc *MasterController) Index() {
	mc.ViewPath = "modules"
	mc.Layout = "admin/views/public/base.html"
	mc.TplName = "master/views/index.html"
}
