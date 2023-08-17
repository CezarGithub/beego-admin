package services

import (
	"net/url"
	"quince/modules/admin/services"
	"quince/modules/master/models"
)

// CurrencyService struct
type currencyService struct {
	services.BaseService
}

// NewCurrencyService - instantiate de IModel filter
func NewCurrencyService() currencyService {
	var cs currencyService
	c := models.Currency{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *currencyService) GetAll(params url.Values) []*models.Currency {
	var list []*models.Currency
	o := cs.DataQuery().QueryTable(new(models.Currency)).RelatedSel()
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}

// Get ById
func (cs *currencyService) GetCurrencyById(id int64) *models.Currency {
	var item models.Currency
	err := cs.DataQuery().QueryTable(new(models.Currency)).Filter("id", id).RelatedSel().One(&item)
	if err == nil {
		return &item
	}

	return nil
}

// Check if we have records in ExchangeRate tables .If Yes ,change of BaseCurrency is not possible.
func (cs *currencyService) ChangeBaseCurrency(id int64) int64 {
	if id > 0 {//is not insert
		c := cs.GetCurrencyById(id)
		if c.BaseCurrency == 1 { //is the base currency -skip check
			return 0
		}
	}
	n, _ := cs.DataQuery().QueryTable(new(models.ExchangeRate)).Count()
	return n
}
