package funcmap

import (
	"quince/internal/generate"
	"quince/internal/i18n"
	"quince/internal/template"
	"quince/internal/toolbar"

	"github.com/beego/beego/v2/server/web"
)

func init() {

	//register template function i18n usage
	// {{i18n .admin.user.Language "app.server_os"}}
	//web.AddFuncMap("i18n", i18n.Tr)

	//dummy function registered.Real function registered in- baseController.go :110 = web.AddFuncMap("i18n", bc.Translate)
	//{{i18n "app.server_os"}} --> .admin.user.Language param is gone
	web.AddFuncMap("i18n", func(s string) string { return "" })

	web.AddFuncMap("flag", i18n.LanguageFlag)
	//multiple args in template calls - not used ,yet
	web.AddFuncMap("args", template.TemplateArgs)
	//button funcmap
	web.AddFuncMap("button", toolbar.RenderButton)
	//code generation funcmap
	web.AddFuncMap("generate", generate.FuncMap)

}
