package modal

import "quince/internal/toolbar/component"

type view struct {
	*component.Button
}

func newViewButton(name string,url string) component.IButton {
	a := view{}
	a.Button = new(component.Button)
	a.Icon = "fa fa-window-restore"
	a.Class = "btn btn-box-tool"
	a.Title = "app.details"
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.are_you_sure"
	a.DataConfirmTitle = "app.details"
	a.Enabled = true
	a.DataUrl = url // utils.LastSegment(url)
	a.DataConfirm = "2"
	a.DataType = "2"
	a.Href = ""
	a.AuthPath = url
	a.Name=name
	return a
}
