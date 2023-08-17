package form


import (
	"quince/internal/copier"
	"quince/modules/{{.Model.Module}}/models"
)

// {{.Model.Name}}Form struct
type {{.Model.Name}}Form struct {
	Id          int64  `form:"id"`
	{{- range $item := .Model.Fields}}
	  {{- if eq $item.Type "ptr" }}
		{{$item.Name}}ID		int64 `form:"{{$item.Name_lcase}}_id"`
	  {{- else}}
		{{$item.Name}}			{{$item.Type}}  `form:"{{$item.Name_lcase}}"`
	  {{- end}}
	{{- end}}
}

func (c *{{.Model.Name}}Form) Validate() (*models.{{.Model.Name}}, error) {
	var m models.{{.Model.Name}}
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		{{- range $item := .Model.Fields}}
			{{- if eq $item.Type "ptr" }}
			m.{{$item.Name}} = &models.{{$item.Name}}{}
			m.{{$item.Name}}.Id = c.{{$item.Name}}ID
			{{- end}}
		{{- end}}
		return &m, m.Validate()
	}
}