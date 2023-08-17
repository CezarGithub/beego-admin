package form

import (
	"quince/internal/copier"
	"quince/modules/admin/models"
)

// UserLevelForm admin_user
type UserLevelForm struct {
	Id          int    `form:"id"`
	Name        string `form:"name"`
	Description string `form:"description"`
	Img         string `form:"img"`
	Status      int8   `form:"status"`
	CreateTime  int    `form:"create_time"`
	UpdateTime  int    `form:"update_time"`
	DeleteTime  int    `form:"delete_time"`
	IsCreate    int    `form:"_create"`
}

func (c *UserLevelForm) Validate() (*models.UserLevel, error) {
	var m models.UserLevel
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		return &m, m.Validate()
	}
}
