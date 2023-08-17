package admin_menu

import (
	"errors"
	m "quince/internal/models"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type update struct {
	*models.AdminMenu //unnamed parameter only !!!
}

func AdminMenuUpdate(c *models.AdminMenu) update {
	t := update{c}
	return t
}
func (u update) Run(txOrm orm.TxOrmer) error {
	adminMenuService := services.NewAdminMenuService()
	if adminMenuService.IsExistUrl(strings.TrimSpace(u.AdminMenu.Name), u.GetID()) {
		return errors.New("error.already_exists")
	}
	return u.Update(txOrm, u.AdminMenu)
}

func (u update) Description() string {
	return "Login user update"
}
func (u update) GetModel() m.IModel {
	return &u
}
