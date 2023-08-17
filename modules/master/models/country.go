package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Country struct
type Country struct {
	Name    string    `orm:"column(name);unique;type(text)" description:"Legal name" json:"name" i18n:"master.country.name"`
	Alpha_2 string    `orm:"column(alpha_2);type(text)" description:"Country code 2 char" json:"alpha_2"  i18n:"master.country.alpha_2"`
	Alpha_3 string    `orm:"column(alpha_3);type(text)" description:"Country code 3 char" json:"alpha_3"  i18n:"master.country.alpha_3"`
	County  []*County `orm:"reverse(many)"` //OneToMany
	models.Base
}
// TableName
func (*Country) TableName() string {
	return "master_country"
}

// SearchField
func (*Country) SearchField() []string {
	return []string{"id", "name", "alpha_2", "alpha_3"}
}
func (*Country) TableIndex() [][]string {
	return [][]string{
		{"Id"},
	}
}
func (t *Country) WhereCondition() *orm.Condition {
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

// TimeField
func (*Country) TimeField() []string {
	return []string{}
}

//init model
func init() {
	database.RegisterModel("master", new(Country))

}

func (c *Country) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Alpha_2, validation.Required, validation.Length(2, 2)))
	rules = append(rules, validation.Field(&c.Alpha_3, validation.Required, validation.Length(3, 3)))

	err := validation.ValidateStruct(c, rules...)
	return err
}

func (t *Country) Export() []byte {
	var items []*Country
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Country) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Country
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

