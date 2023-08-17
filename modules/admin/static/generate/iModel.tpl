package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/internal/validation/is"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// {{.Model.Name}} struct
type {{.Model.Name}} struct {
	models.Base
}

// TableName
func (*{{.Model.Name}}) TableName() string {
	return "{{.Model.Module}}_{{.Model.Name_lcase}}"
}

// SearchField
func (*{{.Model.Name}}) SearchField() []string {
	return []string{}
}

func (cs *{{.Model.Name}}) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
	}
	return cond
}

// TimeField
func (*{{.Model.Name}}) TimeField() []string {
	return []string{}
}

// init model
func init() {
	database.RegisterModel("{{.Model.Module}}", new({{.Model.Name}}))
}
func (c *{{.Model.Name}}) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Id, validation.Required, validation.Min(0)))
	err := validation.ValidateStruct(c, rules...)
	return err
}

func (t *{{.Model.Name}}) Export() []byte {
	var items []*{{.Model.Name}}
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *{{.Model.Name}}) Import(tx orm.TxOrmer, data []byte) error {
	var list []*{{.Model.Name}}
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		return err
	}
	tx.QueryTable(t.TableName()).Filter("id__gt", 0).Delete()
	for _, item := range list {
		_, err := tx.Insert(item)
		if err != nil {
			return err
		}
	}
	return nil
}
