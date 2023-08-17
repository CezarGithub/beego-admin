package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Currency struct
type Currency struct {
	Name         string `orm:"column(name);unique;type(text)" description:"Currency name" json:"name" i18n:"master.currency.name"`
	CurrencyCode string `orm:"column(currency_code);type(text)" description:"Currency code" json:"currency_code" i18n:"master.currency.code"`
	BaseCurrency int8   `orm:"column(status);size(1);default(1)" description:"Status 0:disabled 1:enabled" json:"status" i18n:"master.currency.base"`
	models.Base
}

// TableName
func (*Currency) TableName() string {
	return "master_currency"
}

// SearchField
func (*Currency) SearchField() []string {
	return []string{"name", "currency_code"}
}

// WhereField
func (*Currency) WhereField() []string {
	return []string{}
}

// TimeField
func (*Currency) TimeField() []string {
	return []string{}
}
func (t *Currency) WhereCondition() *orm.Condition {
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
func (c *Currency) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.CurrencyCode, validation.Required, validation.Length(1, 255)))
	err := validation.ValidateStruct(c, rules...)
	return err
}

// init model
func init() {
	database.RegisterModel("master", new(Currency))
}

func (t *Currency) Export() []byte {
	var items []*Currency
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Currency) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Currency
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
