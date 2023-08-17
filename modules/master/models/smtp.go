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

type SMTP struct {
	Server   string   `orm:"column(server);type(text)" description:"Legal name" json:"server"  i18n:"master.smtp.server"`
	Port     string   `orm:"column(port);type(text)" description:"Legal name" json:"port"  i18n:"master.smtp.port"`
	User     string   `orm:"column(user);type(text)" description:"Legal name" json:"user"  i18n:"master.smtp.user"`
	Password string   `orm:"column(password);type(text)" description:"Legal name" json:"password"  i18n:"master.smtp.password"`
	UseTLS   int8     `orm:"column(useTLS);type(int)" description:"Legal name" json:"useTls"  i18n:"master.smtp.useTls"`
	Company  *Company `orm:"rel(fk)"`
	models.Base
}

func NewSMTP(companyID int64) *SMTP {
	smtp := SMTP{}
	smtp.Company = &Company{}
	smtp.Company.Id = companyID
	return &smtp
}

// TableName
func (*SMTP) TableName() string {
	return "master_smtp"
}

// init model
func init() {
	database.RegisterModel("master", new(SMTP))
}

func (*SMTP) SearchField() []string {
	return []string{"id", "server"}
}
func (cs *SMTP) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Server != "" {
			cond = cond.And("name__icontains", cs.Server)
		}
		if cs.Company != nil {
			if cs.Company.Id > 0 {
				cond = cond.And("company_id", cs.Company.Id)
			}
		}
	}
	return cond
}

// TimeField
func (*SMTP) TimeField() []string {
	return []string{}
}
func (c *SMTP) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Server, validation.Required, validation.Length(1, 250)))
	rules = append(rules, validation.Field(&c.Port, validation.Required, is.Digit))
	rules = append(rules, validation.Field(&c.User, validation.Required, validation.Length(1, 25)))
	rules = append(rules, validation.Field(&c.Password, validation.Required, validation.Length(1, 25)))
	rules = append(rules, validation.Field(&c.Company)) //
	err := validation.ValidateStruct(c, rules...)
	return err
}

func (t *SMTP) Export() []byte {
	var items []*SMTP
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *SMTP) Import(tx orm.TxOrmer, data []byte) error {
	var list []*SMTP
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
