package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Site struct
type Site struct {
	Company   *Company `orm:"rel(fk)"`
	Name      string   `orm:"column(name);unique;type(text)" description:"Site description name" json:"name"`
	URL       string   `orm:"column(url);type(text)" description:"Site address" json:"url"`
	IsDefault int8     `orm:"column(is_default);size(1)" description:"Enabled 0:yes 1:no" json:"is_default"`
	models.Base
}

// TableName
func (*Site) TableName() string {
	return "master_site"
}

// SearchField
func (*Site) SearchField() []string {
	return []string{"company_id", "url"}
}


// WhereField
func (*Site) WhereField() []string {
	return []string{}
}

// TimeField
func (*Site) TimeField() []string {
	return []string{}
}
func (t *Site) WhereCondition() *orm.Condition {
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
	database.RegisterModel("master", new(Site))
}

func (t *Site) Export() []byte {
	var items []*Site
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Site) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Site
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
