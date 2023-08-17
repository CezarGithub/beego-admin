package api

import (
	"quince/modules/admin/controllers"
	"quince/modules/master/services"
)

type ApiVersionOneController struct {
	controllers.BaseController
}

func (avo ApiVersionOneController) Countries() {
	countryService := services.NewCountryService()
	data := countryService.GetAll(*avo.GetQueryParams())
	avo.ResponseSuccessWithData("Countries list", data, avo.Ctx)
}
