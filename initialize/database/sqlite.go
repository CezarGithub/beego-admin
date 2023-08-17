package database

import (
	"fmt"
	"quince/global"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// init sqlite
func sqlite() {
	err := orm.RegisterDriver("sqlite", orm.DRSqlite)
	if err != nil {
		logs.Error("mysql register driver error:", err)
	}

	dataSource := global.BA_CONFIG.Sqlite.Database
	name := "default"

	err = orm.RegisterDataBase(name, "sqlite3", dataSource)
	if err != nil {
		logs.Error("sqlite register database error:", err)
	} else {
		logs.Info(fmt.Sprintf("[DATABASE]  sqlite database registered %s", dataSource))
	}

}
