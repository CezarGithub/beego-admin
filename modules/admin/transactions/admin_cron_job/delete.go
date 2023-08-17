package admin_cron_job

import (
	m "quince/internal/models"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

type delete struct {
	*models.AdminCronJob //unnamed parameter only !!!
	idArr             []int64
}

func AdminCronJobDelete(idArr []int64) delete {
	c := models.AdminCronJob{}
	t := delete{&c, idArr}
	return t
}
func (u delete) Run(txOrm orm.TxOrmer) error {
	_, err := txOrm.QueryTable(new(models.AdminCronJob)).Filter("id__in", u.idArr).Delete()
	return err
}
func (u delete) Description() string {
	return "CronJob delete"
}
func (u delete) GetModel() m.IModel {
	return &u
}
