package models

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"quince/initialize/database"
	"quince/internal/validation"
	"strings"

	"github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

// LoginUser struct
type LoginUser struct {
	User      *User  `orm:"rel(fk)"`
	LoginName string `orm:"column(login_name);unique;size(30)" description:"Login name" json:"login_name" i18n:"admin.user.login_name"`
	Password  string `orm:"column(password);size(255)" description:"Password" json:"password"  i18n:"admin.user.password"`
	Roles     string `orm:"column(role);size(200)" description:"Role" json:"role"   i18n:"admin.user.roles"`
	Status    int8   `orm:"column(status);size(1)" description:"Enabled 0：yes 1：no" json:"status" i18n:"admin.user.status"`
	Language  string `orm:"column(language);size(10)" description:"Language" json:"language" i18n:"admin.user.lang"`
	Base
}

func NewLoginUser() LoginUser {
	user := NewUser()
	adminUser := LoginUser{Base: Base{Id: 0}, User: &user}
	return adminUser
}

// TableName
func (*LoginUser) TableName() string {
	return "admin_login"
}

// SearchField
func (*LoginUser) SearchField() []string {
	return []string{"LoginName"}
}

func (cs *LoginUser) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.LoginName != "" {
			cond = cond.And("login_name__icontains", cs.LoginName)
		}
		if cs.User != nil {
			if cs.User.Mobile != "" {
				cond = cond.And("mobile__icontains", cs.User.Mobile)
			}
			if cs.User.Email != "" {
				cond = cond.And("email__icontains", cs.User.Email)
			}
			if cs.User.UserLevel != nil {
				if cs.User.UserLevel.Id != 0 {
					cond = cond.And("user_level_id", cs.User.UserLevel.Id)
				}
			}
		}
	}
	return cond
}

// TimeField
func (*LoginUser) TimeField() []string {
	return []string{}
}

func (c *LoginUser) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.User.Id, validation.Required, validation.Min(1)))
	rules = append(rules, validation.Field(&c.Password, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.LoginName, validation.Required, validation.Length(1, 25)))
	rules = append(rules, validation.Field(&c.Roles, validation.Required, validation.Length(1, 300)))
	rules = append(rules, validation.Field(&c.Language, validation.Required, validation.Length(1, 5)))
	err := validation.ValidateStruct(c, rules...)
	return err
}

func init() {
	database.RegisterModel("admin", new(LoginUser))
}

// GetSignStrByAdminUser
func (adminUser *LoginUser) GetSignStrByAdminUser(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d%s%s", adminUser.Id, adminUser.User.Username, ua))))
}

// GetAuthUrl
func (adminUser *LoginUser) GetMenuUrl() map[string]interface{} {
	var (
		urlArr orm.ParamsList
	)
	authURL := make(map[string]interface{})

	o := orm.NewOrm()
	qs := o.QueryTable(new(AdminRole))

	_, err := qs.Filter("id__in", strings.Split(adminUser.Roles, ",")).Filter("status", 1).ValuesFlat(&urlArr, "url")
	if err == nil {
		urlIDStr := ""
		for k, row := range urlArr {
			urlStr, ok := row.(string)
			if ok {
				if k == 0 {
					urlIDStr = urlStr
				} else {
					urlIDStr += "," + urlStr
				}
			}
		}
		urlIDArr := strings.Split(urlIDStr, ",")

		var authURLArr orm.ParamsList

		if len(urlIDStr) > 0 {
			o = orm.NewOrm()
			qs = o.QueryTable(new(AdminMenu))
			_, err := qs.Filter("id__in", urlIDArr).ValuesFlat(&authURLArr, "url")
			if err == nil {
				for k, row := range authURLArr {
					val, ok := row.(string)
					if ok {
						authURL[val] = k
					}
				}
			}
		}
		return authURL
	}
	return authURL
}

// GetShowMenu
// func (adminUser *LoginUser) GetShowMenu() map[int64]orm.Params {
// 	var maps []orm.Params
// 	returnMaps := make(map[int64]orm.Params)
// 	o := orm.NewOrm()

// 	if adminUser.Id == 1 {
// 		_, err := o.QueryTable(new(AdminMenu)).Filter("is_show", 1).OrderBy("parent_id", "sort_id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
// 		if err == nil {
// 			for _, m := range maps {
// 				returnMaps[m["Id"].(int64)] = m
// 			}
// 			return returnMaps
// 		}
// 		return map[int64]orm.Params{}
// 	}

// 	var list orm.ParamsList
// 	_, err := o.QueryTable(new(AdminRole)).Filter("id__in", strings.Split(adminUser.Roles, ",")).Filter("status", 1).ValuesFlat(&list, "url")
// 	if err == nil {
// 		var urlIDArr []string
// 		for _, m := range list {
// 			urlIDArr = append(urlIDArr, strings.Split(m.(string), ",")...)
// 		}
// 		_, err := o.QueryTable(new(AdminMenu)).Filter("id__in", urlIDArr).Filter("is_show", 1).OrderBy("sort_id", "id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
// 		if err == nil {
// 			for _, m := range maps {
// 				returnMaps[m["Id"].(int64)] = m
// 			}
// 			return returnMaps
// 		}
// 		return map[int64]orm.Params{}
// 	}
// 	return map[int64]orm.Params{}

// }

// GetRoleText
func (adminUser *LoginUser) GetRoleText() map[int64]*AdminRole {
	roleIDArr := strings.Split(adminUser.Roles, ",")
	var adminRole []*AdminRole
	_, err := orm.NewOrm().QueryTable(new(AdminRole)).Filter("id__in", roleIDArr, "id", "name").All(&adminRole)
	if err != nil {
		return nil
	}
	adminRoleMap := make(map[int64]*AdminRole)
	for _, v := range adminRole {
		adminRoleMap[v.Id] = v
	}
	return adminRoleMap
}

// GetLoginUser
func (*LoginUser) GetLoginUser() []*LoginUser {
	var loginUsers []*LoginUser
	_, err := orm.NewOrm().QueryTable(new(LoginUser)).All(&loginUsers)
	if err == nil {
		return loginUsers
	}
	return nil
}

func (t *LoginUser) Export() []byte {
	var items []*LoginUser
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *LoginUser) Import(tx orm.TxOrmer, data []byte) error {
	var list []*LoginUser
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		return err
	}
	//tx.QueryTable(t.TableName()).Filter("id__gt", 1).Delete() //do not delete super_admin
	admin, _ := web.AppConfig.String("adminname")
	//tx.QueryTable(t.TableName()).Exclude("login_name", admin).Delete() //do not delete super_admin
	//tx.Raw("DELETE FROM admin_user WHERE login_name !=?",admin).Exec()
	// for _, item := range list {
	// 	if item.LoginName != admin { //skip super_admin
	// 		_, err := tx.Insert(item)
	// 		if err != nil {
	// 			return err

	// 		}
	// 	}
	// }
	logs.Info(admin)
	return nil
}

// GetShowMenu
func (adminUser *LoginUser) GetLeftMenu() ([]orm.Params, error) {
	var maps []orm.Params
	o := orm.NewOrm()

	if adminUser.Id == 1 {
		_, err := o.QueryTable(new(AdminMenu)).Filter("is_show", 1).OrderBy("parent_id", "sort_id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
		return maps, err
	}

	var list orm.ParamsList
	_, err := o.QueryTable(new(AdminRole)).Filter("id__in", strings.Split(adminUser.Roles, ",")).Filter("status", 1).ValuesFlat(&list, "url")
	if err == nil {
		var urlIDArr []string
		for _, m := range list {
			urlIDArr = append(urlIDArr, strings.Split(m.(string), ",")...)
		}
		_, err := o.QueryTable(new(AdminMenu)).Filter("id__in", urlIDArr).Filter("is_show", 1).OrderBy("parent_id", "sort_id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
		return maps, err
	}
	return nil, err

}
