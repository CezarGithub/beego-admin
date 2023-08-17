package form

import (
	"quince/internal/copier"
	"quince/modules/admin/models"
)

// AdminRoleForm
type AdminRoleForm struct {
	Id          int    `form:"id"`
	Name        string `form:"name" `
	Description string `form:"description" `
	Status      int8   `form:"status" `
	IsCreate    int    `form:"_create"`
}

func (c *AdminRoleForm) Validate() (*models.AdminRole, error) {
	var m models.AdminRole
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		return &m, m.Validate()
	}
}
