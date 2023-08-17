package models

import (
	"encoding/json"
	"quince/initialize/database"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// AdminLog struct
type AdminLog struct {
	AdminUserId int64    `orm:"column(admin_user_id);size(10);default(0)" description:"User id"`
	Name        string `orm:"column(name);size(30)" description:"name" i18n:"admin.user.name"`
	Url         string `orm:"column(url);size(100)" description:"url" i18n:"app.url"`
	LogMethod   string `orm:"column(log_method);size(8);default(OFF)" description:"Logging method"  i18n:"app.logging_method"`
	LogIp       string `orm:"column(log_ip);size(20)" description:"Log IP"    i18n:"app.ip"`
	Base
}

// SearchField
func (*AdminLog) SearchField() []string {
	return []string{"name", "url", "log_ip"}
}

// TimeField
func (*AdminLog) TimeField() []string {
	return []string{"create_time"}
}

func (cs *AdminLog) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Name != "" {
			cond = cond.And("name__icontains", cs.Name)
		}
		if cs.Url != "" {
			cond = cond.And("url__icontains", cs.Url)
		}
		if cs.LogIp != "" {
			cond = cond.And("log_ip__icontains", cs.LogIp)
		}
	}
	return cond
}

// TableName
func (*AdminLog) TableName() string {
	return "admin_log"
}

// init
func init() {
	database.RegisterModel("admin", new(AdminLog))
}

func (t *AdminLog) Export() []byte {
	var items []*AdminLog
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *AdminLog) Import(tx orm.TxOrmer, data []byte) error {
	var list []*AdminLog
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

