package database

import (
	"fmt"
	"quince/global"
	"quince/internal/models"

	"github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

var modelFiles *DataModel

// init database
func init() {
	f, _ := web.AppConfig.String("log::system")
	logs.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(`{"filename":"log/%s","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true,"rotate":true}`, f))
	logs.Info("Init database")
	switch db := global.BA_CONFIG.Database.Type; db {
	case "mysql":
		mysql()
	case "sqlite":
		sqlite()
	default:
		logs.Error("Database configuration error !!!")
	}

	modelFiles = NewDataModelFile()
}

func DbSync() {
	name := "default"
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := true
	//check is new DB
	b := isNewDB()
	if b {
		logs.Info("New database started...")
	}

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		logs.Error("[DATABASE] sync database error:", err)
	} else {
		logs.Info("[DATABASE] sync database done")

	}
	initData()

	o := orm.NewOrm()
	if o.Driver().Type() == orm.DRSqlite {
		sql := fmt.Sprintf("PRAGMA user_version=%d;", global.BA_CONFIG.Sqlite.Version)
		_, err = o.Raw(sql).Exec()
		if err != nil {
			logs.Error("sqlite register version error:", err)
		}
	}
}

func RegisterModel(module string, a models.IModel) {
	orm.RegisterModel(a)
	modelFiles.Register(module, a)
}
func ExportData() []error {
	return modelFiles.Export()
}
func ImportData(folder string,module string) error {
	return modelFiles.Import(folder,module)
}
func AvailableImports() []string {
	return modelFiles.AvailableImports()
}
func RegisteredModels()map[string]models.IModel{
	return modelFiles.Items
}

// trigger initData method from models.
// initData() method is implemented in models.Base and must be overrited in model struct
func initData() error {
	return modelFiles.InitData()
}

func isNewDB() bool {
	b := true
	var maps []orm.Params
	o := orm.NewOrm()
	if o.Driver().Type() == orm.DRMySQL {
		affectRows, err := o.Raw("SHOW TABLE STATUS").Values(&maps)

		if affectRows > 0 && err == nil {
			b = false
		}
	}
	if o.Driver().Type() == orm.DRSqlite {
		sql := "SELECT name FROM  sqlite_master WHERE type ='table' AND name NOT LIKE 'sqlite_%';"
		affectRows, err := o.Raw(sql).Values(&maps)
		if affectRows > 0 && err == nil {
			b = false
		}
	}
	return b
}
