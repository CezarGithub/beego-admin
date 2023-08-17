package admin_cron_job_interval

import (
	m "quince/internal/models"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

type update struct {
	*models.AdminCronJobInterval //unnamed parameter only !!!
}

func AdminCronJobIntervalUpdate(c *models.AdminCronJobInterval) update {
	t := update{c}
	return t
}
func (u update) Run(txOrm orm.TxOrmer) error {
	err := u.Validate()
	if err != nil {
		return err
	} else {
		return u.Update(txOrm, u.AdminCronJobInterval)
	}
}
func (u update) Description() string {
	return "AdminCronJobInterval update"
}
func (u update) GetModel() m.IModel {
	return &u
}
