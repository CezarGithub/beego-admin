package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// ExchangeRate struct
type ExchangeRate struct {
	CurrencyCode string  `orm:"column(currency_code);type(text)" description:"Currency code" json:"currency_code" i18n:"master.exchange_rate.currency_code"`
	Rate         float64 `orm:"column(rate);type(real)" description:"Exchange rate" json:"rate" i18n:"master.exchange_rate.rate"`
	//STRUCT ? Date 		 time.Time `orm:"column(date);type(datetime)" description:"Exchange rate date" json:"date" i18n:"master.exchange_rate.date"`
	// time.Parse(time.DateOnly, dateToParse)
	models.Base
}

func NewExchangeRate() ExchangeRate {
	exchange_rate := ExchangeRate{Base: models.Base{Id: 0}}
	return exchange_rate
}

// TableName
func (*ExchangeRate) TableName() string {
	return "master_exchange_rate"
}

// SearchField
func (*ExchangeRate) SearchField() []string {
	return []string{"currency_code"}
}

func (cs *ExchangeRate) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.CurrencyCode != "" {
			cond = cond.And("currency_code__icontains", cs.CurrencyCode)
		}
	}
	return cond
}

// TimeField
func (*ExchangeRate) TimeField() []string {
	return []string{"date"}
}

// init model
func init() {
	database.RegisterModel("master", new(ExchangeRate))
}
func (c *ExchangeRate) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Rate, validation.Required, validation.Min(0.01)))
	rules = append(rules, validation.Field(&c.CurrencyCode, validation.Required, validation.Length(1, 5)))
	err := validation.ValidateStruct(c, rules...)
	return err
}

func (t *ExchangeRate) Export() []byte {
	var items []*ExchangeRate
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *ExchangeRate) Import(tx orm.TxOrmer, data []byte) error {
	var list []*ExchangeRate
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
