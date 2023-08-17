package transactions

import (
	m "quince/internal/models"
	"quince/modules/master/models"

	"github.com/beego/beego/v2/client/orm"
)

type countryUpdate struct {
	*models.Country //unnamed parameter only !!!
}

func CountryUpdate(c *models.Country) countryUpdate {
	t := countryUpdate{c}
	return t
}
func (u countryUpdate) Run(txOrm orm.TxOrmer) error {
	return u.Update(txOrm, u.Country)
}
func (u countryUpdate) Description() string {
	return "Country update"
}
func (u countryUpdate) GetModel() m.IModel{
	return &u
}