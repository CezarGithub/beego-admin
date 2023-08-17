package services

import (
	"net/url"
	"quince/modules/admin/services"
	"quince/modules/master/models"
	"quince/utils/page"
)

// CountryService struct
type countryService struct {
	services.BaseService
}

// NewCountryService - instantiate de IModel filter
func NewCountryService() countryService {
	var cs countryService
	c := models.Country{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *countryService) GetAll(params url.Values) []*models.Country {
	cs.IModel = new(models.Country)
	var list []*models.Country
	o := cs.DataQuery().QueryTable(new(models.Country)).OrderBy("name")
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}
func (cs *countryService) GetPaginateData(listRows int, params url.Values) ([]*models.Country, page.Pagination) {
	cs.IModel = new(models.Country)
	var list []*models.Country

	o := cs.DataQuery().QueryTable(new(models.Country)).OrderBy("name")
	_, err := cs.PaginateAndScopeWhere(o, listRows, params).All(&list)
	if err != nil {
		return nil, cs.Pagination
	}
	return list, cs.Pagination
}

// GetCountryById
func (cs *countryService) GetCountryById(id int) *models.Country {
	var item models.Country
	err := cs.DataQuery().QueryTable(new(models.Country)).Filter("id", id).One(&item)
	if err == nil {
		return &item
	}
	return nil
}
