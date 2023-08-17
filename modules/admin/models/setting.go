package models

import (
	"encoding/json"
	"quince/initialize/database"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

//Content
type Content struct {
	Name    string
	Field   string
	Type    string
	Content string
	Option  string
	Form    string
}

// Setting struct
type Setting struct {
	SettingGroupId int64    `orm:"column(setting_group_id);size(10);default(0)" description:"Belonging to the group" json:"setting_group_id"`
	Name           string `orm:"column(name);size(50)" description:"name" json:"name" i18n:"admin.settinggroup.name"`
	Description    string `orm:"column(description);size(100)" description:"description" json:"description" i18n:"admin.settinggroup.description"`
	Code           string `orm:"column(code);size(50)" description:"code" json:"code" i18n:"admin.settinggroup.code"`
	Content        string `orm:"column(content);type(text)" description:"Set configuration and content" json:"content"`
	//Convert Content json string to structure
	ContentStrut []*Content `orm:"-"`
	SortNumber   int        `orm:"column(sort_number);size(10);default(1000)" description:"Sort" json:"sort_number" i18n:"admin.settinggroup.sort_number"`
	Base
}

// TableName
func (*Setting) TableName() string {
	return "admin_setting"
}

// SearchField
func (*Setting) SearchField() []string {
	return []string{"name"}
}

// NoDeletionId
// func (*Setting) NoDeletionId() []int {
// 	return []int{1, 2, 3, 4, 5}
// }

func (cs *Setting) WhereCondition() *orm.Condition {
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
func (*Setting) TimeField() []string {
	return []string{}
}

//init
func init() {

	database.RegisterModel("admin", new(Setting))
}

func (t *Setting) Export() []byte {
	var items []*Setting
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Setting) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Setting
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
