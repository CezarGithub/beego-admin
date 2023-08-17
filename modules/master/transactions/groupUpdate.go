package transactions

import (
	m "quince/internal/models"
	"quince/modules/master/models"

	"github.com/beego/beego/v2/client/orm"
)

type groupUpdate struct {
	*models.Group //unnamed parameter only !!!
}

func GroupUpdate(c *models.Group) groupUpdate {
	t := groupUpdate{c}
	return t
}
func (u groupUpdate) Run(txOrm orm.TxOrmer) error {
	return u.Update(txOrm, u.Group)
}
func (u groupUpdate) Description() string {
	return "Group update"
}
func (u groupUpdate) GetModel() m.IModel {
	return &u
}
