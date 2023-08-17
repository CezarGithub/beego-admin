package admin_rbac

import (
	m "quince/internal/models"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

type delete struct {
	*models.AdminRbac //unnamed parameter only !!!
	idArr             []int64
}

func AdminRbacDelete(idArr []int64) delete {
	c := models.AdminRbac{}
	t := delete{&c, idArr}
	return t
}
func (u delete) Run(txOrm orm.TxOrmer) error {
	_, err := txOrm.QueryTable(new(models.AdminRbac)).Filter("id__in", u.idArr).Delete()
	return err
}
func (u delete) Description() string {
	return "Admin role basec acces control delete"
}
func (u delete) GetModel() m.IModel {
	return &u
}
