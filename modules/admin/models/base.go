package models

import (
	"quince/internal/models"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Base struct {
	Id         int64     `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"  i18n:"app.id" form:"id"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);column(create_time)" description:"Creation date" json:"create_time"   i18n:"app.create_time"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);column(update_time)" description:"Update date" json:"update_time"   i18n:"app.update_time"`
	CreateUser int64     `orm:"column(create_user);default(0)" description:"Created by user" json:"create_user"   i18n:"app.create_user"`
	UpdateUser int64     `orm:"column(update_user);default:(0)" description:"Modified by user" json:"update_user"   i18n:"app.update_user"`
}

func (b *Base) GetID() int64 {
	return b.Id
}
func (b *Base) SetID(id int64) {
	b.Id = id
}
func (b *Base) SetUser(userID int64) {
	if b.GetID() == 0 {
		b.CreateUser = userID
		b.CreateTime = time.Now()
	} else {
		b.UpdateUser = userID
		b.UpdateTime = time.Now()
	}
}

// Update method can be overrited in model
func (b *Base) Update(tx orm.TxOrmer, c models.IModel) error {
	var e error
	if b.GetID() == 0 {
		_, e = tx.Insert(c)
	} else {
		_, e = tx.Update(c)
	}
	return e
}

// Spare function.Override in model if neccesary
func (t *Base) InitData(tx orm.TxOrmer) error {
	return nil
}
