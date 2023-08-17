package transactions

import (
	m "quince/internal/models"
	"quince/modules/master/models"

	"github.com/beego/beego/v2/client/orm"
)

type countyUpdate struct {
	*models.County //unnamed parameter only !!!
}

func CountyUpdate(c *models.County) countyUpdate {
	t := countyUpdate{c}
	return t
}
func (u countyUpdate) Run(txOrm orm.TxOrmer) error {
	return u.Update(txOrm, u.County)
}
func (u countyUpdate) Description() string {
	return "County update"
}
func (u countyUpdate) GetModel() m.IModel {
	return &u
}
