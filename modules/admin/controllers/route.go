package controllers

import (
	"quince/modules/admin/services"
)

// AdminMenuController struct.
type RouteController struct {
	BaseController
}

// Index Menu home.
func (amc *RouteController) Index() {
	routesService:= services.NewRoutesService()
	data, pagination := routesService.GetPaginatedData(admin["per_page"].(int), gQueryParams)
	//log.Printf("Params %v \n", gQueryParams)
	amc.Data["method_list"] = routesService.Methods()

	amc.Data["data"] = data
	amc.Data["paginate"] = pagination

	amc.TplName = "admin/views/route/index.html"
}

// Index Menu home.
func (amc *RouteController) View() {

	amc.TplName = "admin/views/route/view.html"
	routesService:= services.NewRoutesService()
	data, _ := routesService.GetPaginatedData(admin["per_page"].(int), gQueryParams)
	//log.Printf("Params %v \n", gQueryParams)

	amc.Data["data"] = data

}
