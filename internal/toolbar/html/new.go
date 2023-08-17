package html

import "quince/internal/toolbar/component"

type newBtn struct {
	*component.Button
}

func newButton(name string,title string,icon string,url string) component.IButton {
	a := new(newBtn)
	a.Button=new(component.Button)
	a.Icon = icon
	a.Class = "btn  btn-box-tool" //btn-primary
	a.Title = title
	a.DataToggle = "tooltip"
	a.Enabled = true
	a.Href = url
	a.AuthPath = url
	a.Name=name
	return a
}
