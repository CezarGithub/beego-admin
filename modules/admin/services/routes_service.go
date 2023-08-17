package services

import (
	"net/url"
	"quince/utils/page"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type RegisteredRoute struct {
	Method     string
	Url        string
	Controller string
	Function   string
}

// RoutesService struct
type routesService struct {
	BaseService
}
//NewRoutesService - instantiate de IModel filter
func NewRoutesService() routesService {
	var cs routesService

	return cs
}
func (rs *routesService) RegisteredRoutes(params url.Values) []*RegisteredRoute {
	var routes []*RegisteredRoute
	wbp := web.BeeApp.PrintTree()
	data := wbp["Data"].(web.M)
	methods := rs.Methods()

	methodParam := params.Get("_method")
	urlParam := params.Get("_keywords")
	controllerParam := params.Get("id") // for modal popup

	for _, item := range methods {
		m := data[item].(*[][]string)
		//log.Printf("%s  - %#v \n", item, m)
		for _, i := range *m { //[]string{"/admin/index/index", "map[GET:Index]", "controllers.IndexController"}

			url := i[0]
			f := i[1]
			c := i[2]
			if (len(methodParam) == 0) || (methodParam == item) { //Filter by _method and _keywords
				if len(urlParam) == 0 || strings.Contains(url, urlParam) {
					if len(controllerParam) == 0 || strings.Contains(c, controllerParam) {
						r := RegisteredRoute{Method: item, Url: url, Function: f, Controller: c}
						routes = append(routes, &r)
					}
				}
			}
		}
	}

	// log.Printf("Tree length %d \n", len(m))
	// for i, num := range m {
	// 	log.Printf("%s - %#v \n", i, num)
	// 	if i == "Data" {
	// 		log.Println("DATA")
	// 		for j, n := range num.(web.M) {
	// 			log.Printf("%s - %#v \n", j, n)
	// 		}
	// 	}
	// }

	return routes
}

func (*routesService) Methods() []string {

	wbp := web.BeeApp.PrintTree()
	if wbp != nil {
		//log.Printf("%#v", wbp["Methods"])
		return wbp["Methods"].([]string)
	}

	return nil
}

func (ars *routesService) GetPaginatedData(listRows int, params url.Values) ([]*RegisteredRoute, page.Pagination) {
	routes := ars.RegisteredRoutes(params)

	return routes, ars.Pagination

}
