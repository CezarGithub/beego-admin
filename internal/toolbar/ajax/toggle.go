package ajax

import c "quince/internal/toolbar/component"

type toggle struct {
	*c.Button
}

func newToggleSelectedRowsButton(name string,url string) c.IButton {
	a := toggle{}
	a.Button = new(c.Button)
	a.Icon = "fa fa-toggle-on"
	a.Class = "btn  btn-box-tool AjaxButton"
	a.Title = "app.toggle_onoff"
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.toggle_onoff_question"
	a.DataConfirmTitle = "app.toggle_onoff"
	a.DataId = "checked"
	a.Enabled = true
	a.DataUrl = url
	a.AuthPath = url
	a.Name=name
	return a
}
