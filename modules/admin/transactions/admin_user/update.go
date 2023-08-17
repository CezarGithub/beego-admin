package admin_user

import (
	"encoding/base64"
	"errors"
	m "quince/internal/models"
	"quince/modules/admin/models"
	"quince/modules/admin/services"
	"quince/utils"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type update struct {
	*models.LoginUser //unnamed parameter only !!!
}

func LoginUserUpdate(c *models.LoginUser) update {
	t := update{c}
	return t
}
func (u update) Run(txOrm orm.TxOrmer) error {
	//Login name check
	adminUserService := services.NewAdminUserService()
	if adminUserService.IsExistName(strings.TrimSpace(u.LoginName), u.GetID()) {
		return errors.New("error.already_exists")
	}

	//check password changed
	adminUser := models.LoginUser{Base: models.Base{Id: u.Id}}
	o := orm.NewOrm()
	if o.Read(&adminUser) == nil || u.Id == 0 {
		if adminUser.Password != u.Password {
			newPasswordForHash, err := utils.PasswordHash(u.Password)
			if err == nil {
				u.Password = base64.StdEncoding.EncodeToString([]byte(newPasswordForHash))
			} else {
				return err
			}
		}
	}
	// if(strings.Trim(u..Avatar)==""){
	// 	u.Avatar = "/static/admin/images/avatar.png"
	// }
	return u.Update(txOrm, u.LoginUser)
}

func (u update) Description() string {
	return "Login user update"
}
func (u update) GetModel() m.IModel {
	return &u
}
