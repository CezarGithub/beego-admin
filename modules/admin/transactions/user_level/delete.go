package user_level

import (
	m "quince/internal/models"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

type delete struct {
	*models.UserLevel //unnamed parameter only !!!
	idArr             []int64
}

func UserLevelDelete(idArr []int64) delete {
	c := models.UserLevel{}
	t := delete{&c, idArr}
	return t
}
func (u delete) Run(txOrm orm.TxOrmer) error {
	_, err := txOrm.QueryTable(new(models.UserLevel)).Filter("id__in", u.idArr).Delete()
	return err
}
func (u delete) Description() string {
	return "User Level delete"
}
func (u delete) GetModel() m.IModel {
	return &u
}
