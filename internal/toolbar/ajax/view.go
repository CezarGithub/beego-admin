package ajax

import (
	c "quince/internal/toolbar/component"
)

type view struct {
	*c.Button
}

func newViewButton(name string, url string) c.IButton {
	a := view{}
	a.Button = new(c.Button)
	a.Icon = "fa fa-info-circle"
	a.Class = "btn btn-box-tool AjaxButton"
	a.Title = "app.details"
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.are_you_sure"
	a.DataConfirmTitle = "app.details"
	a.Enabled = true
	//a.DataUrl, _ = utils.LastSegment(url) // ????
	a.DataUrl = url
	a.AuthPath = url
	a.Name = name
	//a.Href = url
	return a
}
