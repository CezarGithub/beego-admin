package ajax

import c "quince/internal/toolbar/component"

type submit struct {
	*c.Button
}

func newSubmitButton(name string,formName string, url string) c.IButton {
	a := submit{}
	a.Button = new(c.Button)
	a.Icon = "fa fa-save"
	a.Class = "btn btn-box-tool  dataFormSubmit"
	a.Title = "app.save"
	a.DataToggle = "tooltip"
	a.Enabled = true
	a.DataUrl = url
	a.AuthPath = url
	a.DataForm = formName
	a.Name=name
	return a
}
