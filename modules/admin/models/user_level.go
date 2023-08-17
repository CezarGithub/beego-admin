package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// UserLevel struct
type UserLevel struct {
	Name        string `orm:"column(name);size(20)" description:"Name" json:"name"  i18n:"admin.user.username"`
	Description string `orm:"column(description);size(50)" description:"Description" json:"description" i18n:"admin.user.description"`
	Img         string `orm:"column(img);size(255)" description:"Picture" json:"img" i18n:"admin.user.img"`
	Status      int8   `orm:"column(status);size(1);default(1)" description:"Enabled 0：NO 1：YES" json:"status"  i18n:"admin.user.status"`
	Base
}

// TableName
func (*UserLevel) TableName() string {
	return "admin_user_level"
}

// SearchField
func (*UserLevel) SearchField() []string {
	return []string{"name", "description"}
}

func (cs *UserLevel) WhereCondition() *orm.Condition {
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
func (*UserLevel) TimeField() []string {
	return []string{}
}

// init
func init() {
	database.RegisterModel("admin", new(UserLevel))
}

func (c *UserLevel) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Description, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Status, validation.Required, validation.Min(0)))
	err := validation.ValidateStruct(c, rules...)
	return err
}
func (t *UserLevel) Export() []byte {
	var items []*UserLevel
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *UserLevel) Import(tx orm.TxOrmer, data []byte) error {
	var list []*UserLevel
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		return err
	}
	admin, _ := web.AppConfig.String("adminname")
	tx.QueryTable(t.TableName()).Exclude("name", admin).Delete() //do not delete super_admin
	for _, item := range list {
		if item.Name != admin {
			_, err := tx.Insert(item)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
