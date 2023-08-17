package transactions

import (
	m "quince/internal/models"
	"quince/modules/{{.Model.Module}}/models"
	"github.com/beego/beego/v2/client/orm"
)
type {{.Model.Name_lcase}}Update struct {
	*models.{{.Model.Name}} //unnamed parameter only !!!
}

func {{.Model.Name}}Update(c *models.{{.Model.Name}}) {{.Model.Name_lcase}}Update {
	t := {{.Model.Name_lcase}}Update{c}
	return t
}
func (u {{.Model.Name_lcase}}Update) Run(txOrm orm.TxOrmer) error {
	err := u.Validate()
	if err != nil {
		return err
	} else {
		return u.Update(txOrm, u.{{.Model.Name}})
	}
}
func (u {{.Model.Name_lcase}}Update) Description() string {
	return "{{.Model.Name}} update"
}
func (u {{.Model.Name_lcase}}Update) GetModel() m.IModel {
	return &u
}


