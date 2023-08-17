package crm

import (
	_ "quince/modules/crm/init/i18n"
	//_ "quince/modules/crm/models"
	_ "quince/modules/crm/routers"

	"github.com/beego/beego/v2/core/logs"
)

func init() {
	logs.Info("CRM module init")

}
