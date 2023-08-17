package services

import (
	"net/url"
	"quince/modules/admin/models"
	"quince/utils/page"
)

// AdminCronJobService struct
type adminCronJobService struct {
	BaseService
}

// AdminLogService - instantiate de IModel filter
func NewAdminCronJobService() adminCronJobService {
	var cs adminCronJobService
	c := models.AdminCronJob{}
	cs.IModel = &c
	return cs
}

// GetPaginateData adminrole
func (ars *adminCronJobService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminCronJob, page.Pagination) {

	ars.SearchField = append(ars.SearchField, new(models.AdminRoute).SearchField()...)

	var adminCrons []*models.AdminCronJob
	o := ars.DataQuery().QueryTable(new(models.AdminCronJob)).RelatedSel()
	_, err := ars.PaginateAndScopeWhere(o, listRows, params).All(&adminCrons)
	if err != nil {
		return nil, ars.Pagination
	}
	return adminCrons, ars.Pagination
}

// IsExistName
func (ars *adminCronJobService) IsExistName(name string, id int64) bool {
	if id == 0 {
		return ars.DataQuery().QueryTable(new(models.AdminCronJob)).Filter("name", name).Exist()
	}
	return ars.DataQuery().QueryTable(new(models.AdminCronJob)).Filter("name", name).Exclude("id", id).Exist()
}
func (cs *adminCronJobService) GetAdminCronJobById(id int) *models.AdminCronJob {
	var item models.AdminCronJob
	//var itl *models.AdminCronJobInterval
	err := cs.DataQuery().QueryTable(new(models.AdminCronJob)).Filter("id", id).One(&item)
	_, _ = cs.DataQuery().LoadRelated(&item, "Intervals")
	if err == nil {
		return &item
	}
	return nil
}
