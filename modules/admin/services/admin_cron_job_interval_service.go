package services

import (
	"net/url"
	"quince/modules/admin/models"
)

// AdminCronJobIntervalService struct
type admin_cron_job_intervalService struct {
	BaseService
}

// NewAdminCronJobIntervalService - instantiate de IModel filter
func NewAdminCronJobIntervalService() admin_cron_job_intervalService {
	var cs admin_cron_job_intervalService
	c := models.AdminCronJobInterval{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *admin_cron_job_intervalService) GetAll(params url.Values) []*models.AdminCronJobInterval {
	var list []*models.AdminCronJobInterval
	o := cs.DataQuery().QueryTable(new(models.AdminCronJobInterval)).RelatedSel()
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}

// Get ById
func (cs *admin_cron_job_intervalService) GetAdminCronJobIntervalById(id int64) *models.AdminCronJobInterval {
	var item models.AdminCronJobInterval
	err := cs.DataQuery().QueryTable(new(models.AdminCronJobInterval)).Filter("id", id).RelatedSel().One(&item)
	if err == nil {
		return &item
	}

	return nil
}
