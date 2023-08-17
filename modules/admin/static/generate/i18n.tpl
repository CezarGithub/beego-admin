[{{$.Model.Module}}.{{$.Model.Name_lcase}}]
{{- range $item := .Model.Fields}}
{{$item.Name_lcase}} = 
{{- end}}