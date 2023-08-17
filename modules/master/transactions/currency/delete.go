package currency

import (
	m "quince/internal/models"
	"quince/modules/master/models"
	"quince/modules/master/transactions/exchangeRate"

	"github.com/beego/beego/v2/client/orm"
)

type delete struct {
	*models.Currency //unnamed parameter only !!!
	idArr            []int64
}

func CurrencyDelete(idArr []int64) delete {
	c := models.Currency{}
	t := delete{&c, idArr}
	return t
}
func (u delete) Run(txOrm orm.TxOrmer) error {
	var result []models.Currency
	setter := txOrm.QueryTable(new(models.Currency)).Filter("id__in", u.idArr)
	_, err := setter.All(&result, "currency_code")
	if err == nil {
		_, err = setter.Delete()
		var codeArr []string
		for _,v :=range result{
			codeArr = append(codeArr,v.CurrencyCode )
		}
		delete := exchangeRate.ExchangeRateDelete(nil, codeArr)
		delete.Run(txOrm)
	}
	return err
}
func (u delete) Description() string {
	return "Currency  delete"
}
func (u delete) GetModel() m.IModel {
	return &u
}
