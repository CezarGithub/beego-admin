package controllers

import (
	"errors"
	"quince/internal/toolbar"
	"quince/modules/admin/controllers"
	"quince/modules/master/form"
	"quince/modules/master/models"
	"quince/modules/master/services"
	"quince/modules/master/transactions"

	"github.com/beego/beego/v2/server/web"
)

// CountyController struct
type CountyController struct {
	controllers.BaseController
}

// NestPrepare - Init controller executed first
func (mc *CountyController) NestPrepare() {
	//
}
func (mc *CountyController) Index() {

	countyService := services.NewCountyService()
	name := mc.GetString("name")
	id, _ := mc.GetInt64("country_id", 0)
	country := new(models.Country)
	country.SetID(id)
	county := new(models.County)
	county.Country = country
	countyService.IModel = county

	data, pagination := countyService.GetPaginateData(mc.GetAdminMap("per_page").(int), *mc.GetQueryParams())
	mc.Data["paginate"] = pagination
	mc.Data["data"] = data
	mc.Data["name"] = name
	mc.DataSearchBox["searchBoxCountry"] = &models.Country{}
	add := toolbar.Html("county.add").Add(web.URLFor("CountyController.Add"))
	edit := toolbar.Html("county.edit").Edit(web.URLFor("CountyController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	mc.AddButtons(add,edit)
	mc.TplName = "master/views/county/index.html"
}
func (mc *CountyController) Add() {
	item := new(models.County)
	item.Country = &models.Country{}
	mc.show(item)
}
func (mc *CountyController) Edit() {
	id, _ := mc.GetInt("id", -1)
	countyService := services.NewCountyService()
	item := countyService.GetCountyById(id)
	if item == nil {
		mc.ResponseErrorWithMessageAndUrl(errors.New("error.info_not_found"), web.URLFor("CountyController.Index"), mc.Ctx)
		return
	}
	mc.Data["data"] = item
	mc.DataSearchBox["searchBoxCountry"] = &models.Country{}
	mc.show(item)
}
func (mc *CountyController) show(item *models.County) {
	save := toolbar.Ajax("county.submit").Submit("countyForm", web.URLFor("CountyController.Update"))
	mc.AddButtons(save)
	mc.Data["data"] = item
	mc.TplName = "master/views/county/edit.html"
}
func (mc *CountyController) Search() {
	countyService := services.NewCountyService()
	data := countyService.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}
func (mc *CountyController) Update() {
	countyForm := form.CountyForm{}
	if err := mc.ParseForm(&countyForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := countyForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
	t1 := transactions.CountyUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
