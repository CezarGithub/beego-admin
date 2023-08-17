package ajax

import c "quince/internal/toolbar/component"

type delete struct {
	*c.Button
}

func newDeleteSelectedRowsButton(name string,url string) c.IButton {
	a := delete{}
	a.Button=new(c.Button)
	a.Icon = "fa fa-trash"
	a.Class = "btn btn-box-tool AjaxButton"
	a.Title = "app.delete"
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.delete_question"
	a.DataConfirmTitle = "app.delete"
	a.DataId = "checked"
	a.Enabled = true
	a.DataUrl= url
	a.AuthPath =url
	a.Name=name
	return a
}

