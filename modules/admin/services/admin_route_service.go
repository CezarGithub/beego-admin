package services

import (
	"net/url"
	"quince/modules/admin/models"
	"quince/utils/page"

	"github.com/beego/beego/v2/core/logs"
)

// AdminRouteService struct
type adminRouteService struct {
	BaseService
}

// AdminLogService - instantiate de IModel filter
func NewAdminRouteService() adminRouteService {
	var cs adminRouteService
	c := models.AdminRoute{}

	cs.IModel = &c
	return cs
}

// GetAdminMenuByUrl
func (cs *adminRouteService) GetAdminRouteByUrl(url string) *models.AdminRoute {
	var adminRoute models.AdminRoute
	err := cs.DataQuery().QueryTable(new(models.AdminRoute)).Filter("url", url).One(&adminRoute)
	if err == nil {
		return &adminRoute
	}
	return nil
}

// GetAdminMenuByUrl
func (cs *adminRouteService) GetAdminRouteById(id int64) *models.AdminRoute {
	var adminRoute models.AdminRoute
	err := cs.DataQuery().QueryTable(new(models.AdminRoute)).Filter("id", id).One(&adminRoute)
	if err == nil {
		return &adminRoute
	}
	return nil
}

// GetCount
func (cs *adminRouteService) GetCount() int64 {
	count, err := cs.DataQuery().QueryTable(new(models.AdminRoute)).Count()
	if err != nil {
		return 0
	}
	return count
}

// GetCount per module
func (cs *adminRouteService) GetCountModule(moduleName string) int64 {
	count, err := cs.DataQuery().QueryTable(new(models.AdminRoute)).Filter("module", moduleName).Count()
	if err != nil {
		return 0
	}
	return count
}

// AllRoute
func (cs *adminRouteService) AllRoute() []*models.AdminRoute {
	var adminRoutes []*models.AdminRoute
	_, err := cs.DataQuery().QueryTable(new(models.AdminRoute)).OrderBy("module", "controller", "url").All(&adminRoutes)
	if err == nil {
		return adminRoutes
	}
	return nil
}

// AllModules - returns all distinct modules names
func (cs *adminRouteService) AllModules() []*models.AdminRoute {
	var adminRoutes []*models.AdminRoute
	_, err := cs.DataQuery().QueryTable(new(models.AdminRoute)).Distinct().All(&adminRoutes, "module")
	if err == nil {
		return adminRoutes
	}
	return nil
}

// AllControllers  returns all distinct controllers names
func (cs *adminRouteService) AllControllers() []*models.AdminRoute {
	var adminRoutes []*models.AdminRoute
	_, err := cs.DataQuery().QueryTable(new(models.AdminRoute)).Distinct().All(&adminRoutes, "module", "controller")
	if err == nil {
		return adminRoutes
	}
	return nil
}

// IsExistUrl
func (cs *adminRouteService) IsExistUrl(url string, id int) bool {
	if id == 0 {
		return cs.DataQuery().QueryTable(new(models.AdminRoute)).Filter("url", url).Exist()
	}
	return cs.DataQuery().QueryTable(new(models.AdminRoute)).Filter("url", url).Exclude("id", id).Exist()
}

// GetAdminRouteById
func (ars *adminRouteService) GetRoleRoutesById(roleId int64, listRows int, params url.Values) ([]*models.AdminRoute, page.Pagination) {
	var adminRbac []models.AdminRbac
	adminRoute, pagination := ars.GetPaginateData(listRows, params)
	_, err := ars.DataQuery().QueryTable(new(models.AdminRbac)).Filter("role_id", roleId).RelatedSel().All(&adminRbac)
	if err != nil {
		logs.Error(err)
		return nil, pagination
	}

	for i, route := range adminRoute {
		adminRoute[i].Status = 0
		for _, rbac := range adminRbac {
			if rbac.Route.Id == route.Id {
				adminRoute[i].Status = 1 //use status field for selct/show role's routes
				break
			}
		}
	}
	return adminRoute, pagination
}

// GetPaginateData adminrole
func (ars *adminRouteService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminRoute, page.Pagination) {

	ars.SearchField = append(ars.SearchField, new(models.AdminRoute).SearchField()...)

	var adminRoute []*models.AdminRoute
	o := ars.DataQuery().QueryTable(new(models.AdminRoute))
	_, err := ars.PaginateAndScopeWhere(o, listRows, params).All(&adminRoute)
	if err != nil {
		return nil, ars.Pagination
	}
	return adminRoute, ars.Pagination
}
