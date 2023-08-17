package services

import (
	"encoding/json"
	"quince/internal/i18n"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// AdminMenuService struct
type adminMenuService struct {
	BaseService
}

// AdminLogService - instantiate de IModel filter
func NewAdminMenuService() adminMenuService {
	var cs adminMenuService
	c := models.AdminMenu{}

	cs.IModel = &c
	return cs
}

// GetAdminMenuByUrl
func (cs *adminMenuService) GetAdminMenuByUrl(url string) *models.AdminMenu {
	var adminMenu models.AdminMenu
	err := cs.DataQuery().QueryTable(new(models.AdminMenu)).Filter("url", url).One(&adminMenu)
	if err == nil {
		return &adminMenu
	}
	return nil
}

// GetCount
func (cs *adminMenuService) GetCount() int {
	count, err := cs.DataQuery().QueryTable(new(models.AdminMenu)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetCount per module
func (cs *adminMenuService) GetCountModule(moduleName string) int {
	count, err := cs.DataQuery().QueryTable(new(models.AdminMenu)).Filter("module", moduleName).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// AllMenu
func (cs *adminMenuService) AllMenu() []*models.AdminMenu {
	var adminMenus []*models.AdminMenu
	_, err := cs.DataQuery().QueryTable(new(models.AdminMenu)).OrderBy("sort_id", "id").All(&adminMenus)
	if err == nil {
		return adminMenus
	}
	return nil
}

// Menu
func (cs *adminMenuService) Menu(currentID int64) []orm.Params {
	var adminMenusMap []orm.Params
	cs.DataQuery().QueryTable(new(models.AdminMenu)).Exclude("id", currentID).OrderBy("sort_id", "id").Values(&adminMenusMap, "id", "parent_id", "name", "sort_id")
	return adminMenusMap
}

// IsExistUrl
func (cs *adminMenuService) IsExistUrl(url string, id int64) bool {
	if id == 0 {
		return cs.DataQuery().QueryTable(new(models.AdminMenu)).Filter("url", url).Exist()
	}
	return cs.DataQuery().QueryTable(new(models.AdminMenu)).Filter("url", url).Exclude("id", id).Exist()
}

// IsChildMenu
func (cs *adminMenuService) IsChildMenu(ids []int64) bool {
	return cs.DataQuery().QueryTable(new(models.AdminMenu)).Filter("parent_id__in", ids).Exist()
}

// GetAdminMenuById
func (cs *adminMenuService) GetAdminMenuById(id int64) *models.AdminMenu {
	var adminMenu models.AdminMenu
	err := cs.DataQuery().QueryTable(new(models.AdminMenu)).Filter("id", id).One(&adminMenu)
	if err == nil {
		return &adminMenu
	}
	return nil
}

func (cs *adminMenuService) GetLeftMenu(requestUrl string, user models.LoginUser, language string) string {
	eson := ` [
		{
			"Name": "Server error",
			"Id": "1",
			"ParentId": "-1",
			"Url": "#",
			"Icon": "fa-exclamation-triangle "
	
		}]`
	menuMap, err := user.GetLeftMenu()
	if err != nil {
		logs.Error(err.Error())
		return eson
	}
	//var menu []interface{}
	for k, v := range menuMap {
		//v["Name"] = i18n.Tr(language, string(v["Name"].(string)))
		//menu = append(menu, v)
		menuMap[k]["Name"] = i18n.Tr(language, string(v["Name"].(string)))
	}

	json, err := json.Marshal(menuMap)
	if err != nil {
		logs.Error(err.Error())
		return eson
	}
	return string(json)
}
