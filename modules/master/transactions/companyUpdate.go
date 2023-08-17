package transactions

import (
	m "quince/internal/models"
	"quince/modules/master/models"

	"github.com/beego/beego/v2/client/orm"
)

type companyUpdate struct {
	*models.Company //unnamed parameter only !!!
}

func CompanyUpdate(c *models.Company) companyUpdate {
	t := companyUpdate{c}
	return t
}
func (u companyUpdate) Run(txOrm orm.TxOrmer) error {
	return u.Update(txOrm, u.Company)
}
func (u companyUpdate) Description() string {
	return "Company update"
}
func (u companyUpdate) GetModel() m.IModel {
	return &u
}
