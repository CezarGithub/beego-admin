package models

import (
	"encoding/json"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/internal/validation/is"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// User struct
type User struct {
	Avatar      string     `orm:"column(avatar);size(255)" description:"avatar" json:"avatar" i18n:"admin.user.avatar"`
	Username    string     `orm:"column(username);size(30)" description:"User name" json:"username" i18n:"admin.user.username"`
	Nickname    string     `orm:"column(nickname);size(30)" description:"Nickname" json:"nickname" i18n:"admin.user.nickname"`
	Mobile      string     `orm:"column(mobile);size(11)" description:"Mobile phone" json:"mobile" i18n:"admin.user.mobile"`
	Email       string     `orm:"column(email);size(255)" description:"Mobile phone" json:"email" i18n:"admin.user.email"`
	Status      int8       `orm:"column(status);size(1);default(1)" description:"Status 0：disabled 1：enabled" json:"status" i18n:"admin.user.status"`
	Description string     `orm:"column(description);type(text)" description:"Comments" json:"description" i18n:"admin.user.description"`
	UserLevel   *UserLevel `orm:"rel(fk)"`
	Base
}

func NewUser() User {
	user := User{Base: Base{Id: 0}}
	return user
}

// TableName
func (*User) TableName() string {
	return "admin_user"
}

// SearchField
func (*User) SearchField() []string {
	return []string{"username", "mobile", "nickname"}
}

func (cs *User) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Username != "" {
			cond = cond.And("name__icontains", cs.Username)
		}
		if cs.Mobile != "" {
			cond = cond.And("mobile__icontains", cs.Mobile)
		}
		if cs.Email != "" {
			cond = cond.And("email__icontains", cs.Email)
		}
		if cs.UserLevel != nil {
			if cs.UserLevel.Id != 0 {
				cond = cond.And("user_level_id", cs.UserLevel.Id)
			}
		}
	}
	return cond
}

// TimeField
func (*User) TimeField() []string {
	return []string{}
}

// init model
func init() {
	database.RegisterModel("admin", new(User))
}
func (c *User) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Username, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Nickname, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Mobile, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Email, validation.Required, is.Email))
	rules = append(rules, validation.Field(&c.UserLevel.Id, validation.Required, validation.Min(0)))
	err := validation.ValidateStruct(c, rules...)
	return err
}

func (t *User) Export() []byte {
	var items []*User
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *User) Import(tx orm.TxOrmer, data []byte) error {
	var list []*User
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		return err
	}
	admin, _ := web.AppConfig.String("adminname")
	tx.QueryTable(t.TableName()).Exclude("username", admin).Delete() //do not delete super_admin
	for _, item := range list {
		if item.Username != admin {
			_, err := tx.Insert(item)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
