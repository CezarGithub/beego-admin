package i18n

import (
	"path"
	"quince/internal/i18n"
	"quince/utils/file"

	"github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/core/logs"
)

func AddLocaleFiles(folderPath string) {
	// l, _ := web.AppConfig.String("lang::alpha4")
	// langs := strings.Split(l, "|")
	langs := i18n.ApplicationLanguages()
	folder, _ := web.AppConfig.String("lang::folder")
	for _, lang := range langs {
		local := path.Join(folderPath, "locale_"+lang+".ini")
		dest := path.Join(folder, "locale_"+lang+".ini")
		//logs.Info("Loading language: %s - %s -> %s ", lang, local, dest)
		err := file.AppendToFile(local, dest)
		if err != nil {
			logs.Error("Module admin load language files : %v", err.Error())
		}
	}
}
