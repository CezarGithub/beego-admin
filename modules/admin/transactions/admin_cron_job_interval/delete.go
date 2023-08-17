package admin_cron_job_interval

import (
	m "quince/internal/models"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

type delete struct {
	*models.AdminCronJobInterval //unnamed parameter only !!!
	idArr             []int64
}

func AdminCronJobIntervalDelete(idArr []int64) delete {
	c := models.AdminCronJobInterval{}
	t := delete{&c, idArr}
	return t
}
func (u delete) Run(txOrm orm.TxOrmer) error {
	_, err := txOrm.QueryTable(new(models.AdminCronJobInterval)).Filter("id__in", u.idArr).Delete()
	return err
}
func (u delete) Description() string {
	return "CronJobInterval delete"
}
func (u delete) GetModel() m.IModel {
	return &u
}
