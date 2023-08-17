package modal

import (
	c "quince/internal/toolbar/component"
)

type Modal struct {
	Name string
}

func (a *Modal) View(url string) c.IButton {
	return newViewButton(a.Name, url)
}
func (a *Modal) Edit( url string) c.IButton {
	return newEditButton(a.Name, url)
}
func (a *Modal) Add(url string) c.IButton {
	return newAddButton(a.Name, url)
}