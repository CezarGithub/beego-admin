package controllers

import (
	"errors"
	"quince/modules/admin/controllers"

	"quince/internal/toolbar"
	"quince/modules/master/form"
	"quince/modules/master/models"
	"quince/modules/master/services"
	"quince/modules/master/transactions"

	"github.com/beego/beego/v2/server/web"
)

// CountryController struct
type CountryController struct {
	controllers.BaseController
}

func (mc *CountryController) Index() {
	countryService := services.NewCountryService()
	data, pagination := countryService.GetPaginateData(mc.GetAdminMap("per_page").(int), *mc.GetQueryParams())
	mc.Data["paginate"] = pagination
	mc.Data["data"] = data
	add := toolbar.Html("country.add").Add(web.URLFor("CountryController.Add"))
	edit := toolbar.Html("country.edit").Edit(web.URLFor("CountryController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	mc.AddButtons(add, edit)
	mc.TplName = "master/views/country/index.html"
}
func (mc *CountryController) Add() {
	item := new(models.Country)
	mc.show(item)
}
func (mc *CountryController) Edit() {
	id, _ := mc.GetInt("id", -1)
	countryService := services.NewCountryService()
	item := countryService.GetCountryById(id)
	if item == nil {
		mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	mc.show(item)

}
func (mc *CountryController) show(item *models.Country) {
	save := toolbar.Ajax("country.submit").Submit("countryForm", web.URLFor("CountryController.Update"))
	mc.AddButtons(save)
	mc.Data["data"] = item
	mc.TplName = "master/views/country/edit.html"
}
func (mc *CountryController) Search() {
	countryService := services.NewCountryService()
	data := countryService.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}

func (mc *CountryController) Update() {
	countryForm := form.CountryForm{}
	if err := mc.ParseForm(&countryForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := countryForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
		//mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	t1 := transactions.CountryUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
