package services

import (
	"net/url"
	"quince/modules/admin/models"
	"quince/utils/page"
)

// AdminRoleService struct
type adminRoleService struct {
	BaseService
}

// NewAdminUserService - instantiate de IModel filter
func NewAdminRoleService() adminRoleService {
	var cs adminRoleService
	c := models.AdminRole{}
	cs.IModel = &c
	return cs
}

// GetCount
func (cs *adminRoleService) GetCount() int {
	count, err := cs.DataQuery().QueryTable(new(models.AdminRole)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetAllData
func (cs *adminRoleService) GetAllData() []*models.AdminRole {
	var adminRoles []*models.AdminRole
	cs.DataQuery().QueryTable(new(models.AdminRole)).All(&adminRoles)
	return adminRoles
}

// GetPaginateData adminrole
func (ars *adminRoleService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminRole, page.Pagination) {

	ars.SearchField = append(ars.SearchField, new(models.AdminRole).SearchField()...)

	var adminRole []*models.AdminRole
	o := ars.DataQuery().QueryTable(new(models.AdminRole))
	_, err := ars.PaginateAndScopeWhere(o, listRows, params).All(&adminRole)
	if err != nil {
		return nil, ars.Pagination
	}
	return adminRole, ars.Pagination
}

// IsExistName
func (cs *adminRoleService) IsExistName(name string, id int64) bool {
	if id == 0 {
		return cs.DataQuery().QueryTable(new(models.AdminRole)).Filter("name", name).Exist()
	}
	return cs.DataQuery().QueryTable(new(models.AdminRole)).Filter("name", name).Exclude("id", id).Exist()
}

// GetAdminRoleById
func (cs *adminRoleService) GetAdminRoleById(id int64) *models.AdminRole {
	var adminRole models.AdminRole
	err := cs.DataQuery().QueryTable(new(models.AdminRole)).Filter("id", id).One(&adminRole)
	if err == nil {
		return &adminRole
	}
	return nil
}
