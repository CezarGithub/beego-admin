package ajax

import c "quince/internal/toolbar/component"

type print struct {
	*c.Button
}

func newPrintButton(name string,title string,url string) c.IButton {
	a := print{}
	a.Button=new(c.Button)
	a.Icon = "fa fa-print"
	a.Class = "btn btn-box-tool  AjaxButton"
	a.Title = title
	a.DataToggle = "tooltip"
	a.Enabled = true
	a.Href = url
	a.AuthPath = url
	a.Name=name
	return a
}
