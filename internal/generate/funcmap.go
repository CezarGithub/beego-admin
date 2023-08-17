package generate

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/beego/beego/v2/core/logs"
)

type Field struct {
	I18n       string
	Type       string
	Name       string
	Name_lcase string
	Orm        string
}
type ModelData struct {
	Template   string
	Name       string
	Module     string
	Name_lcase string
	Fields     []Field
	Type       string
}
type templateData struct {
	Model ModelData
}

// Generate - funcmap - returns generatet code
func FuncMap(model ModelData) template.HTML {
	var buf bytes.Buffer
	var tpl *template.Template
	var err error

	file := fmt.Sprintf("%s%s", "modules/admin/static/generate/", model.Template)
	ext := filepath.Ext(file)
	if ext != ".tpl" { //for html use [[]] delims to allow {{}} marks for go template
		tpl, err = template.New(model.Template).Delims("[[", "]]").ParseFiles(file)
	} else {
		tpl, err = template.New(model.Template).ParseFiles(file)
	}
	if err != nil {
		logs.Error(err)
	}
	temp := template.Must(tpl, nil)

	data := templateData{Model: model}
	//err = temp.Execute(os.Stdout, data)//console
	err = temp.Execute(&buf, data)
	if err != nil {
		logs.Error(err)
	}
	s := buf.String()
	return template.HTML(s)
}
