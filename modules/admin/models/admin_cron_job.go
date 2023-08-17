package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// AdminCronJob struct
type AdminCronJob struct {
	Name        string `orm:"column(name);unique;size(30)" description:"name" json:"name"  i18n:"app.name"`
	Description string `orm:"column(description);unique;size(230)" description:"name" json:"description"  i18n:"app.description"`
	Module      string `orm:"column(module);size(30)" description:"module" json:"module"  i18n:"app.module"`
	Status      int8   `orm:"column(status);size(1);default(1)" description:"Enabled 0：NO 1：YES" json:"status" i18n:"admin.user.status"`
	Base
	Intervals []*AdminCronJobInterval `orm:"reverse(many)"`
}

// SearchField
func (*AdminCronJob) SearchField() []string {
	return []string{"name", "module", "description"}
}

// TimeField
func (*AdminCronJob) TimeField() []string {
	return []string{"create_time"}
}

func (cs *AdminCronJob) WhereCondition() *orm.Condition {
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

// TableName
func (*AdminCronJob) TableName() string {
	return "admin_cron_job"
}
func (c *AdminCronJob) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 50)))
	rules = append(rules, validation.Field(&c.Module, validation.Required, validation.Length(1, 25)))
	rules = append(rules, validation.Field(&c.Description, validation.Length(1, 255)))
	err := validation.ValidateStruct(c, rules...)
	return err
}

// init
func init() {
	database.RegisterModel("admin", new(AdminCronJob))
}

func (t *AdminCronJob) Export() []byte {
	var items []*AdminCronJob
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *AdminCronJob) Import(tx orm.TxOrmer, data []byte) error {
	var list []*AdminCronJob
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
