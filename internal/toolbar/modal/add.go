package modal

import (
	c "quince/internal/toolbar/component"
)

type add struct {
	*c.Button
}

func newAddButton(name string,url string) c.IButton {
	a := add{}
	a.Button = new(c.Button)
	a.Icon = "fa fa-plus"
	a.Class = "btn btn-box-tool AjaxButton"
	a.Title = "app.add"
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.are_you_sure"
	a.DataConfirmTitle = "app.add"
	a.DataType = "2"
	a.DataConfirm = "2"
	a.Enabled = true
	a.DataUrl = url // utils.LastSegment(url)
	a.AuthPath = url
	a.Name = name
	return a
}
