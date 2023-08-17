package services

import (
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/core/logs"

)

// AdminRbacService struct
type adminRbacService struct {
	BaseService
}

// AdminLogService - instantiate de IModel filter
func NewAdminRbacService() adminRbacService {
	var cs adminRbacService
	c := models.AdminRbac{}
	cs.IModel = &c
	return cs
}

// GetAdminMenuByUrl
func (cs *adminRbacService) GetAdminRbacByRole(role_id int) *models.AdminRbac {
	var adminRbac models.AdminRbac
	err := cs.DataQuery().QueryTable(new(models.AdminRoute)).RelatedSel().Filter("role_id", role_id).One(&adminRbac)
	if err == nil {
		return &adminRbac
	}
	return nil
}

// GetCount
func (cs *adminRbacService) GetCount() int {
	count, err := cs.DataQuery().QueryTable(new(models.AdminRbac)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetCount per module
func (cs *adminRbacService) GetCountRbac(role_id int) int {
	count, err := cs.DataQuery().QueryTable(new(models.AdminRbac)).Filter("role_id", role_id).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// AllRoute
func (cs *adminRbacService) AllRbac() []*models.AdminRbac {
	var adminRbac []*models.AdminRbac
	_, err := cs.DataQuery().QueryTable(new(models.AdminRbac)).RelatedSel().OrderBy("role_id").All(&adminRbac)
	if err == nil {
		return adminRbac
	}
	return nil
}

// IsExistUrl
func (cs *adminRbacService) IsExistUrl(role_id int64, route_id int64) *models.AdminRbac {
	var adminRbac models.AdminRbac
	err := cs.DataQuery().QueryTable(new(models.AdminRbac)).Filter("role_id", role_id).Filter("route_id", route_id).One(&adminRbac)
	if err != nil {
		logs.Error(err)
		return nil
	}
	return &adminRbac
}

// GetAdminRouteById
func (cs *adminRbacService) GetAdminRbacById(id int64) *models.AdminRbac {
	var adminRbac models.AdminRbac
	err := cs.DataQuery().QueryTable(new(models.AdminRbac)).RelatedSel().Filter("id", id).One(&adminRbac)
	if err == nil {
		return &adminRbac

	}
	return nil
}
