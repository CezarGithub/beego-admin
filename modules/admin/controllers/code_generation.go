package controllers

import (
	"fmt"
	"quince/initialize/database"
	"quince/initialize/module"
	"quince/internal/generate"
	"quince/internal/structs"
	"quince/internal/toolbar"
	"quince/modules/admin/services"
	"strings"
	"unicode"

	"github.com/beego/beego/v2/server/web"
)

// CodeGenerationController struct.
type CodeGenerationController struct {
	BaseController
}

func (dc *CodeGenerationController) Index() {
	var list []string
	modules := module.GetModulesNames()
	reqModule := dc.GetString("_module")
	databaseService := services.NewDatabaseService()
	_, affectRows := databaseService.GetTableStatus()
	models := database.RegisteredModels()

	for k := range models {
		m := strings.Split(k, ".")
		module := m[0]
		//name := m[1]
		if reqModule == module {
			list = append(list, k)
		}

	}
	dc.Data["data"] = list
	dc.Data["total"] = affectRows
	dc.Data["module"] = reqModule
	dc.Data["modules"] = modules
	save := toolbar.Html("generate.new").Submit("generateForm", web.URLFor("CodeGenerationController.New"))
	save.SetTitle("admin.new_model")
	save.SetIcon("fa fa-file-code-o")
	dc.AddButtons(save)
	dc.TplName = "admin/views/generate/index.html"
}

func (dc *CodeGenerationController) New() {
	module := dc.GetString("_module")
	name := dc.GetString("_name")
	data := new(generate.ModelData)
	data.Name_lcase = lCase(name)
	data.Module = module
	data.Name = name
	data.Template = "iModel.tpl" // model.tpl require a registered model
	dc.Data["model"] = data
	dc.TplName = "admin/views/generate/view.html"

}
func (dc *CodeGenerationController) Doc() {
	//n := strings.Split(dc.GetString("_name", ""), ".")
	name := dc.GetString("_name", "")
	doc := dc.GetString("_doc", "")
	module := strings.Split(name, ".")[0]
	//data := new(generate.ModelData)
	// if len(n) != 2 {
	// 	data.Template = "error.tpl"
	// } else {
	// 	module := n[0]
	// 	name := n[1]
	// 	data.Module = module
	// 	data.Name = name
	// 	data.Template = doc + ".tpl"
	// }
	data := getData(name)
	data.Module = module
	data.Template = doc // + ".tpl"
	dc.Data["model"] = data
	dc.TplName = "admin/views/generate/view.html"
}

func getData(name string) *generate.ModelData {
	models := database.RegisteredModels()
	m := models[name]
	modelData := new(generate.ModelData)
	modelData.Fields = *new([]generate.Field)
	myModel := structs.New(m)
	fields := myModel.Fields()
	for _, f := range fields {
		fName := f.Name()
		kind := f.Kind().String()
		name_lcase := lCase(fName)
		orm := ormType(name_lcase, kind)
		i18n := lCase(name) + "." + name_lcase
		newField := generate.Field{Name: fName, Name_lcase: name_lcase, Orm: orm, I18n: i18n, Type: kind}
		if fName != "Base" {
			modelData.Fields = append(modelData.Fields, newField)
		}
	}
	// for k := range models {
	// 	m := strings.Split(k, ".")
	// 	module := m[0]
	// 	//name := m[1]
	// 	if reqModule == module {
	// 		list = append(list, k)
	// 	}
	modelData.Name = myModel.Name()
	modelData.Name_lcase = lCase(modelData.Name)

	fmt.Printf("%#v", myModel)
	return modelData

}

func lCase(s string) string {
	result := ""
	for i, r := range s {
		if i == 0 {
			result += strings.ToLower(string(r))
		} else {
			if unicode.IsUpper(r) && !unicode.IsNumber(r) {
				result += "_"
			}
			result += strings.ToLower(string(r))
		}
	}
	return result
}
func ormType(name_lcase string, kind string) string {
	var result string
	switch kind {
	case "int8":
		result = fmt.Sprintf("orm:\"column(%s);size(1);default(0)\"", name_lcase)
	case "int64":
		result = fmt.Sprintf("orm:\"column(%s);size(11);default(1)\"", name_lcase)
	case "int":
		result = fmt.Sprintf("orm:\"column(%s);size(8);default(1)\"", name_lcase)
	case "string":
		result = fmt.Sprintf("orm:\"column(%s);type(text)\"", name_lcase)
	case "time.Time":
		result = fmt.Sprintf("orm:\"column(%s);auto_now;type(datetime)\"", name_lcase)
	case "ptr":
		result = "`orm:\"null;rel(one);on_delete(set_null)\"` OR `orm:\"rel(fk)\"` OR `orm:\"reverse(many)\"` - check Beego docs "
	default:
		result = ""
	}
	return result
}
