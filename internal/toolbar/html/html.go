package html

import (
	c "quince/internal/toolbar/component"
)

type Html struct {
	Name string
}

func (a *Html) Add(url string) c.IButton {
	return newAddButton(a.Name, url)
}
func (a *Html) View(url string) c.IButton {
	return newViewButton(a.Name, url)
}
func (a *Html) Edit(url string) c.IButton {
	return newEditButton(a.Name, url)
}
func (a *Html) Submit(formName string,url string) c.IButton {
	return newSubmitButton(a.Name,formName, url)
}
func (a *Html) New(title string, icon string, url string) c.IButton {
	return newButton(a.Name, title, icon, url)
}
