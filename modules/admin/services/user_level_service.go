package services

import (
	"net/url"
	"quince/modules/admin/models"
	"quince/utils/page"
)

// UserLevelService struct
type userLevelService struct {
	BaseService
}

func NewUserLevelService() userLevelService {
	var cs userLevelService
	c := models.UserLevel{}

	cs.IModel = &c
	return cs
}

// GetPaginateData
func (uls *userLevelService) GetPaginateData(listRows int, params url.Values) ([]*models.UserLevel, page.Pagination) {

	uls.SearchField = append(uls.SearchField, new(models.UserLevel).SearchField()...)

	var userLevel []*models.UserLevel
	o := uls.DataQuery().QueryTable(new(models.UserLevel))
	_, err := uls.PaginateAndScopeWhere(o, listRows, params).All(&userLevel)
	if err != nil {
		return nil, uls.Pagination
	}
	return userLevel, uls.Pagination
}

// IsExistName
func (uls *userLevelService) IsExistName(name string, id int64) bool {
	if id == 0 {
		return uls.DataQuery().QueryTable(new(models.UserLevel)).Filter("name", name).Exist()
	}
	return uls.DataQuery().QueryTable(new(models.UserLevel)).Filter("name", name).Exclude("id", id).Exist()
}

// GetExportData
func (uls *userLevelService) GetExportData(params url.Values) []*models.UserLevel {
	//
	uls.SearchField = append(uls.SearchField, new(models.UserLevel).SearchField()...)
	var userLevel []*models.UserLevel
	o := uls.DataQuery().QueryTable(new(models.UserLevel))
	_, err := uls.ScopeWhere(o, params).All(&userLevel)
	if err != nil {
		return nil
	}
	return userLevel
}

// Create
// func (*userLevelService) Create(form *form.UserLevelForm) int {
// 	userLevel := models.UserLevel{
// 		Name:        form.Name,
// 		Description: form.Description,
// 		Status:      int8(form.Status),
// 	}
// 	if form.Img != "" {
// 		userLevel.Img = form.Img
// 	}
// 	id, err := orm.NewOrm().Insert(&userLevel)

// 	if err == nil {
// 		return int(id)
// 	}
// 	return 0
// }

// Update
// func (*userLevelService) Update(form *form.UserLevelForm) int {
// 	o := orm.NewOrm()
// 	userLevel := models.UserLevel{Base: models.Base{Id: form.Id}}
// 	if o.Read(&userLevel) == nil {
// 		userLevel.Name = form.Name
// 		userLevel.Description = form.Description
// 		userLevel.Status = int8(form.Status)
// 		userLevel.UpdateTime = time.Now()
// 		if form.Img != "" {
// 			userLevel.Img = form.Img
// 		}
// 		userLevel.Name = form.Name
// 		num, err := o.Update(&userLevel)
// 		if err == nil {
// 			return int(num)
// 		}
// 		return 0
// 	}
// 	return 0
// }

// GetUserLevelById
func (usl *userLevelService) GetUserLevelById(id int64) *models.UserLevel {
	userLevel := models.UserLevel{Base: models.Base{Id: id}}
	err := usl.DataQuery().Read(&userLevel)
	if err != nil {
		return nil
	}
	return &userLevel
}

// GetUserLevel
func (uls *userLevelService) GetUserLevel() []*models.UserLevel {
	var userLevels []*models.UserLevel
	_, err := uls.DataQuery().QueryTable(new(models.UserLevel)).All(&userLevels)
	if err == nil {
		return userLevels
	}
	return nil
}

// Enable
// func (*userLevelService) Enable(ids []int) int {
// 	num, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Update(orm.Params{
// 		"status": 1,
// 	})
// 	if err == nil {
// 		return int(num)
// 	}
// 	return 0
// }

// // Disable
// func (*userLevelService) Disable(ids []int) int {
// 	num, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Update(orm.Params{
// 		"status": 0,
// 	})
// 	if err == nil {
// 		return int(num)
// 	}
// 	return 0
// }

// Del
// func (*userLevelService) Del(ids []int) int {
// 	count, err := orm.NewOrm().QueryTable(new(models.UserLevel)).Filter("id__in", ids).Delete()
// 	if err == nil {
// 		return int(count)
// 	}
// 	return 0
// }
