package currency

import (
	"errors"
	m "quince/internal/models"
	"quince/modules/master/models"
	"quince/modules/master/services"

	"github.com/beego/beego/v2/client/orm"
)

type currencyUpdate struct {
	*models.Currency //unnamed parameter only !!!
}

func CurrencyUpdate(c *models.Currency) currencyUpdate {
	t := currencyUpdate{c}
	return t
}
func (u currencyUpdate) Run(txOrm orm.TxOrmer) error {
	err := u.Validate()
	if err != nil {
		return err
	}
	if u.BaseCurrency == 1 {
		currencyService := services.NewCurrencyService()
		currencyService.Tx = txOrm
		b := currencyService.ChangeBaseCurrency(u.Id)
		if b > 0 {
			return errors.New("master.currency.cannot_change")
		}
		//set other currencies on 0
		r := txOrm.Raw("UPDATE "+u.TableName()+" SET status=0 WHERE id !=?", u.Id)
		r.Exec()
	}
	return u.Update(txOrm, u.Currency)
}
func (u currencyUpdate) Description() string {
	return "Currency update"
}
func (u currencyUpdate) GetModel() m.IModel {
	return &u
}
