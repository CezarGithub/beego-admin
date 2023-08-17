package admin_rbac

import (
	"errors"
	m "quince/internal/models"
	"quince/modules/admin/models"
	"quince/modules/admin/services"

	"github.com/beego/beego/v2/client/orm"
)

type update struct {
	*models.AdminRbac //unnamed parameter only !!!
}

func AdminRbacUpdate(c *models.AdminRbac) update {
	t := update{c}
	return t
}
func (u update) Run(txOrm orm.TxOrmer) error {
	//Name check
	adminRbacService := services.NewAdminRbacService()
	b := adminRbacService.IsExistUrl(u.Role.Id, u.Route.Id)
	if b != nil {
		return errors.New("error.already_exists")
	}
	return u.Update(txOrm, u.AdminRbac)
}
func (u update) Description() string {
	return "Admin rbac update"
}
func (u update) GetModel() m.IModel {
	return &u
}
