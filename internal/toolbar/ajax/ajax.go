package ajax

import (
	c "quince/internal/toolbar/component"
)

type Ajax struct {
	Name string
}

func (a *Ajax) Export(url string) c.IButton {
	return newExportGridButton(a.Name, url)
}
func (a *Ajax) View(url string) c.IButton {
	return newViewButton(a.Name, url)
}
func (a *Ajax) DeleteSelectedRows(url string) c.IButton {
	return newDeleteSelectedRowsButton(a.Name, url)
}
func (a *Ajax) Print(title string, url string) c.IButton {
	return newPrintButton(a.Name, title, url)
}
func (a *Ajax) New(title string, icon string, url string, isPopup bool) c.IButton {
	return newButton(a.Name, title, icon, url, isPopup)
}
func (a *Ajax) PrintSelectedRows(url string) c.IButton {
	return newPrintSelectedRowsButton(a.Name, url)
}
func (a *Ajax) ToggleSelectedRows(url string) c.IButton {
	return newToggleSelectedRowsButton(a.Name, url)
}

// AJAX response only !!! Used only for update etc..and OK or Error message on response.
// For normal HTML response use Html.Submit()
func (a *Ajax) Submit(formName string, url string) c.IButton {
	return newSubmitButton(a.Name, formName, url)
}
