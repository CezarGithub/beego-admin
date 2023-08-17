package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Companies group struct
type Group struct {
	Name        string `orm:"column(name);unique;type(text)" description:"Legal name" json:"name" i18n:"master.company.groupname"`
	Status      int8   `orm:"column(status);size(1);default(1)" description:"Status 0:disabled 1:enabled" json:"status" i18n:"master.company.status"`
	Description string `orm:"column(description);type(text)" description:"Comments" json:"description" i18n:"master.company.description"`
	models.Base
}

// TableName
func (*Group) TableName() string {
	return "master_company_group"
}

// SearchField
func (*Group) SearchField() []string {
	return []string{"name"}
}

func (cs *Group) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Name != "" {
			cond = cond.And("name__icontains", cs.Name)
		}

	}
	return cond
}

// TimeField
func (*Group) TimeField() []string {
	return []string{}
}

// init model
func init() {
	database.RegisterModel("master", new(Group))
}

func (c *Group) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	err := validation.ValidateStruct(c, rules...)
	return err
}
func (t *Group) Export() []byte {
	var items []*Group
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Group) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Group
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
