package controllers

import (
	"quince/modules/admin/controllers"
)

// CrmController struct
type CrmController struct {
	controllers.BaseController
}

func (mc *CrmController) Index() {
	mc.ViewPath = "modules"
	mc.Layout = "admin/views/public/base.html"
	mc.TplName = "crm/views/index.html"
}
