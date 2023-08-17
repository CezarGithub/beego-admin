package module

import (
	"context"
	"path/filepath"
	"quince/utils/file"
	"quince/utils/i18n"
	s "quince/utils/string"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/client/orm"

	"github.com/beego/beego/v2/server/web"
)

var modulesList *Modules

type Modules struct {
	Items map[string]IModule
}

func init() {
	modulesList = new(Modules)
	modulesList.Items = make(map[string]IModule)
}
func Register(name string, module IModule) {
	name = s.Alphanumeric(name)
	modulesList.Items[name] = module
	modulesPath, err := web.AppConfig.String("modulesFolder")
	if err != nil {
		modulesPath = "modules"
	}
	i18nPath := filepath.Join(modulesPath, name, "static", "i18n")
	i18n.AddLocaleFiles(i18nPath)
	//data init files
	initFilesDest := filepath.Join("static", "data", "init", name)
	initFilesSource := filepath.Join(modulesPath, name, "static", "data", "init")
	file.CopyDir(initFilesSource, initFilesDest)
}
func GetModules() *Modules {
	return modulesList
}
func GetModulesNames() []string {
	var list []string
	for k := range modulesList.Items {
		list = append(list, k)
	}
	return list
}
func DBSync() {
	o := orm.NewOrm()
	var er error
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		for _, v := range modulesList.Items {
			er = v.Init(txOrm)
		}
		return er
	})
	if err != nil {
		logs.Error(err)
	}

}
