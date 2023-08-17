package routers

import (
	//"net/http"
	"quince/initialize/router"
	"quince/middleware"
	"quince/modules/admin/api"
	"quince/modules/admin/controllers"

	"quince/internal/captcha"
)

func init() {

	//Authorized login middleware
	middleware.AuthMiddle()

	//Test middleware
	//middleware.TestMiddle2()

	// web.Get("/", func(ctx *context.Context) {
	// 	ctx.Redirect(http.StatusFound, "/admin/index/index")
	// })
	router.Redirect("/", "/admin/index/index")
	//admin Module routing
	admin := router.NewNamespace("admin",
		//UEditor Controller
		router.NSRouter("admin.route.uieditor", "OFF", "/editor/server", &controllers.EditorController{}, "get,post:Server"),

		//Login page
		router.NSRouter("admin.route.login", "OFF", "/auth/login", &controllers.AuthController{}, "get:Login"),
		//sign out
		router.NSRouter("admin.route.logout", "OFF", "/auth/logout", &controllers.AuthController{}, "get:Logout"),
		//QR code image output
		router.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
		//Login authentication
		router.NSRouter("admin.route.login", "OFF", "/auth/check_login", &controllers.AuthController{}, "post:CheckLogin"),
		//APILogin authentication - excepted from API
		//router.APIRouter("admin.route.APIlogin", router.Version1, "OFF", "/check_login", &api.ApiController{}, "post:CheckLogin"),
		router.NSRouter("admin.route.APIlogin", "OFF", "/api/login", &api.ApiController{}, "post:CheckLogin"),

		//Refresh Code
		router.NSRouter("admin.route.captcha", "OFF", "/auth/refresh_captcha", &controllers.AuthController{}, "post:RefreshCaptcha"),

		//front page
		router.NSRouter("app.route.index", "OFF", "/index/index", &controllers.IndexController{}, "get:Index"),

		//User Management
		router.NSRouter("app.route.index", "OFF", "/admin_user/index", &controllers.AdminUserController{}, "get:Index"),

		//Operation log
		router.NSRouter("app.route.index", "OFF", "/admin_log/index", &controllers.AdminLogController{}, "get:Index"),
		router.NSRouter("app.route.view", "OFF", "/admin_log/view", &controllers.AdminLogController{}, "get:View"),

		//Menu management
		router.NSRouter("app.route.index", "OFF", "/admin_menu/index", &controllers.AdminMenuController{}, "get:Index"),
		//Menu Management-Add Menu-Interface
		router.NSRouter("app.route.add", "OFF", "/admin_menu/add", &controllers.AdminMenuController{}, "get:Add"),
		//Menu Management-Add Menu-Create
		//web.NSRouter("/admin_menu/create", &controllers.AdminMenuController{}, "post:Create"),
		//Menu management-modify menu-interface
		router.NSRouter("app.route.edit", "OFF", "/admin_menu/edit", &controllers.AdminMenuController{}, "get:Edit"),
		//Menu Management-Update Menu
		router.NSRouter("app.route.update", "POST", "/admin_menu/update", &controllers.AdminMenuController{}, "post:Update"),
		//Menu management-delete menu
		router.NSRouter("app.route.delete", "POST", "/admin_menu/delete", &controllers.AdminMenuController{}, "post:Delete"),
		router.NSRouter("app.route.toggle", "POST", "/admin_menu/toggle", &controllers.AdminMenuController{}, "post:Toggle"),
		router.NSRouter("app.route.sidebar", "OFF", "/admin_menu/sidebar", &controllers.AdminMenuController{}, "get:LeftMenu"),

		//System Management-Personal Information
		router.NSRouter("admin.route.admin_user", "OFF", "/admin_user/profile", &controllers.AdminUserController{}, "get:Profile"),
		//System Management-Personal Information-Modify Nickname
		router.NSRouter("admin.route.admin_user", "POST", "/admin_user/update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
		//System Management-Personal Information-Change Password
		router.NSRouter("admin.route.admin_user", "POST", "/admin_user/update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
		//System Management-Personal Information-Modify Avatar
		router.NSRouter("admin.route.admin_user", "POST", "/admin_user/update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
		//System Management-User Management-Add Interface
		router.NSRouter("app.route.add", "POST", "/admin_user/add", &controllers.AdminUserController{}, "get:Add"),
		//System Management-User Management-Add
		router.NSRouter("app.route.toggle", "POST", "/admin_user/toggle", &controllers.AdminUserController{}, "post:Toggle"),
		//System Management-User Management-Modify Interface
		router.NSRouter("app.route.edit", "POST", "/admin_user/edit", &controllers.AdminUserController{}, "get:Edit"),
		//System Management-User Management-Modification
		router.NSRouter("app.route.update", "POST", "/admin_user/update", &controllers.AdminUserController{}, "post:Update"),
		//System Management-User Management-Enable
		//web.NSRouter("/admin_user/enable", &controllers.AdminUserController{}, "post:Enable"),
		//System Management-User Management-Disable
		//web.NSRouter("/admin_user/disable", &controllers.AdminUserController{}, "post:Disable"),
		//System Management-User Management-Delete
		router.NSRouter("app.route.delete", "POST", "/admin_user/del", &controllers.AdminUserController{}, "post:Del"),

		//System Management-Role Management
		router.NSRouter("app.route.index", "OFF", "/admin_role/index", &controllers.AdminRoleController{}, "get:Index"),
		//System Management-Role Management-Add Interface
		router.NSRouter("app.route.add", "POST", "/admin_role/add", &controllers.AdminRoleController{}, "get:Add"),
		//System Management-Role Management-Add
		router.NSRouter("app.route.toggle", "POST", "/admin_role/toggle", &controllers.AdminRoleController{}, "post:Toggle"),
		//web.NSRouter("/admin_role/create", &controllers.AdminRoleController{}, "post:Create"),
		//Menu Management-Role Management-Modify Interface
		router.NSRouter("app.route.edit", "POST", "/admin_role/edit", &controllers.AdminRoleController{}, "get:Edit"),
		//Menu Management-Role Management-Modify
		router.NSRouter("app.route.update", "POST", "/admin_role/update", &controllers.AdminRoleController{}, "post:Update"),
		//Menu Management-Role Management-Delete
		router.NSRouter("app.route.delete", "POST", "/admin_role/del", &controllers.AdminRoleController{}, "post:Delete"),
		//Menu Management-Role Management-Enable Role
		//web.NSRouter("/admin_role/enable", &controllers.AdminRoleController{}, "post:Enable"),
		//Menu Management-Role Management-Disable Role
		//web.NSRouter("/admin_role/disable", &controllers.AdminRoleController{}, "post:Disable"),
		//Menu Management-Role Management-Role Authorization Interface
		router.NSRouter("admin.route.admin_role", "OFF", "/admin_role/access", &controllers.AdminRoleController{}, "get:Menus"),
		router.NSRouter("admin.route.admin_user", "OFF", "/admin_role/route", &controllers.AdminRoleController{}, "get,post:Routes"),
		//Menu Management-Role Management-Role Authorization
		router.NSRouter("admin.route.admin_user", "POST", "/admin_role/access_operate", &controllers.AdminRoleController{}, "post:AccessOperate"),
		router.NSRouter("admin.route.admin_user", "POST", "/admin_role/rbac_operate", &controllers.AdminRoleController{}, "post:AccessRbac"),
		//Settings Center-Background Settings
		router.NSRouter("admin.route.settings", "OFF", "/setting/admin", &controllers.SettingController{}, "get:Admin"),
		//Settings Center-Update Settings
		router.NSRouter("admin.route.settings", "POST", "/setting/update", &controllers.SettingController{}, "post:Update"),

		//System Management-Development Management-Data Maintenance
		router.NSRouter("admin.route.database", "OFF", "/database/table", &controllers.DatabaseController{}, "get:Table"),
		//System Management-Development Management-Data Maintenance-Optimization Table
		router.NSRouter("admin.route.database", "POST", "/database/optimize", &controllers.DatabaseController{}, "post:Optimize"),
		//System Management-Development Management-Data Maintenance-Repair Table
		router.NSRouter("admin.route.database", "OFF", "/database/repair", &controllers.DatabaseController{}, "post:Repair"),
		//System Management-Development Management-Data Maintenance-View Details
		router.NSRouter("admin.route.database", "OFF", "/database/view", &controllers.DatabaseController{}, "get,post:View"),
		router.NSRouter("admin.route.database", "OFF", "/database/export-json", &controllers.DatabaseController{}, "get,post:ExportDatabaseToJSON"),
		router.NSRouter("admin.route.database", "OFF", "/database/import-json", &controllers.DatabaseController{}, "get,post:ImportDatabaseFromJSON"),
		router.NSRouter("admin.route.database", "OFF", "/database/init-db", &controllers.DatabaseController{}, "get,post:InitDatabaseFromJSON"),

		//Code generation
		router.NSRouter("admin.route.generate", "OFF", "/generate/index", &controllers.CodeGenerationController{}, "get:Index"),
		router.NSRouter("admin.route.generate", "OFF", "/generate/new", &controllers.CodeGenerationController{}, "get,post:New"),
		router.NSRouter("admin.route.generate", "OFF", "/generate/doc", &controllers.CodeGenerationController{}, "get:Doc"),
		//User level management
		router.NSRouter("admin.route.user_level", "OFF", "/user_level/index", &controllers.UserLevelController{}, "get:Index"),
		//User level management-add interface
		router.NSRouter("admin.route.user_level", "OFF", "/user_level/add", &controllers.UserLevelController{}, "get:Add"),
		//User level management-add
		//web.NSRouter("/user_level/create", &controllers.UserLevelController{}, "post:Create"),
		//User level management-modify interface
		router.NSRouter("admin.route.user_level", "OFF", "/user_level/edit", &controllers.UserLevelController{}, "get:Edit"),
		//User level management-modification
		router.NSRouter("admin.route.user_level", "POST", "/user_level/update", &controllers.UserLevelController{}, "post:Update"),
		//User level management-enable
		router.NSRouter("admin.route.user_level", "OFF", "/user_level/toggle", &controllers.UserLevelController{}, "post:Toggle"),
		//User level management-disabled
		//web.NSRouter("/user_level/disable", &controllers.UserLevelController{}, "post:Disable"),
		//User level management-delete
		router.NSRouter("admin.route.user_level", "OFF", "/user_level/del", &controllers.UserLevelController{}, "post:Del"),
		//User level management-export
		router.NSRouter("admin.route.user_level", "OFF", "/user_level/export", &controllers.UserLevelController{}, "get:Export"),

		//Cron jobs Management
		router.NSRouter("admin.cron.title", "OFF", "/cronjob/index", &controllers.AdminCronJobController{}, "get:Index"),
		//Cron jobs management-add interface
		router.NSRouter("admin.cron.title", "OFF", "/cronjob/add", &controllers.AdminCronJobController{}, "get:Add"),
		//Cron jobs management-modify interface
		router.NSRouter("admin.cron.title", "OFF", "/cronjob/edit", &controllers.AdminCronJobController{}, "get:Edit"),
		//Cron jobs Management-Modification
		router.NSRouter("admin.cron.title", "POST", "/cronjob/update", &controllers.AdminCronJobController{}, "post:Update"),
		//Cron jobs Management-Enable
		router.NSRouter("admin.cron.title", "OFF", "/cronjob/toggle", &controllers.AdminCronJobController{}, "post:Toggle"),
		router.NSRouter("admin.cron.title", "POST", "/cronjob/delete", &controllers.AdminCronJobController{}, "post:Delete"),
		router.NSRouter("app.route.search", "OFF", "/cronjob/search", &controllers.AdminCronJobController{}, "get,post:Search"),
		//User Management
		router.NSRouter("admin.route.user", "OFF", "/user/index", &controllers.UserController{}, "get:Index"),
		//User management-add interface
		router.NSRouter("admin.route.user", "OFF", "/user/add", &controllers.UserController{}, "get:Add"),
		//User Management-Add
		//web.NSRouter("/user/create", &controllers.UserController{}, "post:Create"),
		//User management-modify interface
		router.NSRouter("admin.route.user", "OFF", "/user/edit", &controllers.UserController{}, "get:Edit"),
		//User Management-Modification
		router.NSRouter("admin.route.user", "POST", "/user/update", &controllers.UserController{}, "post:Update"),
		//User Management-Enable
		router.NSRouter("admin.route.user", "OFF", "/user/toggle", &controllers.UserController{}, "post:Toggle"),
		//User Management-Disable
		//web.NSRouter("/user/disable", &controllers.UserController{}, "post:Disable"),
		//User Management-Delete
		router.NSRouter("admin.route.user", "OFF", "/user/del", &controllers.UserController{}, "post:Del"),
		//User Management-Export
		router.NSRouter("admin.route.user", "OFF", "/user/export", &controllers.UserController{}, "get:Export"),
		//Search
		router.NSRouter("admin.route.user", "OFF", "/user/search", &controllers.UserController{}, "get:Search"),

		//Registered routes list
		router.NSRouter("admin.menu.routes", "OFF", "/route/index", &controllers.RouteController{}, "get:Index"),
		router.NSRouter("admin.menu.routes", "OFF", "/route/view", &controllers.RouteController{}, "get:View"),
	)

	router.AddNamespace(admin)

}

// admin := web.NewNamespace("/admin",
// //UEditor Controller
// web.NSRouter("/editor/server", &controllers.EditorController{}, "get,post:Server"),

// //Login page
// web.NSRouter("/auth/login", &controllers.AuthController{}, "get:Login"),
// //sign out
// web.NSRouter("/auth/logout", &controllers.AuthController{}, "get:Logout"),
// //QR code image output
// web.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
// //Login authentication
// web.NSRouter("/auth/check_login", &controllers.AuthController{}, "post:CheckLogin"),
// //Refresh Code
// web.NSRouter("/auth/refresh_captcha", &controllers.AuthController{}, "post:RefreshCaptcha"),

// //front page
// web.NSRouter("/index/index", &controllers.IndexController{}, "get:Index"),

// //User Management
// web.NSRouter("/admin_user/index", &controllers.AdminUserController{}, "get:Index"),

// //Operation log
// web.NSRouter("/admin_log/index", &controllers.AdminLogController{}, "get:Index"),
// web.NSRouter("/admin_log/view", &controllers.AdminLogController{}, "get:View"),

// //Menu management
// web.NSRouter("/admin_menu/index", &controllers.AdminMenuController{}, "get:Index"),
// //Menu Management-Add Menu-Interface
// web.NSRouter("/admin_menu/add", &controllers.AdminMenuController{}, "get:Add"),
// //Menu Management-Add Menu-Create
// //web.NSRouter("/admin_menu/create", &controllers.AdminMenuController{}, "post:Create"),
// //Menu management-modify menu-interface
// web.NSRouter("/admin_menu/edit", &controllers.AdminMenuController{}, "get:Edit"),
// //Menu Management-Update Menu
// web.NSRouter("/admin_menu/update", &controllers.AdminMenuController{}, "post:Update"),
// //Menu management-delete menu
// web.NSRouter("/admin_menu/delete", &controllers.AdminMenuController{}, "post:Delete"),
// web.NSRouter("/admin_menu/toggle", &controllers.AdminMenuController{}, "post:Toggle"),

// //System Management-Personal Information
// web.NSRouter("/admin_user/profile", &controllers.AdminUserController{}, "get:Profile"),
// //System Management-Personal Information-Modify Nickname
// web.NSRouter("/admin_user/update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
// //System Management-Personal Information-Change Password
// web.NSRouter("/admin_user/update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
// //System Management-Personal Information-Modify Avatar
// web.NSRouter("/admin_user/update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
// //System Management-User Management-Add Interface
// web.NSRouter("/admin_user/add", &controllers.AdminUserController{}, "get:Add"),
// //System Management-User Management-Add
// web.NSRouter("/admin_user/toggle", &controllers.AdminUserController{}, "post:Toggle"),
// //System Management-User Management-Modify Interface
// web.NSRouter("/admin_user/edit", &controllers.AdminUserController{}, "get:Edit"),
// //System Management-User Management-Modification
// web.NSRouter("/admin_user/update", &controllers.AdminUserController{}, "post:Update"),
// //System Management-User Management-Enable
// //web.NSRouter("/admin_user/enable", &controllers.AdminUserController{}, "post:Enable"),
// //System Management-User Management-Disable
// //web.NSRouter("/admin_user/disable", &controllers.AdminUserController{}, "post:Disable"),
// //System Management-User Management-Delete
// web.NSRouter("/admin_user/del", &controllers.AdminUserController{}, "post:Del"),

// //System Management-Role Management
// web.NSRouter("/admin_role/index", &controllers.AdminRoleController{}, "get:Index"),
// //System Management-Role Management-Add Interface
// web.NSRouter("/admin_role/add", &controllers.AdminRoleController{}, "get:Add"),
// //System Management-Role Management-Add
// web.NSRouter("/admin_role/toggle", &controllers.AdminRoleController{}, "post:Toggle"),
// //web.NSRouter("/admin_role/create", &controllers.AdminRoleController{}, "post:Create"),
// //Menu Management-Role Management-Modify Interface
// web.NSRouter("/admin_role/edit", &controllers.AdminRoleController{}, "get:Edit"),
// //Menu Management-Role Management-Modify
// web.NSRouter("/admin_role/update", &controllers.AdminRoleController{}, "post:Update"),
// //Menu Management-Role Management-Delete
// web.NSRouter("/admin_role/del", &controllers.AdminRoleController{}, "post:Delete"),
// //Menu Management-Role Management-Enable Role
// //web.NSRouter("/admin_role/enable", &controllers.AdminRoleController{}, "post:Enable"),
// //Menu Management-Role Management-Disable Role
// //web.NSRouter("/admin_role/disable", &controllers.AdminRoleController{}, "post:Disable"),
// //Menu Management-Role Management-Role Authorization Interface
// web.NSRouter("/admin_role/access", &controllers.AdminRoleController{}, "get:Menus"),
// web.NSRouter("/admin_role/route", &controllers.AdminRoleController{}, "get,post:Routes"),
// //Menu Management-Role Management-Role Authorization
// web.NSRouter("/admin_role/access_operate", &controllers.AdminRoleController{}, "post:AccessOperate"),
// web.NSRouter("/admin_role/rbac_operate", &controllers.AdminRoleController{}, "post:AccessRbac"),
// //Settings Center-Background Settings
// web.NSRouter("/setting/admin", &controllers.SettingController{}, "get:Admin"),
// //Settings Center-Update Settings
// web.NSRouter("/setting/update", &controllers.SettingController{}, "post:Update"),

// //System Management-Development Management-Data Maintenance
// web.NSRouter("/database/table", &controllers.DatabaseController{}, "get:Table"),
// //System Management-Development Management-Data Maintenance-Optimization Table
// web.NSRouter("/database/optimize", &controllers.DatabaseController{}, "post:Optimize"),
// //System Management-Development Management-Data Maintenance-Repair Table
// web.NSRouter("/database/repair", &controllers.DatabaseController{}, "post:Repair"),
// //System Management-Development Management-Data Maintenance-View Details
// web.NSRouter("/database/view", &controllers.DatabaseController{}, "get,post:View"),
// web.NSRouter("/database/export-json", &controllers.DatabaseController{}, "get,post:ExportDatabaseToJSON"),
// web.NSRouter("/database/import-json", &controllers.DatabaseController{}, "get,post:ImportDatabaseFromJSON"),
// //User level management
// web.NSRouter("/user_level/index", &controllers.UserLevelController{}, "get:Index"),
// //User level management-add interface
// web.NSRouter("/user_level/add", &controllers.UserLevelController{}, "get:Add"),
// //User level management-add
// //web.NSRouter("/user_level/create", &controllers.UserLevelController{}, "post:Create"),
// //User level management-modify interface
// web.NSRouter("/user_level/edit", &controllers.UserLevelController{}, "get:Edit"),
// //User level management-modification
// web.NSRouter("/user_level/update", &controllers.UserLevelController{}, "post:Update"),
// //User level management-enable
// web.NSRouter("/user_level/toggle", &controllers.UserLevelController{}, "post:Toggle"),
// //User level management-disabled
// //web.NSRouter("/user_level/disable", &controllers.UserLevelController{}, "post:Disable"),
// //User level management-delete
// web.NSRouter("/user_level/del", &controllers.UserLevelController{}, "post:Del"),
// //User level management-export
// web.NSRouter("/user_level/export", &controllers.UserLevelController{}, "get:Export"),

// //User Management
// web.NSRouter("/user/index", &controllers.UserController{}, "get:Index"),
// //User management-add interface
// web.NSRouter("/user/add", &controllers.UserController{}, "get:Add"),
// //User Management-Add
// //web.NSRouter("/user/create", &controllers.UserController{}, "post:Create"),
// //User management-modify interface
// web.NSRouter("/user/edit", &controllers.UserController{}, "get:Edit"),
// //User Management-Modification
// web.NSRouter("/user/update", &controllers.UserController{}, "post:Update"),
// //User Management-Enable
// web.NSRouter("/user/toggle", &controllers.UserController{}, "post:Toggle"),
// //User Management-Disable
// //web.NSRouter("/user/disable", &controllers.UserController{}, "post:Disable"),
// //User Management-Delete
// web.NSRouter("/user/del", &controllers.UserController{}, "post:Del"),
// //User Management-Export
// web.NSRouter("/user/export", &controllers.UserController{}, "get:Export"),
// //Search
// web.NSRouter("/user/search", &controllers.UserController{}, "get:Search"),

// //Registered routes list
// web.NSRouter("/route/index", &controllers.RouteController{}, "get:Index"),
// web.NSRouter("/route/view", &controllers.RouteController{}, "get:View"),
// )

// web.AddNamespace(admin)
