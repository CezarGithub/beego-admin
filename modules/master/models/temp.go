package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/client/orm"
)

type Temp struct {
	Name string `orm:"column(name);unique;type(text)" description:"Legal name" json:"name"  i18n:"master.county.name" form:"name"`
	models.Base
}

// TableName
func (*Temp) TableName() string {
	return "temp"
}

//init model
func init() {
	database.RegisterModel("master", new(Temp))
}

func (*Temp) SearchField() []string {
	return nil
}
func (*Temp) WhereCondition() *orm.Condition {
	return nil
}
// TimeField
func (*Temp) TimeField() []string {
	return []string{}
}

func (t *Temp) Export() []byte {
	var items []*Temp
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Temp) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Temp
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		return err
	}
	tx.QueryTable(t.TableName()).Filter("id__gt",0).Delete()
	for _, item := range list {
		_, err := tx.Insert(item)
		if err != nil {
			return err
		}
	}
	return nil
}

