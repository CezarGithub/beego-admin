package admin

import (
	"encoding/base64"
	"fmt"
	"quince/initialize/module"
	"quince/internal/i18n"
	_ "quince/modules/admin/cron"
	"quince/modules/admin/models"
	_ "quince/modules/admin/routers"
	"quince/utils"
	"time"

	"github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/client/orm"

	"github.com/beego/beego/v2/core/logs"
)

type admin struct {
	module.IModule
}

func init() {
	a := *new(admin)
	module.Register("admin", a)
}
func (m admin) Init(tx orm.TxOrmer) error {
	var err error
	var cnt int64
	logs.Info("Module admin initialisation")
	//super_admin
	t, ee := tx.QueryTable(new(models.LoginUser).TableName()).Count()

	if t == 0 && ee == nil {
		v := models.UserLevel{}
		v.Description = "Administrators level"
		v.Name, _ = web.AppConfig.String("adminname")
		v.Status = 1
		_, err = tx.Insert(&v)
		if err != nil {
			logs.Error("Default admin user error %v", err)
		}
		u := models.User{}
		u.Username, _ = web.AppConfig.String("adminname")
		u.Description = "Default admin user"
		u.Nickname, _ = web.AppConfig.String("appname")
		u.Email, _ = web.AppConfig.String("email") //"suport@kodis.ro"
		u.Status = 1
		u.Avatar = "/static/admin/images/avatar64.png"
		u.Base = models.Base{CreateTime: time.Now()}
		u.UserLevel = &v
		_, err = tx.Insert(&u)
		lu := models.LoginUser{}
		lu.Base = models.Base{}
		lu.User = &u
		lu.Base.CreateTime = time.Now()
		lu.Base.UpdateTime = time.Now()
		lu.LoginName, _ = web.AppConfig.String("adminname") // "super_admin"
		pass, _ := web.AppConfig.String("adminname")
		newPasswordForHash, err := utils.PasswordHash(pass)
		if err == nil {
			lu.Password = base64.StdEncoding.EncodeToString([]byte(newPasswordForHash))
		} else {
			return err
		}
		lu.Status = 1
		// l, _ := web.AppConfig.String("lang::alpha4")
		// langs := strings.Split(l, "|")
		langs := i18n.ApplicationLanguages()
		lu.Language = langs[0] //set admin language as default language of app from conf ex.: "en-US"
		id, er := tx.Insert(&lu)
		if id != 1 || er != nil {
			logs.Error("Default admin user id : %d , error %v", id, er.Error())
		}
	}

	//menus
	cnt, er := tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", "admin").Count()
	if cnt == 0 && er == nil {
		m1 := models.AdminMenu{Name: "app.home", Url: "admin/index/index", Icon: "fa-home", Module: "admin", IsShow: 1, Status: 1, SortId: 10, LogMethod: "OFF"}
		tx.Insert(&m1)
		//2
		m2 := models.AdminMenu{Name: "admin.menu.system", Url: "admin/sys", Icon: "fa-desktop", Module: "admin", IsShow: 1, Status: 1, SortId: 20, LogMethod: "OFF"}
		id, _ := tx.Insert(&m2)
		m21 := models.AdminMenu{Name: "admin.menu.login_user", Url: "admin/admin_user/index", Icon: "fa-user", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&m21)
		m22 := models.AdminMenu{Name: "admin.menu.role", Url: "admin/admin_role/index", Icon: "fa-group", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&m22)
		m23 := models.AdminMenu{Name: "admin.menu.menu", Url: "admin/admin_menu/index", Icon: "fa-align-justify", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&m23)
		m24 := models.AdminMenu{Name: "admin.menu.logs", Url: "admin/admin_log/index", Icon: "fa-keyboard-o", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&m24)
		m25 := models.AdminMenu{Name: "admin.menu.personal", Url: "admin/admin_user/profile", Icon: "fa-smile-o", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&m25)
		//dev
		m26 := models.AdminMenu{Name: "admin.menu.dev", Url: "admin/develop/manager", Icon: "fa-code", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		dev, _ := tx.Insert(&m26)
		dev1 := models.AdminMenu{Name: "admin.menu.db_tables", Url: "admin/database/table", Icon: "fa-database", ParentId: dev, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&dev1)
		dev2 := models.AdminMenu{Name: "admin.menu.cronjobs", Url: "admin/cronjob/index", Icon: "fa-clock-o", ParentId: dev, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&dev2)
		dev3 := models.AdminMenu{Name: "admin.menu.generate", Url: "admin/generate/index", Icon: "fa-terminal", ParentId: dev, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&dev3)
		//settings
		m27 := models.AdminMenu{Name: "admin.menu.settings_area", Url: "admin/setting/center", Icon: "fa-cogs", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		set, _ := tx.Insert(&m27)
		set1 := models.AdminMenu{Name: "admin.menu.settings", Url: "admin/setting/admin", Icon: "fa-adjust", ParentId: set, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&set1)

		m28 := models.AdminMenu{Name: "admin.menu.routes", Url: "admin/route/index", Icon: "fa-link", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&m28)
		//3
		m3 := models.AdminMenu{Name: "admin.menu.user", Url: "admin/user/mange", Icon: "fa-users", Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		id, _ = tx.Insert(&m3)
		m31 := models.AdminMenu{Name: "admin.menu.users", Url: "admin/user/index", Icon: "fa-user", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&m31)
		m32 := models.AdminMenu{Name: "admin.menu.user_level", Url: "admin/user_level/index", Icon: "fa-th-list", ParentId: id, Module: "admin", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
		tx.Insert(&m32)
	}
	if er != nil {
		err = er
	}

	c, e := tx.QueryTable(new(models.Setting).TableName()).Count()
	if c == 0 && e == nil {
		m1 := models.SettingGroup{Module: "admin", Name: "Settings", Description: "Settings management", SortNumber: 1, AutoCreateMenu: 1, AutoCreateFile: 1, Icon: "fa-adjust"}
		id, _ := tx.Insert(&m1)
		content := fmt.Sprintf("[{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}", "Name", "Name", "Field", "name", "Type", "text", "Content", "⚡⚡⚡", "Option", "", "Form", "")
		content += fmt.Sprintf(",{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}", "Name", "Abbreviation", "Field", "short_name", "Type", "text", "Content", "⛬", "Option", "", "Form", "")
		content += fmt.Sprintf(",{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}", "Name", "Author", "Field", "author", "Type", "text", "Content", "quince", "Option", "", "Form", "")
		content += fmt.Sprintf(",{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}]", "Name", "Version", "Field", "version", "Type", "text", "Content", "0.1", "Option", "", "Form", "")
		s1 := models.Setting{SettingGroupId: id, Name: "basic_settings", Description: "Basic information settings ", Code: "base", Content: content, SortNumber: 1000}
		tx.Insert(&s1)

		content = fmt.Sprintf("[{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}", "Name", "Login token verification", "Field", "token", "Type", "switch", "Content", "1", "Option", "", "Form", "")
		content += fmt.Sprintf(",{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}", "Name", "Verification code", "Field", "captcha", "Type", "select", "Content", "1", "Option", "0||Do not open\r\n1||Graphic verification code", "Form", "")
		content += fmt.Sprintf(",{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}]", "Name", "Login background", "Field", "background", "Type", "image", "Content", "/static/uploads/attachment/ea08c391-0eb4-4c6a-8e4f-9846c51d61cc.jpg", "Option", "", "Form", "")
		s2 := models.Setting{SettingGroupId: id, Name: "login_settings", Description: "Login related settings  ", Code: "login", Content: content, SortNumber: 1}
		tx.Insert(&s2)

		content = fmt.Sprintf("[{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}", "Name", "Default password warning", "Field", "password_warning", "Type", "switch", "Content", "1", "Option", "", "Form", "")
		content += fmt.Sprintf(",{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}", "Name", "Whether to display prompt information", "Field", "show_notice", "Type", "switch", "Content", "1", "Option", "", "Form", "")
		content += fmt.Sprintf(",{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:%q}]", "Name", "Prompt message content", "Field", "notice_content", "Type", "text", "Content", "Welcome to use this system, the left side is the menu area, and the right side is the function area。", "Option", "", "Form", "")
		s3 := models.Setting{SettingGroupId: id, Name: "home_page_settings", Description: "Home page parameter setting ", Code: "index", Content: content, SortNumber: 1}
		tx.Insert(&s3)

	}
	if er != nil {
		err = e
	}
	return err
}
