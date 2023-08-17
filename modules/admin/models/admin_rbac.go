package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Admin route based acces control struct
type AdminRbac struct {
	Role   *AdminRole  `orm:"rel(fk)"`
	Route  *AdminRoute `orm:"rel(fk)"`
	Status int8        `orm:"column(status);size(1);default(1)" description:"Enabled 0：NO 1：YES" i18n:"admin.user.status"`
	Base
}

// TableName
func (*AdminRbac) TableName() string {
	return "admin_rbac"
}

// init
func init() {
	database.RegisterModel("admin", new(AdminRbac))
}

// GetLogMethod
func (*AdminRbac) GetLogMethod() []string {
	return []string{"OFF", "GET", "POST", "PUT", "DELETE"}
}

// SearchField
func (*AdminRbac) SearchField() []string {
	return []string{}
}

func (cs *AdminRbac) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
	}
	return cond
}

// TimeField
func (*AdminRbac) TimeField() []string {
	return []string{}
}

func (c *AdminRbac) Validate() error {
	rules := []*validation.FieldRules{}
	err := validation.ValidateStruct(c, rules...)
	return err
}

// 
func (t *AdminRbac) Export() []byte {
	var items []*AdminRbac
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *AdminRbac) Import(tx orm.TxOrmer, data []byte) error {
	var list []*AdminRbac
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

