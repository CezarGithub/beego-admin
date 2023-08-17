package user_level

import (
	"errors"
	m "quince/internal/models"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type update struct {
	*models.UserLevel //unnamed parameter only !!!
}

func UserLevelUpdate(c *models.UserLevel) update {
	t := update{c}
	return t
}
func (u update) Run(txOrm orm.TxOrmer) error {
	//Name check
	userLevelService := services.NewUserLevelService()
	if userLevelService.IsExistName(strings.TrimSpace(u.UserLevel.Name), u.GetID()) {
		return errors.New("error.already_exists")
	}
	return u.Update(txOrm, u.UserLevel)
}
func (u update) Description() string {
	return "User Level update"
}
func (u update) GetModel() m.IModel {
	return &u
}
