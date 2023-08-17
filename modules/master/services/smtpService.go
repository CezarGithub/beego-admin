package services

import (
	"net/url"
	"quince/global"
	"quince/modules/admin/services"
	"quince/modules/master/models"
	"quince/utils/encrypter"
	"quince/utils/page"
)

// SmtpService struct
type smtpService struct {
	services.BaseService
}

// NewCountyService - instantiate de IModel filter
func NewSMTPService() smtpService {
	var cs smtpService
	c := models.SMTP{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *smtpService) GetAll(params url.Values) []*models.SMTP {
	var list []*models.SMTP
	o := cs.DataQuery().QueryTable(new(models.SMTP)).RelatedSel().OrderBy("server")
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}
func (cs *smtpService) GetPaginateData(listRows int, params url.Values) ([]*models.SMTP, page.Pagination) {
	var list []*models.SMTP
	o := cs.DataQuery().QueryTable(new(models.SMTP)).RelatedSel().OrderBy("server")
	_, err := cs.PaginateAndScopeWhere(o, listRows, params).All(&list)
	if err != nil {
		return nil, cs.Pagination
	}
	return list, cs.Pagination
}
func (cs *smtpService) GetSMTPById(id int64) *models.SMTP {
	var item models.SMTP
	err := cs.DataQuery().QueryTable(new(models.SMTP)).RelatedSel().Filter("id", id).One(&item)
	if err == nil {
		cryptData := encrypter.Decrypt(item.Password, []byte(global.BA_CONFIG.Other.LogAesKey))
		item.Password = cryptData
		return &item
	}
	return nil
}
func (cs *smtpService) GetSMTPByCompanyId(companyId int64) *models.SMTP {
	var item models.SMTP
	err := cs.DataQuery().QueryTable(new(models.SMTP)).RelatedSel().Filter("company_id", companyId).One(&item)
	if err == nil {
		return &item
	}
	return nil
}
