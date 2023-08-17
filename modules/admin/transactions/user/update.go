package user

import (
	m "quince/internal/models"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

type update struct {
	*models.User //unnamed parameter only !!!
}

func UserUpdate(c *models.User) update {
	t := update{c}
	return t
}
func (u update) Run(txOrm orm.TxOrmer) error {
	//Name check
	//userService := services.NewUserService()
	// if userService.IsExistName(strings.TrimSpace(u.UserLevel.Name), u.GetID()) {
	// 	return errors.New("error.already_exists")
	// }
	return u.Update(txOrm, u.User)
}
func (u update) Description() string {
	return "User update"
}
func (u update) GetModel() m.IModel {
	return &u
}
