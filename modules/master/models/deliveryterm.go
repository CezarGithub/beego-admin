package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// DeliveryTerm struct
type DeliveryTerm struct {
	Name string `orm:"column(name);unique;type(text)" description:"Legal name" json:"name"`
	Code string `orm:"column(code);type(text)" description:"UnitMeasure code" json:"code"`
	models.Base
}

// TableName
func (*DeliveryTerm) TableName() string {
	return "master_delivery_term"
}

// SearchField
func (*DeliveryTerm) SearchField() []string {
	return []string{"name", "code"}
}

// NoDeletionId
func (*DeliveryTerm) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*DeliveryTerm) WhereField() []string {
	return []string{}
}

// TimeField
func (*DeliveryTerm) TimeField() []string {
	return []string{}
}
func (t *DeliveryTerm) WhereCondition() *orm.Condition {
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
	database.RegisterModel("master", new(DeliveryTerm))
}

func (t *DeliveryTerm) Export() []byte {
	var items []*DeliveryTerm
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *DeliveryTerm) Import(tx orm.TxOrmer, data []byte) error {
	var list []*DeliveryTerm
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

