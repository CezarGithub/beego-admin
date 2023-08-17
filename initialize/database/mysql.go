package database

import (
	"fmt"
	"quince/global"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// init mysql
func mysql() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error("mysql register driver error:", err)
	}

	//dataSource := "root:root@tcp(127.0.0.1:3306)/test"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		global.BA_CONFIG.Mysql.Username,
		global.BA_CONFIG.Mysql.Password,
		global.BA_CONFIG.Mysql.Host,
		global.BA_CONFIG.Mysql.Port,
		global.BA_CONFIG.Mysql.Database,
	)
	name := "default"
	err = orm.RegisterDataBase(name, "mysql", dataSource)
	if err != nil {
		logs.Error("mysql register database error:", err)
		panic("Fatal database error. Check logs...")
	} else {
		logs.Info(fmt.Sprintf("[DATABASE] mysql database registered %s", global.BA_CONFIG.Mysql.Database))
	}

}
