package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/client/orm"
)

// SettingGroup struct
type SettingGroup struct {
	Module         string `orm:"column(module);size(30)" description:"Function module" json:"module" i18n:"admin.settinggroup.module"`
	Name           string `orm:"column(name);size(50)" description:"name" json:"name" i18n:"admin.settinggroup.name"`
	Description    string `orm:"column(description);size(100)" description:"description" json:"description" i18n:"admin.settinggroup.description"`
	Code           string `orm:"column(code);size(50)" description:"code" json:"code" i18n:"admin.settinggroup.code"`
	SortNumber     int    `orm:"column(sort_number);size(10);default(1000)" description:"sort" json:"sort_number" i18n:"admin.settinggroup.sort_number"`
	AutoCreateMenu int    `orm:"column(auto_create_menu);size(1);default(0)" description:"Automatically generate menu" json:"auto_create_menu"  i18n:"admin.settinggroup.auto_create_menu"`
	AutoCreateFile int    `orm:"column(auto_create_file);size(1);default(0)" description:"Automatic configuration file generation" json:"auto_create_file" i18n:"admin.settinggroup.auto_create_file"`
	Icon           string `orm:"column(icon);size(30);default(fa-list)" description:"Icon" json:"icon" i18n:"admin.settinggroup.icon"`
	Base
}

// TableName
func (*SettingGroup) TableName() string {
	return "admin_setting_group"
}

// SearchField
func (*SettingGroup) SearchField() []string {
	return []string{"name", "description", "code"}
}

func (cs *SettingGroup) WhereCondition() *orm.Condition {
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
func (*SettingGroup) TimeField() []string {
	return []string{}
}

// init
func init() {
	database.RegisterModel("admin", new(SettingGroup))
}
func (c *SettingGroup) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Code, validation.Required, validation.Length(1, 25)))
	err := validation.ValidateStruct(c, rules...)
	return err
}
func (t *SettingGroup) Export() []byte {
	var items []*SettingGroup
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *SettingGroup) Import(tx orm.TxOrmer, data []byte) error {
	var list []*SettingGroup
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
