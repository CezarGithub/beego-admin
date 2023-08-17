package ajax

import c "quince/internal/toolbar/component"

type printgrid struct {
	*c.Button
}

func newPrintSelectedRowsButton(name string,url string) c.IButton {
	a := printgrid{}
	a.Button=new(c.Button)
	a.Icon = "fa fa-info"
	a.Class = "btn  btn-box-tool AjaxButton"
	a.Title = "app.print"
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.print_question"
	a.DataConfirmTitle = "app.print"
	a.DataId = "checked"
	a.Enabled = true
	a.DataUrl= url
	a.AuthPath=url
	a.Name=name
	return a
}
