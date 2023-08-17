package modules

import (
	_ "quince/modules/admin"
	_ "quince/modules/crm"
	_ "quince/modules/master"

	"github.com/beego/beego/v2/core/logs"
)

//Init - init all modules
func init() {
	logs.Info("Modules initialisation")
}
