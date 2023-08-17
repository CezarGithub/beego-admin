package toolbar

import (
	"bytes"
	"html/template"
	"os"
	c "quince/internal/toolbar/component"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

type tplData struct {
	Button c.IButton
	Item   interface{}
}

// RenderButton - funcmap - returns toolbar buttons for grid rows - see views\public\toolbarRow.html
func RenderButton(name string, buttons []c.IButton, item interface{}) template.HTML {
	var buf bytes.Buffer
	var current c.IButton
	var tpl *template.Template
	var err error
	var new string

	for _, t := range buttons {
		if t.GetName() == name {
			current = t
			break
		}
	}
	if current == nil {
		logs.Error("%s - Button rendering WRONG name !!!!! \n %v", name, buttons)
	}
	v := current.HasVariableData()
	if item == nil {
		v = false
	}
	//if data-data has variable, insert variables in template before parse
	if v { //contains variable data and variable data is not nil
		content, err := os.ReadFile("modules/admin/views/public/button.html")
		if err != nil {
			logs.Error(err)
		}
		str := string(content)
		currentData := current.GetData()
		if len(currentData) > 0 {
			newData := currentData[0:10] + "'" + currentData[10:] + "'"    // 0:10 -> data-data=
			new = strings.ReplaceAll(str, "{{.Button.DataData}}", newData) //{{.Button.DataData}} from button.html template
			str=new
		}
		params := current.GetUrlParams()
		if len(params) > 0 {
			new = strings.ReplaceAll(str, "{{.Button.Params}}", params)
		}
		tpl, err = template.New("button.html").Parse(new)
		if err != nil {
			logs.Error(err)
		}
	} else {
		tpl, err = template.New("button.html").ParseFiles("modules/admin/views/public/button.html")
		if err != nil {
			logs.Error(err)
		}
	}

	temp := template.Must(tpl, nil)
	data := tplData{Button: current, Item: item}
	//err = temp.Execute(os.Stdout, data)//console
	err = temp.Execute(&buf, data)
	if err != nil {
		logs.Error(err)
	}
	s := buf.String()
	return template.HTML(s)
}

//tpl, err := template.New("toolbar.html").Funcs(funcMap).ParseFiles("modules/admin/views/public/toolbar.html")
