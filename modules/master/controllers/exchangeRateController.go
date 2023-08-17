package controllers

import (
	"errors"
	"quince/internal/toolbar"
	"quince/modules/admin/controllers"
	"quince/modules/master/form"
	"quince/modules/master/models"
	"quince/modules/master/services"
	"quince/modules/master/transactions/exchangeRate"

	"github.com/beego/beego/v2/server/web"
)

// ExchangeRateController struct
type ExchangeRateController struct {
	controllers.BaseController
}

// NestPrepare - Init controller
func (mc *ExchangeRateController) NestPrepare() {}

func (mc *ExchangeRateController) Index() {
	exchange_rateService := services.NewExchangeRateService()
	data := exchange_rateService.GetAll(*mc.GetQueryParams())
	mc.Data["data"] = data

	add := toolbar.Modal("exchange.add").Add(web.URLFor("ExchangeRateController.Add"))
	edit := toolbar.Modal("exchange.edit").Edit(web.URLFor("ExchangeRateController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	mc.AddButtons(edit,add)
	mc.TplName = "master/views/exchange_rate/index.html"
}
func (mc *ExchangeRateController) Add() {
	item := models.ExchangeRate{}
	mc.show(&item)
}
func (mc *ExchangeRateController) Edit() {
	id, _ := mc.GetInt64("id", -1)
	exchange_rateService := services.NewExchangeRateService()
	item := exchange_rateService.GetExchangeRateById(id)
	if item == nil {
		mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	mc.show(item)
}
func (mc *ExchangeRateController) show(item *models.ExchangeRate) {
	currencyService := services.NewCurrencyService()
	currency_list := currencyService.GetAll(*mc.GetQueryParams())
	mc.Data["currency_list"] = currency_list
	mc.Data["data"] = item
	save := toolbar.Ajax("exchange_rate.submit").Submit("exchange_rateForm", web.URLFor("ExchangeRateController.Update"))
	mc.AddButtons(save)
	mc.TplName = "master/views/exchange_rate/edit.html"
}
func (mc *ExchangeRateController) Search() {
	exchange_rateService := services.NewExchangeRateService()
	data := exchange_rateService.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}
func (mc *ExchangeRateController) Update() {
	exchange_rateForm := form.ExchangeRateForm{}
	if err := mc.ParseForm(&exchange_rateForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := exchange_rateForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
		//mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	t1 := exchangeRate.ExchangeRateUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
