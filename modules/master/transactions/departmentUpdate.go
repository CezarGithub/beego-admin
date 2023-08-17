package transactions

import (
	m "quince/internal/models"
	"quince/modules/master/models"

	"github.com/beego/beego/v2/client/orm"
)

type departmentUpdate struct {
	*models.Department //unnamed parameter only !!!
}

func DepartmentUpdate(c *models.Department) departmentUpdate {
	t := departmentUpdate{c}
	return t
}
func (u departmentUpdate) Run(txOrm orm.TxOrmer) error {
	err := u.Validate()
	if err != nil {
		return err
	} else {
		return u.Update(txOrm, u.Department)
	}
}
func (u departmentUpdate) Description() string {
	return "Department update"
}
func (u departmentUpdate) GetModel() m.IModel {
	return &u
}
