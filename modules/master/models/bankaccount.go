package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// BankAccount struct
type BankAccount struct {
	Company      *Company `orm:"rel(fk)"`
	Name         string   `orm:"column(name);unique;type(text)" description:"Legal name" json:"name"`
	BankName     string   `orm:"column(bank_name);type(text)" description:"Bank name" json:"bank_name"`
	IBAN         string   `orm:"column(iban);type(text)" description:"IBAN" json:"iban"`
	CurrencyCode string   `orm:"column(currency_code);type(text)" description:"Currency code" json:"currency_code"`
	IsDefault    int8     `orm:"column(is_default);size(1)" description:"Enabled 0:yes 1:no" json:"is_default"`
	models.Base
}

// TableName
func (*BankAccount) TableName() string {
	return "master_bank_account"
}

// SearchField
func (*BankAccount) SearchField() []string {
	return []string{"iban", "currency_code"}
}

// NoDeletionId
func (*BankAccount) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*BankAccount) WhereField() []string {
	return []string{}
}

// TimeField
func (*BankAccount) TimeField() []string {
	return []string{}
}
func  (*BankAccount)InitData(tx orm.TxOrmer)error {
	return nil
}
func (t *BankAccount) WhereCondition() *orm.Condition {
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
// init model
func init() {
	database.RegisterModel("master", new(BankAccount))
}

func (t *BankAccount) Export() []byte {
	var items []*BankAccount
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *BankAccount) Import(tx orm.TxOrmer, data []byte) error {
	var list []*BankAccount
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
