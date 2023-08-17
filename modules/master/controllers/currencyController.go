package controllers

import (
	"errors"
	"quince/internal/toolbar"
	"quince/modules/admin/controllers"
	"quince/modules/master/form"
	"quince/modules/master/models"
	"quince/modules/master/services"
	tx "quince/modules/master/transactions/currency"

	"github.com/beego/beego/v2/server/web"
)

// CurrencyController struct
type CurrencyController struct {
	controllers.BaseController
}

// NestPrepare - Init controller
func (mc *CurrencyController) NestPrepare() {}

func (mc *CurrencyController) Index() {
	currencyService := services.NewCurrencyService()
	data := currencyService.GetAll(*mc.GetQueryParams())
	mc.Data["data"] = data
	add := toolbar.Modal("currency.add").Add(web.URLFor("CurrencyController.Add"))
	del := toolbar.Ajax("currency.del").DeleteSelectedRows(web.URLFor("CurrencyController.Delete"))
	edit := toolbar.Modal("currency.edit").Edit(web.URLFor("CurrencyController.Edit"))
	edit.AddQueryParams("id", "{{.Item.Id}}")
	mc.AddButtons(add,del,edit)
	mc.TplName = "master/views/currency/index.html"
}
func (mc *CurrencyController) Add() {
	item := models.Currency{}
	mc.show(&item)
}
func (mc *CurrencyController) Edit() {
	id, _ := mc.GetInt64("id", -1)
	currencyService := services.NewCurrencyService()
	item := currencyService.GetCurrencyById(id)
	if item == nil {
		mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	mc.show(item)
}
func (mc *CurrencyController) show(item *models.Currency) {
	mc.Data["data"] = item

	save := toolbar.Ajax("currency.submit").Submit("currencyForm", web.URLFor("CurrencyController.Update"))
	mc.AddButtons(save)
	mc.TplName = "master/views/currency/edit.html"
}
func (mc *CurrencyController) Search() {
	currencyService := services.NewCurrencyService()
	data := currencyService.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}

// Del
func (uc *CurrencyController) Delete() {
	idArr := uc.GetSelectedIDs()
	t1 := tx.CurrencyDelete(idArr)
	uc.Transaction.Add(t1)
	err := uc.Transaction.Execute()

	if err == nil {
		uc.ResponseSuccess(uc.Ctx)
	} else {
		uc.ResponseErrorWithMessage(err, uc.Ctx)
	}
}
func (mc *CurrencyController) Update() {
	currencyForm := form.CurrencyForm{}
	if err := mc.ParseForm(&currencyForm); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := currencyForm.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
	t1 := tx.CurrencyUpdate(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
