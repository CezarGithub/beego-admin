package models

import (
	"encoding/json"
	"quince/initialize/database"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// AdminLogTx struct
// WARNING : Raw sql insert in /internal/action.go:110
type AdminLogTx struct {
	Table      string `orm:"column(table_name);type(text)" description:"Table name"  json:"table"  i18n:"admin.log.table"`
	TableRowId int    `orm:"column(table_row_id);size(11)" description:"Row ID" json:"table_row_id"  i18n:"admin.log.row_id"`
	Data       string `orm:"column(data);type(text)" description:"Log content"  json:"data"  i18n:"admin.log.data"`
	Title       string `orm:"column(title);type(text)" description:"Log content"  json:"title"  i18n:"admin.log.data"`
	Base
}

// TableName
func (*AdminLogTx) TableName() string {
	return "admin_log_tx"
}

// SearchField
func (*AdminLogTx) SearchField() []string {
	return []string{}
}

func (cs *AdminLogTx) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Data != "" {
			cond = cond.And("data__icontains", cs.Data)
		}
		if cs.Table != "" {
			cond = cond.And("data__icontains", cs.Table)
		}
	}
	return cond
}

// TimeField
func (*AdminLogTx) TimeField() []string {
	return []string{}
}

// init
func init() {
	database.RegisterModel("admin", new(AdminLogTx))
}
func (t *AdminLogTx) Export() []byte {
	var items []*AdminLogTx
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *AdminLogTx) Import(tx orm.TxOrmer, data []byte) error {
	var list []*AdminLogTx
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
