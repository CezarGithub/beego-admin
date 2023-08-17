package form

import (
	"quince/internal/copier"
	"quince/modules/master/models"
)

// GroupForm
type GroupForm struct {
	Id          int64  `form:"id"`
	Name        string `form:"name"`
	Description string `form:"description"`
	Status      int8   `form:"status"`
}

func (c *GroupForm) Validate() (*models.Group, error) {
	var m models.Group
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		return &m, m.Validate()
	}
}
