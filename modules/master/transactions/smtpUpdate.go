package transactions

import (
	"quince/global"
	m "quince/internal/models"
	"quince/modules/master/models"
	"quince/utils/encrypter"

	"github.com/beego/beego/v2/client/orm"
)

type smtpUpdate struct {
	*models.SMTP //unnamed parameter only !!!
}

func SmtpUpdate(c *models.SMTP) smtpUpdate {
	t := smtpUpdate{c}
	return t
}
func (u smtpUpdate) Run(txOrm orm.TxOrmer) error {
	cryptData := encrypter.Encrypt([]byte(u.Password), []byte(global.BA_CONFIG.Other.LogAesKey))
	u.Password = cryptData
	return u.Update(txOrm, u.SMTP)
}
func (u smtpUpdate) Description() string {
	return "SMTP update"
}
func (u smtpUpdate) GetModel() m.IModel {
	return &u
}
