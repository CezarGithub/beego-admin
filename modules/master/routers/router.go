package routers

import (
	"quince/initialize/router"
	"quince/modules/master/api"
	"quince/modules/master/controllers"
)

func init() {

	masterModule := router.NewNamespace("master",
		router.NSRouter("app.route.index", "OFF", "/index", &controllers.MasterController{}, "get:Index"),
		router.NSRouter("app.route.index", "OFF", "/company/index", &controllers.CompanyController{}, "get:Index"),
		router.NSRouter("app.route.edit", "OFF", "/company/edit", &controllers.CompanyController{}, "get:Edit"),
		router.NSRouter("app.route.add", "OFF", "/company/add", &controllers.CompanyController{}, "get:Add"),
		router.NSRouter("app.route.search", "OFF", "/company/search", &controllers.CompanyController{}, "get:Search"),
		router.NSRouter("app.route.update", "POST", "/company/update", &controllers.CompanyController{}, "post:Update"),
		//country
		router.NSRouter("app.route.index", "OFF", "/country/index", &controllers.CountryController{}, "get:Index"),
		router.NSRouter("app.route.edit", "OFF", "/country/edit", &controllers.CountryController{}, "get:Edit"),
		router.NSRouter("app.route.add", "OFF", "/country/add", &controllers.CountryController{}, "get:Add"),
		router.NSRouter("app.route.search", "OFF", "/country/search", &controllers.CountryController{}, "get:Search"),
		router.NSRouter("app.route.update", "OFF", "/country/update", &controllers.CountryController{}, "post:Update"),
		//county
		router.NSRouter("app.route.index", "OFF", "/county/index", &controllers.CountyController{}, "get:Index"),
		router.NSRouter("app.route.edit", "OFF", "/county/edit", &controllers.CountyController{}, "get:Edit"),
		router.NSRouter("app.route.add", "OFF", "/county/add", &controllers.CountyController{}, "get:Add"),
		router.NSRouter("app.route.update", "POST", "/county/update", &controllers.CountyController{}, "post:Update"),
		router.NSRouter("app.route.search", "OFF", "/county/search", &controllers.CountyController{}, "get:Search"),
		//API
		router.APIRouter("master.companies.list", router.Version1, "OFF", "/country/list", &api.ApiVersionOneController{}, "post:Countries"),
		//smtp
		router.NSRouter("app.route.index", "OFF", "/smtp/index", &controllers.SMTPController{}, "get:Index"),
		router.NSRouter("app.route.edit", "OFF", "/smtp/edit", &controllers.SMTPController{}, "get:Edit"),
		router.NSRouter("app.route.add", "OFF", "/smtp/add", &controllers.SMTPController{}, "get:Add"),
		router.NSRouter("app.route.update", "POST", "/smtp/update", &controllers.SMTPController{}, "post:Update"),
		//group
		router.NSRouter("app.route.index", "OFF", "/group/index", &controllers.GroupController{}, "get:Index"),
		router.NSRouter("app.route.edit", "OFF", "/group/edit", &controllers.GroupController{}, "get:Edit"),
		router.NSRouter("app.route.add", "OFF", "/group/add", &controllers.GroupController{}, "get:Add"),
		router.NSRouter("app.route.update", "POST", "/group/update", &controllers.GroupController{}, "post:Update"),
		//department - TO DO
		router.NSRouter("app.department.index", "OFF", "/department/index", &controllers.DepartmentController{}, "get:Index"),
		router.NSRouter("app.department.edit", "OFF", "/department/edit", &controllers.DepartmentController{}, "get:Edit"),
		router.NSRouter("app.department.add", "OFF", "/department/add", &controllers.DepartmentController{}, "get:Add"),
		router.NSRouter("app.department.update", "POST", "/department/update", &controllers.DepartmentController{}, "post:Update"),
		//currency
		router.NSRouter("app.currency.index", "OFF", "/currency/index", &controllers.CurrencyController{}, "get:Index"),
		router.NSRouter("app.currency.edit", "OFF", "/currency/edit", &controllers.CurrencyController{}, "get:Edit"),
		router.NSRouter("app.currency.add", "OFF", "/currency/add", &controllers.CurrencyController{}, "get:Add"),
		router.NSRouter("app.currency.add", "OFF", "/currency/delete", &controllers.CurrencyController{}, "post:Delete"),
		router.NSRouter("app.currency.update", "POST", "/currency/update", &controllers.CurrencyController{}, "post:Update"),
		//exchangerate
		router.NSRouter("app.exchangerate.index", "OFF", "/exchange_rate/index", &controllers.ExchangeRateController{}, "get:Index"),
		router.NSRouter("app.exchangerate.edit", "OFF", "/exchange_rate/edit", &controllers.ExchangeRateController{}, "get:Edit"),
		router.NSRouter("app.exchangerate.add", "OFF", "/exchange_rate/add", &controllers.ExchangeRateController{}, "get:Add"),
		router.NSRouter("app.exchangerate.update", "POST", "/exchange_rate/update", &controllers.ExchangeRateController{}, "post:Update"),
	)
	router.AddNamespace(masterModule)

}
