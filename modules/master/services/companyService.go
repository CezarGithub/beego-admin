package services

import (
	"net/url"
	"quince/modules/admin/services"
	"quince/modules/master/models"
)

// CompanyService struct
type companyService struct {
	services.BaseService
}

// NewCountryService - instantiate de IModel filter
func NewCompanyService() companyService {
	var cs companyService
	c := models.Company{}
	cs.IModel = &c
	return cs
}

// GetAllCompanies
func (cs *companyService) GetAll(params url.Values) []*models.Company {
	var list []*models.Company
	o := cs.DataQuery().QueryTable(new(models.Company)).RelatedSel().OrderBy("name")
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}

// GetCompanyById
func (cs *companyService) GetCompanyById(id int64) *models.Company {
	var item models.Company
	err := cs.DataQuery().QueryTable(new(models.Company)).Filter("id", id).RelatedSel().One(&item)
	if err == nil {
		return &item
	}

	return nil
}
