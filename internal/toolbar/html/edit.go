package html

import "quince/internal/toolbar/component"

type edit struct {
	*component.Button
}

func newEditButton(name string,url string) component.IButton {
	a := edit{}
	a.Button = new(component.Button)
	a.Icon = "fa fa-pencil"
	a.Class = "btn btn-box-tool"
	a.Title = "app.edit"
	a.DataToggle = "tooltip"
	a.Enabled = true
	a.Href = url
	a.AuthPath = url
	a.Name=name
	return a
}
