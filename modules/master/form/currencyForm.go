package form

import (
	"quince/internal/copier"
	"quince/modules/master/models"
)

// CurrencyForm struct
type CurrencyForm struct {
	Id           int64  `form:"id"`
	Name         string `form:"name"`
	CurrencyCode string `form:"currency_code"`
	BaseCurrency int8   `form:"base_currency"`
}

func (c *CurrencyForm) Validate() (*models.Currency, error) {
	var m models.Currency
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		return &m, m.Validate()
	}
}
