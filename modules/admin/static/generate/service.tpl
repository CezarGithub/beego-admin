package services


import (
	"net/url"
	"quince/modules/admin/services"
	"quince/modules/{{.Model.Module}}/models"
)

// {{.Model.Name}}Service struct
type {{.Model.Name_lcase}}Service struct {
	services.BaseService
}

// New{{.Model.Name}}Service - instantiate de IModel filter
func New{{.Model.Name}}Service() {{.Model.Name_lcase}}Service {
	var cs {{.Model.Name_lcase}}Service
	c := models.{{.Model.Name}}{}
	cs.IModel = &c
	return cs
}

// GetAll
func (cs *{{.Model.Name_lcase}}Service) GetAll(params url.Values) []*models.{{.Model.Name}} {
	var list []*models.{{.Model.Name}}
	o := cs.DataQuery().QueryTable(new(models.{{.Model.Name}})).RelatedSel()
	_, err := cs.GetAllAndScopeWhere(o, params).All(&list)
	if err != nil {
		return nil
	}
	return list
}
// Get ById
func (cs *{{.Model.Name_lcase}}Service) Get{{.Model.Name}}ById(id int64) *models.{{.Model.Name}} {
	var item models.{{.Model.Name}}
	err := cs.DataQuery().QueryTable(new(models.{{.Model.Name}})).Filter("id", id).RelatedSel().One(&item)
	if err == nil {
		return &item
	}

	return nil
}