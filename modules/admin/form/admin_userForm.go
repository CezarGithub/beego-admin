package form

import (
	"quince/internal/copier"
	"quince/modules/admin/models"
)

// AdminUserForm admin_user
type AdminUserForm struct {
	Id        int64   `form:"id"`
	UserID    int64    `form:"user_id"`
	Username  string `form:"username" `
	LoginName string `form:"loginname" `
	Password  string `form:"password" `
	Nickname  string `form:"nickname" `
	Avatar    string `form:"avatar"`
	Roles     string `form:"role" `
	Status    int    `form:"status"`
	Language  string `form:"language" `
	IsCreate  int    `form:"_create"`
}

func (c *AdminUserForm) Validate() (*models.LoginUser, error) {
	var m models.LoginUser
	if err := copier.Copy(&m, c); err != nil {
		return &m, err
	} else {
		m.User = &models.User{}
		m.User.Id = c.UserID
		return &m, m.Validate()
	}
}
