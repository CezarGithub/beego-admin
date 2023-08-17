package exchangeRate


import (
	m "quince/internal/models"
	"quince/modules/master/models"

	"github.com/beego/beego/v2/client/orm"
)

type exchange_rateUpdate struct {
	*models.ExchangeRate //unnamed parameter only !!!
}

func ExchangeRateUpdate(c *models.ExchangeRate) exchange_rateUpdate {
	t := exchange_rateUpdate{c}
	return t
}
func (u exchange_rateUpdate) Run(txOrm orm.TxOrmer) error {
	return u.Update(txOrm, u.ExchangeRate)
}
func (u exchange_rateUpdate) Description() string {
	return "ExchangeRate update"
}
func (u exchange_rateUpdate) GetModel() m.IModel {
	return &u
}
