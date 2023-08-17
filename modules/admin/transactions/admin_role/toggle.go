package admin_role

import (
	m "quince/internal/models"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

type toggle struct {
	*models.AdminRole //unnamed parameter only !!!
	idArr             []int64
}

func AdminRoleToggleStatus(idArr []int64) toggle {
	c := models.AdminRole{}
	t := toggle{&c, idArr}
	return t
}
func (u toggle) Run(txOrm orm.TxOrmer) error {
	r := txOrm.Raw("UPDATE "+u.TableName()+" SET status=ABS(status-1) WHERE id IN (?)", u.idArr)
	_, e := r.Exec()
	return e
}
func (u toggle) Description() string {
	return "Admin role toggle status"
}
func (u toggle) GetModel() m.IModel {
	return &u
}
