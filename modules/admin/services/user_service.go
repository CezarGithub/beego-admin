package services

import (
	"net/url"
	"quince/modules/admin/models"
	"quince/utils/page"
)

// UserService struct
type userService struct {
	BaseService
}

func NewUserService() userService {
	var cs userService
	c := models.User{}

	cs.IModel = &c
	return cs
}

// GetPaginateData
func (us *userService) GetPaginateData(listRows int, params url.Values) ([]*models.User, page.Pagination) {
	//
	us.SearchField = append(us.SearchField, new(models.User).SearchField()...)

	var users []*models.User
	o := us.DataQuery().QueryTable(new(models.User))
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&users)
	if err != nil {
		return nil, us.Pagination
	}
	return users, us.Pagination
}

// GetUserById
func (us *userService) GetUserById(id int64) *models.User {
	user := models.User{}
	user.Id = id
	err := us.DataQuery().Read(&user)
	if err != nil {
		return nil
	}
	return &user
}

// GetExportData
func (us *userService) GetExportData(params url.Values) []*models.User {
	//
	us.SearchField = append(us.SearchField, new(models.User).SearchField()...)
	var user []*models.User
	o := us.DataQuery().QueryTable(new(models.User))
	_, err := us.ScopeWhere(o, params).All(&user)
	if err != nil {
		return nil
	}
	return user
}

// GetAll
func (us *userService) GetAll(params url.Values) []*models.User {

	var list []*models.User
	o := us.DataQuery().QueryTable(new(models.User)).OrderBy("name")
	_, err := us.GetAllAndScopeWhere(o.OrderBy("name"), params).All(&list)
	if err != nil {
		return nil
	}
	return list
}
