package html

import "quince/internal/toolbar/component"

type view struct {
	*component.Button
}

func newViewButton(name string, url string) component.IButton {
	a := view{}
	a.Button = new(component.Button)
	a.Icon = "fa fa-info-circle"
	a.Class = "btn btn-box-tool"
	a.Title = "app.details"
	a.DataToggle = "tooltip"
	a.Enabled = true
	a.Href = url
	a.AuthPath = url
	a.Name = name
	return a
}
