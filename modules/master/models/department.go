package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Department struct
type Department struct {
	Company *Company `orm:"rel(fk)"`
	Name    string   `orm:"column(name);unique;type(text)" description:"Legal name" json:"name"`
	Code    string   `orm:"column(code);type(text)" description:"Department code" json:"code"`
	models.Base
}

func NewDepartment() Department {
	company := new(Company)
	department := Department{Base: models.Base{Id: 0}, Company: company}
	return department
}

// TableName
func (*Department) TableName() string {
	return "master_department"
}

// SearchField
func (*Department) SearchField() []string {
	return []string{"name"}
}

// TimeField
func (*Department) TimeField() []string {
	return []string{}
}
func (cs *Department) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Name != "" {
			cond = cond.And("name__icontains", cs.Name)
		}
		if cs.Code != "" {
			cond = cond.And("code__icontains", cs.Code)
		}
	}
	return cond
}
func (c *Department) Validate() error {
	rules := []*validation.FieldRules{}
	//rules = append(rules, validation.Field(&c.Id, validation.Required, validation.Min(-1)))
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Code, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Company))
	err := validation.ValidateStruct(c, rules...)
	return err
}

// init model
func init() {
	database.RegisterModel("master", new(Department))
}

func (t *Department) Export() []byte {
	var items []*Department
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Department) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Department
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
