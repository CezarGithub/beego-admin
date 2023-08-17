package admin_menu

import (
	"quince/modules/admin/models"
	m "quince/internal/models"
	"github.com/beego/beego/v2/client/orm"
)

type toggle struct {
	*models.AdminMenu //unnamed parameter only !!!
	idArr             []int64
}

func AdminMenuToggleStatus(idArr []int64) toggle {
	c := models.AdminMenu{}
	t := toggle{&c, idArr}
	return t
}
func (u toggle) Run(txOrm orm.TxOrmer) error {
	r := txOrm.Raw("UPDATE "+u.TableName()+" SET is_show=ABS(is_show-1) WHERE id IN (?)", u.idArr)
	_, e := r.Exec()
	return e
}
func (u toggle) Description() string {
	return "Admin menu toggle status"
}
func (u toggle) GetModel() m.IModel {
	return &u
}