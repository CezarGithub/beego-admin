package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// AdminMenu struct
type AdminMenu struct {
	ParentId     int64  `orm:"column(parent_id);size(11);default(0)" description:"Parent menu" json:"parent_id" i18n:"admin.menu.parent"`
	Name         string `orm:"column(name);size(30)" description:"name" json:"name"  i18n:"app.name"`
	Module       string `orm:"column(module);size(30)" description:"module"  json:"module"  i18n:"app.module"`
	Url          string `orm:"column(url);size(100);index" description:"url"  json:"url"   i18n:"app.url"`
	Icon         string `orm:"column(icon);size(30);default(fa-list)" description:"icon"  json:"icon"   i18n:"app.icon"`
	IsShow       int8   `orm:"column(is_show);size(1);default(1)" description:"grade"  json:"is_show"  i18n:"admin.menu.is_show"`
	IsShowMobile int8   `orm:"column(is_show_mobile);size(1);default(1)" description:"grade"  json:"is_show_mobile"  i18n:"admin.menu.is_show_mobile"`
	SortId       int    `orm:"column(sort_id);size(10);default(1000)" description:"Sort"  json:"sort_id"  i18n:"app.sort"`
	LogMethod    string `orm:"column(log_method);size(8);default(OFF)" description:"Logging method"  json:"log_method"  i18n:"app.logging_method"`
	Status       int8   `orm:"column(status);size(1);default(1)" description:"Enabled 0：NO 1：YES"  json:"status"  i18n:"admin.user.status"`
	Base
}

// TableName
func (*AdminMenu) TableName() string {
	return "admin_menu"
}

// init
func init() {
	database.RegisterModel("admin", new(AdminMenu))
}

// GetLogMethod
func (*AdminMenu) GetLogMethod() []string {
	return []string{"OFF", "GET", "POST", "PUT", "DELETE"}
}

// SearchField
func (*AdminMenu) SearchField() []string {
	return []string{"name", "module"}
}

func (cs *AdminMenu) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Name != "" {
			cond = cond.And("name__icontains", cs.Name)
		}
		if cs.Module != "" {
			cond = cond.And("module__icontains", cs.Module)
		}

	}
	return cond
}

// TimeField
func (*AdminMenu) TimeField() []string {
	return []string{}
}

func (c *AdminMenu) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Module, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Url, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Icon, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.LogMethod, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.ParentId, validation.Required, validation.Min(1)))
	//rules = append(rules, validation.Field(&c.SortId, validation.Required))
	err := validation.ValidateStruct(c, rules...)
	return err
}

// NoDeletionId
//
//	func (*AdminMenu) NoDeletionId() []int {
//		return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
//	}
func (t *AdminMenu) Export() []byte {
	var items []*AdminMenu
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *AdminMenu) Import(tx orm.TxOrmer, data []byte) error {
	var list []*AdminMenu
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
