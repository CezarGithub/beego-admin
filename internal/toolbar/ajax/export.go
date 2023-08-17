package ajax

import (
	"fmt"
	c "quince/internal/toolbar/component"
)

type export struct {
	*c.Button
}

func newExportGridButton(name string,url string) c.IButton {
	a := export{}
	a.Button=new(c.Button)
	a.Icon = "fa fa-download"
	a.Class = "btn btn-box-tool exportData"
	a.Title = "app.export"
	a.DataToggle = "tooltip"
	a.DataConfirmContent = "app.export_question"
	a.DataConfirmTitle = "app.export"
	a.DataId = "checked"
	a.Enabled = true
	a.DataUrl = ""
	a.AuthPath = url
	a.Name=name
	a.OnClick = fmt.Sprintf("exportData(%s);", url)
	return a
}
