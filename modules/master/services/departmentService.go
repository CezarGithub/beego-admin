package services


import (
	"net/url"
	"quince/modules/admin/services"
	"quince/modules/master/models"
)

// DepartmentService struct
type departmentService struct {
	services.BaseService
}

// NewDepartmentService - instantiate de IModel filter
func NewDepartmentService() departmentService {
	var cs departmentService
	c := models.Department{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *departmentService) GetAll(params url.Values) []*models.Department {
	var list []*models.Department
	o :=cs.DataQuery().QueryTable(new(models.Department)).RelatedSel()
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}
// Get ById
func (cs *departmentService) GetDepartmentById(id int64) *models.Department {
	var item models.Department
	err := cs.DataQuery().QueryTable(new(models.Department)).Filter("id", id).RelatedSel().One(&item)
	if err == nil {
		return &item
	}

	return nil
}
                