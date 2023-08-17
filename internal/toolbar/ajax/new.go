package ajax

import (
	c "quince/internal/toolbar/component"
)

type newBtn struct {
	*c.Button
}

func newButton(name string, title string, icon string, url string, isPopup bool) c.IButton {
	a := newBtn{}
	a.Button = new(c.Button)
	a.Icon = icon
	a.Class = "btn btn-box-tool AjaxButton"
	a.Title = title
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.are_you_sure"
	a.DataConfirmTitle = title
	if isPopup {
		a.DataType = "2"
	} else {
		a.DataType = "1"
	}
	a.DataConfirm = "2"
	a.Enabled = true
	a.DataUrl = url // utils.LastSegment(url)
	a.AuthPath = url
	a.Name = name
	return a
}
