package services

import (
	"net/url"
	"quince/modules/admin/services"
	"quince/modules/master/models"
	"quince/utils/page"
)

// GroupService struct
type groupService struct {
	services.BaseService
}

// NewCountyService - instantiate de IModel filter
func NewGroupService() groupService {
	var cs groupService
	c := models.Group{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *groupService) GetAll(params url.Values) []*models.Group {
	var list []*models.Group
	o := cs.DataQuery().QueryTable(new(models.Group)).OrderBy("name")
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}
func (cs *groupService) GetPaginateData(listRows int, params url.Values) ([]*models.Group, page.Pagination) {
	var list []*models.Group
	o := cs.DataQuery().QueryTable(new(models.Group)).OrderBy("name")
	_, err := cs.PaginateAndScopeWhere(o, listRows, params).All(&list)
	if err != nil {
		return nil, cs.Pagination
	}
	return list, cs.Pagination
}

// GetCountyById
func (cs *groupService) GetGroupById(id int) *models.Group {
	var item models.Group
	err := cs.DataQuery().QueryTable(new(models.Group)).Filter("id", id).One(&item)
	if err == nil {
		return &item
	}
	return nil
}
