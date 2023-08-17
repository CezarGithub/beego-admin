package form

import (
	"quince/internal/copier"
	"quince/modules/admin/models"
)

// UserForm user
type UserForm struct {
	Id          int    `form:"id"`
	Avatar      string `form:"avatar"`
	Username    string `form:"username"`
	Nickname    string `form:"nickname"`
	Mobile      string `form:"mobile"`
	Email       string `form:"email"`
	UserLevelId int64  `form:"user_level_id"`
	Status      int8   `form:"status"`
	Description string `form:"description"`
	IsCreate    int    `form:"_create"`
}

func (c *UserForm) Validate() (*models.User, error) {
	var m models.User
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		m.UserLevel = &models.UserLevel{}
		m.UserLevel.Id = c.UserLevelId
		return &m, m.Validate()
	}
}
