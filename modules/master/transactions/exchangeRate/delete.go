package exchangeRate

import (
	m "quince/internal/models"
	"quince/modules/master/models"

	"github.com/beego/beego/v2/client/orm"
)

type delete struct {
	*models.ExchangeRate //unnamed parameter only !!!
	idArr                []int64
	currencyCode         []string
}

// idArr har priority if both arguments are supplied
func ExchangeRateDelete(idArr []int64, currencyCode []string) delete {
	c := models.ExchangeRate{}
	t := delete{&c, idArr, currencyCode}
	return t
}
func (u delete) Run(txOrm orm.TxOrmer) error {
	var err error
	if len(u.idArr) > 0 {
		_, err = txOrm.QueryTable(new(models.ExchangeRate)).Filter("id__in", u.idArr).Delete()
	} else {
		_, err = txOrm.QueryTable(new(models.ExchangeRate)).Filter("currency_code__in", u.currencyCode).Delete()
	}
	return err
}
func (u delete) Description() string {
	return "ExchangeRate  delete"
}
func (u delete) GetModel() m.IModel {
	return &u
}
