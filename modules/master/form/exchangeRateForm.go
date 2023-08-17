package form

import (
	"quince/internal/copier"
	"quince/modules/master/models"
)

// ExchangeRate struct
type ExchangeRateForm struct {
	Id           int64   `form:"id"`
	CurrencyCode string  `form:"currency_code"`
	Rate         float64 `form:"rate"`
}

func (c *ExchangeRateForm) Validate() (*models.ExchangeRate, error) {
	var m models.ExchangeRate
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		return &m, m.Validate()
	}
}
