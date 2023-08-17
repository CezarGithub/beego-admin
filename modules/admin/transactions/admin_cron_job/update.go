package admin_cron_job

import (
	"errors"
	m "quince/internal/models"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type update struct {
	*models.AdminCronJob //unnamed parameter only !!!
}

func AdminCronJobUpdate(c *models.AdminCronJob) update {
	t := update{c}
	return t
}
func (u update) Run(txOrm orm.TxOrmer) error {
	adminCronJobService := services.NewAdminCronJobService()
	if adminCronJobService.IsExistName(strings.TrimSpace(u.Name), u.GetID()) {
		return errors.New("error.already_exists")
	}
	return u.Update(txOrm, u.AdminCronJob)
}

func (u update) Description() string {
	return "AdminCronJobInterval update"
}
func (u update) GetModel() m.IModel {
	return &u
}
