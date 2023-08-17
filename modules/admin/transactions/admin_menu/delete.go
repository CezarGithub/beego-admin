package admin_menu

import (
	"errors"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	m "quince/internal/models"
	"github.com/beego/beego/v2/client/orm"
)

type delete struct {
	*models.AdminMenu //unnamed parameter only !!!
	idArr             []int64
}

func AdminMenuDelete(idArr []int64) delete {
	c := models.AdminMenu{}
	t := delete{&c, idArr}
	return t
}
func (u delete) Run(txOrm orm.TxOrmer) error {
	adminMenuService := services.NewAdminMenuService()
	//Determine whether there is a submenu
	if adminMenuService.IsChildMenu(u.idArr) {
		return errors.New("error.cannot_delete")
	}

	_, err := txOrm.QueryTable(new(models.AdminMenu)).Filter("id__in", u.idArr).Delete()
	return err
}
func (u delete) Description() string {
	return "Admin menu delete"
}
func (u delete) GetModel() m.IModel {
	return &u
}
