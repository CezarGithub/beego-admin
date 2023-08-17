package services

import (
	"net/url"
	"quince/modules/admin/services"
	"quince/modules/master/models"
	"quince/utils/page"
)

// CountyService struct
type countyService struct {
	services.BaseService
}

// NewCountyService - instantiate de IModel filter
func NewCountyService() countyService {
	var cs countyService
	c := models.County{}
	c.Country = &models.Country{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *countyService) GetAll(params url.Values) []*models.County {
	var list []*models.County
	//o := orm.NewOrm().QueryTable(new(models.County)).RelatedSel().OrderBy("name")
	o := cs.DataQuery().QueryTable(new(models.County)).RelatedSel().OrderBy("name")
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}
func (cs *countyService) GetPaginateData(listRows int, params url.Values) ([]*models.County, page.Pagination) {
	var list []*models.County
	//o := orm.NewOrm().QueryTable(new(models.County)).RelatedSel().OrderBy("name")
	o := cs.DataQuery().QueryTable(new(models.County)).RelatedSel().OrderBy("name")
	_, err := cs.PaginateAndScopeWhere(o, listRows, params).All(&list)
	if err != nil {
		return nil, cs.Pagination
	}
	return list, cs.Pagination
}

// GetCountyById
func (cs *countyService) GetCountyById(id int) *models.County {
	var item models.County
	err := cs.DataQuery().QueryTable(new(models.County)).RelatedSel().Filter("id", id).One(&item)
	if err == nil {
		return &item
	}
	return nil
}
