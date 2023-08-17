package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// AdminRole struct
type AdminRole struct {
	Name        string `orm:"column(name);size(50)" description:"Name" i18n:"admin.user.name"`
	Description string `orm:"column(description);size(100)" description:"Description" i18n:"admin.user.description"`
	Url         string `orm:"column(url);size(1000)" description:"Menus URLs" i18n:"admin.user.url"`
	Route       string `orm:"column(route);size(1000)" description:"Routes URLs" i18n:"admin.user.menu"`
	Status      int8   `orm:"column(status);size(1);default(1)" description:"Enabled 0：NO 1：YES" i18n:"admin.user.status"`
	Base
}

// SearchField
func (*AdminRole) SearchField() []string {
	return []string{"name", "description"}
}

func (cs *AdminRole) WhereCondition() *orm.Condition {
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
func (*AdminRole) TimeField() []string {
	return []string{}
}

// TableName
func (*AdminRole) TableName() string {
	return "admin_role"
}

//Register the defined model in init
func init() {
	database.RegisterModel("admin", new(AdminRole))
}

func (c *AdminRole) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Description, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Status, validation.Required, validation.Min(0)))
	err := validation.ValidateStruct(c, rules...)
	return err
}
func (t *AdminRole) Export() []byte {
	var items []*AdminRole
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *AdminRole) Import(tx orm.TxOrmer, data []byte) error {
	var list []*AdminRole
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
