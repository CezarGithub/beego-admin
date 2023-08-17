package admin_rbac

import (
	"quince/modules/admin/models"
	m "quince/internal/models"
	"github.com/beego/beego/v2/client/orm"
)

type toggle struct {
	*models.AdminRbac //unnamed parameter only !!!
	idArr             []int
}

func AdminRbacToggleStatus(idArr []int) toggle {
	c := models.AdminRbac{}
	t := toggle{&c, idArr}
	return t
}
func (u toggle) Run(txOrm orm.TxOrmer) error {
	r := txOrm.Raw("UPDATE "+u.TableName()+" SET status=ABS(status-1) WHERE id IN (?)", u.idArr)
	_, e := r.Exec()
	return e
}
func (u toggle) Description() string {
	return "Admin rbac toggle status"
}
func (u toggle) GetModel() m.IModel {
	return &u
}