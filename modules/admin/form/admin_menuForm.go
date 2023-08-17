package form

import (
	"quince/internal/copier"
	"quince/modules/admin/models"
	"strings"
)

// AdminMenuForm
type AdminMenuForm struct {
	Id        int    `form:"id"`
	ParentId  int    `form:"parent_id" `
	Name      string `form:"name" `
	I18n      string `form:"i18n" `
	Url       string `form:"url" `
	Icon      string `form:"icon" `
	IsShow    int8   `form:"is_show"`
	SortId    int    `form:"sort_id"`
	LogMethod string `form:"log_method" `
	IsCreate  int    `form:"_create"`
}

func (c *AdminMenuForm) Validate() (*models.AdminMenu, error) {
	var m models.AdminMenu
	c.Url = strings.TrimSpace(c.Url) //remove spaces
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		return &m, m.Validate()
	}
}
