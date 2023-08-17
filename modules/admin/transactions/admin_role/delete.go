package admin_role

import (
	m "quince/internal/models"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

type delete struct {
	*models.AdminRole //unnamed parameter only !!!
	idArr             []int64
}

func AdminRoleDelete(idArr []int64) delete {
	c := models.AdminRole{}
	t := delete{&c, idArr}
	return t
}
func (u delete) Run(txOrm orm.TxOrmer) error {
	_, err := txOrm.QueryTable(new(models.AdminRole)).Filter("id__in", u.idArr).Delete()
	return err
}
func (u delete) Description() string {
	return "Admin role delete"
}
func (u delete) GetModel() m.IModel {
	return &u
}
