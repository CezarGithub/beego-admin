package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/internal/validation/is"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Company struct
type Company struct {
	Name        string   `orm:"column(name);unique;type(text)" description:"Legal name" json:"name" i18n:"master.company.name"`
	Mark        string   `orm:"column(mark);unique;type(text)" description:"Commercial name" json:"mark" i18n:"master.company.mark"`
	Code        string   `orm:"column(code);unique;size(3)" description:"Unique 3 chars code" json:"code" i18n:"master.company.code"`
	VAT         string   `orm:"column(vat);unique;type(text)" description:"VAT number" json:"vat" i18n:"master.company.vat"`
	RegNo       string   `orm:"column(regno);unique;type(text)" description:"Registration number ,if any" json:"regno" i18n:"master.company.regno"`
	Web         string   `orm:"column(web);type(text)" description:"Oficial site" json:"web" i18n:"master.company.web"`
	Email       string   `orm:"column(email);type(text)" description:"Oficial email address" json:"email" i18n:"master.company.email"`
	Country     *Country `orm:"null;rel(one);on_delete(set_null)"`
	Group       *Group   `orm:"null;rel(one);on_delete(set_null)"`
	Status      int8     `orm:"column(status);size(1);default(1)" description:"Status 0:disabled 1:enabled" json:"status" i18n:"master.company.status"`
	Description string   `orm:"column(description);type(text)" description:"Comments" json:"description" i18n:"master.company.description"`
	models.Base
}

// TableName
func (*Company) TableName() string {
	return "master_company"
}

// SearchField
func (*Company) SearchField() []string {
	return []string{"name", "vat", "email", "web"}
}

func (cs *Company) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Name != "" {
			cond = cond.And("name__icontains", cs.Name)
		}
		if cs.VAT != "" {
			cond = cond.And("vat__icontains", cs.VAT)
		}
		if cs.Web != "" {
			cond = cond.And("web__icontains", cs.Web)
		}
		if cs.Email != "" {
			cond = cond.And("email__icontains", cs.Email)
		}
	}
	return cond
}

// TimeField
func (*Company) TimeField() []string {
	return []string{}
}

// init model
func init() {
	database.RegisterModel("master", new(Company))
}

func (c *Company) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Code, validation.Required, validation.Length(1, 25)))
	rules = append(rules, validation.Field(&c.VAT, validation.Required, validation.Length(1, 25)))
	rules = append(rules, validation.Field(&c.Email, is.Email))
	rules = append(rules, validation.Field(&c.Web, is.URL))
	err := validation.ValidateStruct(c, rules...)
	return err
}
func (t *Company) Export() []byte {
	var items []*Company
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Company) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Temp
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
