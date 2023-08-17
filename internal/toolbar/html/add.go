package html

import "quince/internal/toolbar/component"

type add struct {
	*component.Button
}

func newAddButton(name string,url string) component.IButton {
	a := new(add)
	a.Button=new(component.Button)
	a.Icon = "fa fa-plus"
	a.Class = "btn  btn-box-tool" //btn-primary
	a.Title = "app.add"
	a.DataToggle = "tooltip"
	a.Enabled = true
	a.Href = url
	a.AuthPath = url
	a.Name=name
	return a
}
