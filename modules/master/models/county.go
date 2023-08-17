package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/modules/admin/models"

	//"quince/internal/validation/is"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// County struct
type County struct {
	Name    string   `orm:"column(name);unique;type(text)" description:"Legal name" json:"name"  i18n:"master.county.name"`
	Code    string   `orm:"column(code);type(text)" description:"County code" json:"code" i18n:"master.county.code" `
	Country *Country `orm:"rel(fk)"`
	models.Base
}

// TableName
func (*County) TableName() string {
	return "master_county"
}

// SearchField
func (*County) SearchField() []string {
	return []string{"id", "name", "code"}
}

func (cs *County) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Name != "" {
			cond = cond.And("name__icontains", cs.Name)
		}
		if cs.Country != nil {
			if cs.Country.Id != 0 {
				cond = cond.And("country_id", cs.Country.Id)
			}
		}
	}
	return cond
}

// TimeField
func (*County) TimeField() []string {
	return []string{}
}

func (c *County) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Code, validation.Required, validation.Length(1, 25)))
	rules = append(rules, validation.Field(&c.Country))
	err := validation.ValidateStruct(c, rules...)
	return err
}

//init model
func init() {
	database.RegisterModel("master", new(County))

}

func (t *County) Export() []byte {
	var items []*County
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *County) Import(tx orm.TxOrmer, data []byte) error {
	var list []*County
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

