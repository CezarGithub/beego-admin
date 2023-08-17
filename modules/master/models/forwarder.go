package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Forwarder struct
type Forwarder struct {
	Name string `orm:"column(name);unique;type(text)" description:"Legal name" json:"name"`
	Code string `orm:"column(code);type(text)" description:"Forwarder code" json:"code"`
	models.Base
}

// TableName
func (*Forwarder) TableName() string {
	return "master_forwarder"
}

// SearchField
func (*Forwarder) SearchField() []string {
	return []string{"name", "code"}
}


// WhereField
func (*Forwarder) WhereField() []string {
	return []string{}
}

// TimeField
func (*Forwarder) TimeField() []string {
	return []string{}
}
func (t *Forwarder) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if t != nil {
		if t.Id != 0 {
			cond = cond.And("id", t.Id)
		}
		if t.Name != "" {
			cond = cond.And("name__icontains", t.Name)
		}
	}
	return cond
}
//init model
func init() {
	database.RegisterModel("master", new(Forwarder))
}

func (t *Forwarder) Export() []byte {
	var items []*Forwarder
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Forwarder) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Forwarder
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

