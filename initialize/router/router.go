package router

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"

	"quince/modules/admin/models"
	s "quince/utils/string"

	"github.com/beego/beego/v2/client/orm"

	"github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

// override Beego web functions
func init() {
	logs.Info("Internal router init")
}

//	type route struct {
//		models.AdminRoute
//	}
type routes struct {
	module string
	List   []models.AdminRoute
}

var moduleRoutes routes
var appRoutes []routes

type Namespace struct {
	*web.Namespace
}

func DBSync() {
	logs.Info("[DATABASE] Start routes db sync for %d modules ", len(appRoutes))
	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		var er error

		_, er = txOrm.Raw("UPDATE admin_route SET status = 0").Exec() //set all existing routes disabled
		for _, m := range appRoutes {
			inserted := 0
			updated := 0
			for _, r := range m.List {
				var route models.AdminRoute
				url:="/" + m.module + r.Url
				_ = txOrm.Raw("SELECT id FROM admin_route WHERE url = ? ",url).QueryRow(&route) //ignore error for no rows found
				r.Id = route.Id
				r.Module = m.module
				r.Url = "/" + m.module + r.Url
				r.Status = 1
				if r.Id == 0 {
					r.CreateTime = time.Now()
					inserted += 1
					txOrm.Insert(&r)
				} else {
					r.UpdateTime = time.Now()
					updated += 1
					txOrm.Update(&r)
				}

				//logs.Info("Module : %s - route : %s", m.module, r.Url)
			}
			logs.Info("Module : %s -> routes inserted : %d , updated : %d", m.module, inserted, updated)
		}

		return er
	})
	if err != nil {
		logs.Error("[DATABASE] routes db sync failed %v", err.Error())
		panic("[DATABASE] routes db sync failed")
	}
}

func NewNamespace(moduleName string, params ...web.LinkNamespace) *web.Namespace {
	prefix := s.Alphanumeric(moduleName)
	ns := web.NewNamespace(prefix, params...)

	moduleRoutes.module = moduleName
	return ns
}

func AddNamespace(nl ...*web.Namespace) {
	appRoutes = append(appRoutes, moduleRoutes)
	moduleRoutes = routes{}
	web.AddNamespace(nl...)
}

// see web.Get
//
//	web.Get("/", func(ctx *context.Context) {
//		ctx.Redirect(http.StatusFound, "/admin/index/index")
//	})
func Redirect(rootpath string, localpath string) {
	web.Get(rootpath, func(ctx *beecontext.Context) {
		ctx.Redirect(http.StatusFound, localpath)
	})

}

// NSBefore Namespace BeforeRouter filter
func NSBefore(filterList ...web.FilterFunc) web.LinkNamespace {
	return web.NSBefore(filterList...)
}

// NSAfter add Namespace FinishRouter filter
func NSAfter(filterList ...web.FilterFunc) web.LinkNamespace {
	return web.NSAfter(filterList...)
}

// NSInclude Namespace Include ControllerInterface
func NSInclude(cList ...web.ControllerInterface) web.LinkNamespace {
	return web.NSInclude(cList...)
}

// NSRouter call Namespace Router
func NSRouter(description string, log_method string, rootpath string, c web.ControllerInterface, mappingMethods ...string) web.LinkNamespace {
	t := reflect.TypeOf(c)
	r := models.AdminRoute{Url: rootpath, IsAPI: 0, Name: description, LogMethod: log_method, MappingMethod: strings.Join(mappingMethods, ","), Controller: t.String()}
	moduleRoutes.List = append(moduleRoutes.List, r)
	return web.NSRouter(rootpath, c, mappingMethods...)
}

// Wrapper for API routes
// Header ["Authorization"] = X-API-Key
// Header [X-API-Key] contains JWT key
// Header [X-Requested-With] = XMLHttpRequest
// POST method only
// Returns myController.ResponseSuccessWithData("My message", data, myController.Ctx) or ResponseWIthError...
func APIRouter(description string, version API, log_method string, rootpath string, c web.ControllerInterface, mappingMethods ...string) web.LinkNamespace {
	rootpath = "/api" + "/" + fmt.Sprintf("%d", version) + rootpath
	t := reflect.TypeOf(c)
	r := models.AdminRoute{Url: rootpath, IsAPI: 1, Name: description, LogMethod: log_method, MappingMethod: strings.Join(mappingMethods, ","), Controller: t.String()}
	moduleRoutes.List = append(moduleRoutes.List, r)
	return web.NSRouter(rootpath, c, mappingMethods...)
}

// NSGet call Namespace Get
func NSGet(rootpath string, f web.HandleFunc) web.LinkNamespace {
	return web.NSGet(rootpath, f)
}

// NSPost call Namespace Post
func NSPost(rootpath string, f web.HandleFunc) web.LinkNamespace {
	return web.NSPost(rootpath, f)
}

// NSHead call Namespace Head
func NSHead(rootpath string, f web.HandleFunc) web.LinkNamespace {
	return web.NSHead(rootpath, f)
}

// NSPut call Namespace Put
func NSPut(rootpath string, f web.HandleFunc) web.LinkNamespace {
	return web.NSPut(rootpath, f)
}

// NSDelete call Namespace Delete
func NSDelete(rootpath string, f web.HandleFunc) web.LinkNamespace {
	return web.NSDelete(rootpath, f)
}

// NSAny call Namespace Any
func NSAny(rootpath string, f web.HandleFunc) web.LinkNamespace {
	return web.NSAny(rootpath, f)
}

// NSOptions call Namespace Options
func NSOptions(rootpath string, f web.HandleFunc) web.LinkNamespace {
	return web.NSOptions(rootpath, f)
}

// NSPatch call Namespace Patch
func NSPatch(rootpath string, f web.HandleFunc) web.LinkNamespace {
	return web.NSPatch(rootpath, f)
}

// NSCtrlGet call Namespace CtrlGet
func NSCtrlGet(rootpath string, f interface{}) web.LinkNamespace {
	return web.NSCtrlGet(rootpath, f)
}

// NSCtrlPost call Namespace CtrlPost
func NSCtrlPost(rootpath string, f interface{}) web.LinkNamespace {
	return web.NSCtrlPost(rootpath, f)
}

// NSCtrlHead call Namespace CtrlHead
func NSCtrlHead(rootpath string, f interface{}) web.LinkNamespace {
	return web.NSCtrlHead(rootpath, f)
}

// NSCtrlPut call Namespace CtrlPut
func NSCtrlPut(rootpath string, f interface{}) web.LinkNamespace {
	return web.NSCtrlPut(rootpath, f)
}

// NSCtrlDelete call Namespace CtrlDelete
func NSCtrlDelete(rootpath string, f interface{}) web.LinkNamespace {
	return web.NSCtrlDelete(rootpath, f)
}

// NSCtrlAny call Namespace CtrlAny
func NSCtrlAny(rootpath string, f interface{}) web.LinkNamespace {
	return web.NSCtrlAny(rootpath, f)
}

// NSCtrlOptions call Namespace CtrlOptions
func NSCtrlOptions(rootpath string, f interface{}) web.LinkNamespace {
	return web.NSCtrlOptions(rootpath, f)
}

// NSCtrlPatch call Namespace CtrlPatch
func NSCtrlPatch(rootpath string, f interface{}) web.LinkNamespace {
	return web.NSCtrlPatch(rootpath, f)
}

// NSAutoRouter call Namespace AutoRouter
func NSAutoRouter(c web.ControllerInterface) web.LinkNamespace {
	return web.NSAutoRouter(c)
}

// NSAutoPrefix call Namespace AutoPrefix
func NSAutoPrefix(prefix string, c web.ControllerInterface) web.LinkNamespace {
	return web.NSAutoPrefix(prefix, c)
}

// NSNamespace add sub Namespace
func NSNamespace(prefix string, params ...web.LinkNamespace) web.LinkNamespace {
	return web.NSNamespace(prefix, params...)
}

// NSHandler add handler
func NSHandler(rootpath string, h http.Handler) web.LinkNamespace {
	return web.NSHandler(rootpath, h)
}
