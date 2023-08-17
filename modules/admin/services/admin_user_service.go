package services

import (
	"encoding/base64"
	"errors"
	"net/url"
	"quince/global"
	"quince/modules/admin/form"
	"quince/modules/admin/models"
	"quince/utils"
	"quince/utils/page"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"


	"github.com/beego/beego/v2/server/web/context"
)

// AdminUserService struct
type adminUserService struct {
	BaseService
}

// NewAdminUserService - instantiate de IModel filter
func NewAdminUserService() adminUserService {
	var cs adminUserService
	c := models.LoginUser{}
	cs.IModel = &c
	return cs
}

// GetAdminUserById
func (cs *adminUserService) GetAdminUserById(id int64) *models.LoginUser {

	adminUser := models.LoginUser{Base: models.Base{Id: id}}
	err := cs.DataQuery().Read(&adminUser)
	cs.DataQuery().LoadRelated(&adminUser, "User")
	if err != nil {
		return nil
	}

	return &adminUser
}

// MenuCheck
func (*adminUserService) MenuCheck(url string, authExcept map[string]interface{}, loginUser *models.LoginUser) bool {
	if loginUser.Id == 1 { //is SUPER_ADMIN - has full rights
		return true
	}
	authURL := loginUser.GetMenuUrl()
	if utils.KeyInMap(url, authExcept) || utils.KeyInMap(url, authURL) {
		return true
	}
	return false
}
func (cs *adminUserService) UrlCheck(url string, authExcept map[string]interface{}, loginUser *models.LoginUser) bool {
	var user_rbac []models.AdminRbac
	check := false
	if loginUser.Id == 1 { //is SUPER_ADMIN - has full rights
		return true
	}
	user_roles := strings.Split(loginUser.Roles, ",")
	sql := "SELECT b.role_id FROM admin_route a LEFT JOIN admin_rbac b ON a.id=b.route_id WHERE a.url=?"
	_, err := cs.DataManipulation().Raw(sql, url).QueryRows(&user_rbac)
	if err != nil {
		logs.Error(err.Error())
	}
	for _, r := range user_roles {
		role, _ := strconv.ParseInt(r, 10, 64)
		for _, rbac := range user_rbac {
			if role == rbac.Role.Id {
				check = true
				break
			}
		}
	}
	return check
}
func (cs *adminUserService) APIUrlCheck(url string, userID int64, roles string) bool {
	var user_rbac []models.AdminRbac
	check := false
	if userID == 1 { //is SUPER_ADMIN - has full rights
		return true
	}
	user_roles := strings.Split(roles, ",")
	sql := "SELECT b.role_id FROM admin_route a LEFT JOIN admin_rbac b ON a.id=b.route_id WHERE a.url=? AND a.isAPI=1 " //must be API url
	_, err := cs.DataManipulation().Raw(sql, url).QueryRows(&user_rbac)
	if err != nil {
		logs.Error(err.Error())
	}
	for _, r := range user_roles {
		role, _ := strconv.ParseInt(r, 10, 64)
		for _, rbac := range user_rbac {
			if role == rbac.Role.Id {
				check = true
				break
			}
		}
	}
	return check
}

// CheckLogin
func (cs *adminUserService) CheckLogin(loginForm form.LoginForm, ctx *context.Context) (*models.LoginUser, error) {
	var adminUser models.LoginUser
	err := cs.DataQuery().QueryTable(new(models.LoginUser)).RelatedSel().Filter("login_name", loginForm.Username).Limit(1).One(&adminUser)
	if err != nil {
		return nil, errors.New("login.no_user")
	}

	decodePasswdStr, err := base64.StdEncoding.DecodeString(adminUser.Password)

	if err != nil || !utils.PasswordVerify(loginForm.Password, string(decodePasswdStr)) {
		return nil, errors.New("login.password_wrong")
	}

	if adminUser.Status != 1 {
		return nil, errors.New("login.user_disabled")
	}

	ctx.Output.Session(global.LOGIN_USER, adminUser)

	if loginForm.Remember != "" {
		ctx.SetCookie(global.LOGIN_USER_ID, (string)(adminUser.Id), 7200)
		ctx.SetCookie(global.LOGIN_USER_ID_SIGN, adminUser.GetSignStrByAdminUser(ctx), 7200)
	} else {
		ctx.SetCookie(global.LOGIN_USER_ID, ctx.GetCookie(global.LOGIN_USER_ID), -1)
		ctx.SetCookie(global.LOGIN_USER_ID_SIGN, ctx.GetCookie(global.LOGIN_USER_ID_SIGN), -1)
	}

	return &adminUser, nil

}

// GetCount
func (cs *adminUserService) GetCount() int {
	count, err := cs.DataQuery().QueryTable(new(models.LoginUser)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetAllAdminUser
func (cs *adminUserService) GetAllAdminUser() []*models.LoginUser {
	var adminUser []*models.LoginUser
	o := cs.DataQuery().QueryTable(new(models.LoginUser))
	_, err := o.All(&adminUser)
	if err != nil {
		return nil
	}
	return adminUser
}

// GetPaginateData
func (aus *adminUserService) GetPaginateData(listRows int, params url.Values) ([]*models.LoginUser, page.Pagination) {
	//
	aus.SearchField = append(aus.SearchField, new(models.LoginUser).SearchField()...)

	var adminUser []*models.LoginUser
	o := aus.DataQuery().QueryTable(new(models.LoginUser)).RelatedSel()
	_, err := aus.PaginateAndScopeWhere(o, listRows, params).All(&adminUser)
	if err != nil {
		return nil, aus.Pagination
	}
	return adminUser, aus.Pagination
}

// IsExistName
func (cs *adminUserService) IsExistName(username string, id int64) bool {
	if id == 0 {
		return cs.DataQuery().QueryTable(new(models.LoginUser)).Filter("login_name", username).Exist()
	}
	return cs.DataQuery().QueryTable(new(models.LoginUser)).Filter("login_name", username).Exclude("id", id).Exist()
}
