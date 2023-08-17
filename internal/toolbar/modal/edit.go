package modal

import (
	c "quince/internal/toolbar/component"
)

type edit struct {
	*c.Button
}

func newEditButton(name string,url string) c.IButton {
	a := edit{}
	a.Button = new(c.Button)
	a.Icon = "fa fa-pencil"
	a.Class = "btn btn-box-tool AjaxButton"
	a.Title = "app.edit"
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.are_you_sure"
	a.DataConfirmTitle = "app.edit"
	a.Enabled = true
	a.DataUrl = url // utils.LastSegment(url)
	a.DataConfirm="2"
	a.DataType="2"
	a.AuthPath = url
	a.Name=name
	return a
}
