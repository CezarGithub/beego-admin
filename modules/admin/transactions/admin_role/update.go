package admin_role

import (
	"errors"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"strings"
	m "quince/internal/models"
	"github.com/beego/beego/v2/client/orm"
)

type update struct {
	*models.AdminRole //unnamed parameter only !!!
}

func AdminRoleUpdate(c *models.AdminRole) update {
	t := update{c}
	return t
}
func (u update) Run(txOrm orm.TxOrmer) error {
	//Name check
	adminRoleService := services.NewAdminRoleService()
	if adminRoleService.IsExistName(strings.TrimSpace(u.AdminRole.Name), u.GetID()) {
		return errors.New("error.already_exists")
	}
	return u.Update(txOrm, u.AdminRole)
}
func (u update) Description() string {
	return "Admin role update"
}
func (u update) GetModel() m.IModel {
	return &u
}