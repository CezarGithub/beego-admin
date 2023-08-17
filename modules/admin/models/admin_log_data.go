package models

import (
	"encoding/json"
	"quince/initialize/database"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// AdminLogData struct
type AdminLogData struct {
	AdminLogId int    `orm:"column(admin_log_id);size(11)" description:"Log ID"   json:"logID"  i18n:"admin.log.row_id"`
	Data       string `orm:"column(data);type(text)" description:"Log content"   json:"data"  i18n:"admin.log.data"`
	Base
}

// TableName
func (*AdminLogData) TableName() string {
	return "admin_log_data"
}

// SearchField
func (*AdminLogData) SearchField() []string {
	return []string{}
}

func (cs *AdminLogData) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Data != "" {
			cond = cond.And("data__icontains", cs.Data)
		}
	}
	return cond
}

// TimeField
func (*AdminLogData) TimeField() []string {
	return []string{}
}

// init
func init() {
	database.RegisterModel("admin", new(AdminLogData))
}
func (t *AdminLogData) Export() []byte {
	var items []*AdminLogData
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *AdminLogData) Import(tx orm.TxOrmer, data []byte) error {
	var list []*AdminLogData
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
