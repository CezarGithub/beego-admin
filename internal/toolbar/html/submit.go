package html

import c "quince/internal/toolbar/component"

type submit struct {
	*c.Button
}

func newSubmitButton(name string, formName string, url string) c.IButton {
	a := submit{}
	a.Button = new(c.Button)
	a.Icon = "fa fa-exclamation-circle"
	a.Class = "btn btn-box-tool  AjaxButton"
	a.Title = "?"
	a.DataToggle = "tooltip"
	a.Enabled = true
	a.DataUrl = url
	a.AuthPath = url
	a.DataForm = formName
	a.DataMethod = "GET" //!Important
	a.DataConfirm = "2"
	a.DataType = "2"
	a.Name = name
	return a
}
