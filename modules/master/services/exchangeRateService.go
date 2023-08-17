package services

import (
	"net/url"
	"quince/modules/admin/services"
	"quince/modules/master/models"
)

// ExchangeRateService struct
type exchange_rateService struct {
	services.BaseService
}

// NewExchangeRateService - instantiate de IModel filter
func NewExchangeRateService() exchange_rateService {
	var cs exchange_rateService
	c := models.ExchangeRate{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *exchange_rateService) GetAll(params url.Values) []*models.ExchangeRate {
	var list []*models.ExchangeRate
	o := cs.DataQuery().QueryTable(new(models.ExchangeRate)).RelatedSel()
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}

// Get ById
func (cs *exchange_rateService) GetExchangeRateById(id int64) *models.ExchangeRate {
	var item models.ExchangeRate
	err := cs.DataQuery().QueryTable(new(models.ExchangeRate)).Filter("id", id).RelatedSel().One(&item)
	if err == nil {
		return &item
	}

	return nil
}
