package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/initialize/scheduler"
	"quince/internal/validation"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// AdminCronJob struct
type AdminCronJobInterval struct {
	AdminCronJob *AdminCronJob `orm:"rel(fk)"`
	scheduler.Interval
	Status      int8   `orm:"column(status);size(1);default(1)" description:"Enabled 0：NO 1：YES" json:"status" i18n:"admin.user.status"`
	RunOnce     int8   `orm:"column(run_once);size(1);default(1)" description:"Enabled 0：NO 1：YES" json:"runOnce" i18n:"admin.user.runOnce"`
	Description string `orm:"column(description);unique;size(230)" description:"name" json:"description"  i18n:"app.description"`
	Base
}

func NewAdminCronJobInterval() AdminCronJobInterval {
	admin_cron_job := new(AdminCronJob)
	interval := scheduler.Interval{}
	admin_cron_job_interval := AdminCronJobInterval{Base: Base{Id: 0}, AdminCronJob: admin_cron_job, Interval: interval}
	return admin_cron_job_interval
}

// SearchField
func (*AdminCronJobInterval) SearchField() []string {
	return []string{}
}

// TimeField
func (*AdminCronJobInterval) TimeField() []string {
	return []string{"create_time"}
}

func (cs *AdminCronJobInterval) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
	}
	return cond
}

// TableName
func (*AdminCronJobInterval) TableName() string {
	return "admin_cron_job_interval"
}
func (c *AdminCronJobInterval) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.AdminCronJob))
	err := validation.ValidateStruct(c, rules...)
	return err
}

// init
func init() {
	database.RegisterModel("admin", new(AdminCronJobInterval))
}

func (t *AdminCronJobInterval) Export() []byte {
	var items []*AdminCronJobInterval
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *AdminCronJobInterval) Import(tx orm.TxOrmer, data []byte) error {
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
