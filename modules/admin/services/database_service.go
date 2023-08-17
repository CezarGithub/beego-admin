package services

import (
	"fmt"
	"quince/initialize/database"

	"github.com/beego/beego/v2/client/orm"
)

// DbVersion struct
type DbVersion struct {
	DbVersion string
}

// DatabaseService struct
type databaseService struct {
}

// AdminLogService - instantiate de IModel filter
func NewDatabaseService() databaseService {
	var cs databaseService
	return cs
}

// GetMysqlVersion
func (*databaseService) GetMysqlVersion() string {
	var dbVersion DbVersion
	o := orm.NewOrm()
	dr := o.Driver()
	if dr.Type() == orm.DRMySQL {
		err := o.Raw("select VERSION() as db_version").QueryRow(&dbVersion)
		if err != nil {
			return "-"
		}
		dbVersion.DbVersion = fmt.Sprintf("MySql - v.%s", dbVersion.DbVersion)
	}
	if dr.Type() == orm.DRSqlite {
		var s string
		err := o.Raw("PRAGMA user_version;").QueryRow(&s)
		if err != nil {
			return "-"
		}
		dbVersion.DbVersion = fmt.Sprintf("SQLite - v.%s", s)
	}
	return dbVersion.DbVersion
}

// GetTableStatus
func (ds *databaseService) GetTableStatus() ([]map[string]string, int) {
	var maps []orm.Params
	var resultMaps []map[string]string
	o := orm.NewOrm()
	if o.Driver().Type() == orm.DRMySQL {
		affectRows, err := o.Raw("SHOW TABLE STATUS").Values(&maps)

		if affectRows > 0 && err == nil {
			for _, item := range maps {
				resultMaps = append(resultMaps, map[string]string{
					"name":        ds.nil2String(item["Name"]),
					"comment":     ds.nil2String(item["Comment"]),
					"engine":      ds.nil2String(item["Engine"]),
					"collation":   ds.nil2String(item["Collation"]),
					"data_length": ds.nil2String(item["Data_length"]),
					"create_time": ds.nil2String(item["Create_time"]),
					"update_time": ds.nil2String(item["Update_time"]),
				})
			}
		}
		return resultMaps, int(affectRows)
	}
	if o.Driver().Type() == orm.DRSqlite {
		sql := "SELECT name FROM  sqlite_master WHERE type ='table' AND name NOT LIKE 'sqlite_%';"
		affectRows, err := o.Raw(sql).Values(&maps)
		if affectRows > 0 && err == nil {
			for _, item := range maps {
				resultMaps = append(resultMaps, map[string]string{
					"name":        ds.nil2String(item["name"]),
					"comment":     ds.nil2String(item["table"]),
					"engine":      "",
					"collation":   "",
					"data_length": "",
					"create_time": "",
					"update_time": "",
				})
			}
			return resultMaps, int(affectRows)
		}
	}
	return nil, 0
}

// OptimizeTable
func (*databaseService) OptimizeTable(tableName string) bool {
	o := orm.NewOrm()
	if o.Driver().Type() == orm.DRMySQL {
		_, err := o.Raw("OPTIMIZE TABLE `" + tableName + "`").Exec()
		return err == nil
	}
	if o.Driver().Type() == orm.DRSqlite {
		_, err := o.Raw("VACUUM;").Exec()
		return err == nil
	}
	return false
}

// RepairTable
func (*databaseService) RepairTable(tableName string) bool {
	o := orm.NewOrm()
	if o.Driver().Type() == orm.DRMySQL {
		_, err := o.Raw("REPAIR TABLE `" + tableName + "`").Exec()
		return err == nil
	}
	if o.Driver().Type() == orm.DRSqlite {
		_, err := o.Raw("VACUUM;").Exec()
		return err == nil
	}
	return false
}

// GetFullColumnsFromTable
func (ds *databaseService) GetFullColumnsFromTable(tableName string) []map[string]string {
	var maps []orm.Params
	var resultMaps []map[string]string
	o := orm.NewOrm()
	if o.Driver().Type() == orm.DRMySQL {
		affectRows, err := o.Raw("SHOW FULL COLUMNS FROM `" + tableName + "`").Values(&maps)
		if affectRows > 0 && err == nil {
			for _, item := range maps {
				resultMaps = append(resultMaps, map[string]string{
					"name":       ds.nil2String(item["Field"]),
					"type":       ds.nil2String(item["Type"]),
					"collation":  ds.nil2String(item["Collation"]),
					"null":       ds.nil2String(item["Null"]),
					"key":        ds.nil2String(item["Key"]),
					"default":    ds.nil2String(item["Default"]),
					"extra":      ds.nil2String(item["Extra"]),
					"privileges": ds.nil2String(item["Privileges"]),
					"comment":    ds.nil2String(item["Comment"]),
				})
			}
		}
	}
	if o.Driver().Type() == orm.DRSqlite {
		affectRows, err := o.Raw("PRAGMA table_info('" + tableName + "')").Values(&maps)
		if affectRows > 0 && err == nil {
			for _, item := range maps {
				resultMaps = append(resultMaps, map[string]string{
					"name":       ds.nil2String(item["name"]),
					"type":       ds.nil2String(item["type"]),
					"collation":  ds.nil2String(item["Collation"]),
					"null":       ds.nil2String(item["notnull"]),
					"key":        ds.nil2String(item["pk"]),
					"default":    ds.nil2String(item["dflt_value"]),
					"extra":      "",
					"privileges": "",
					"comment":    "",
				})
			}
		}
	}
	return resultMaps
}
func (ds *databaseService) ExportDatabaseToJSON() []error {
	err := database.ExportData()
	//err = database.ImportData("20220727092558")
	return err
}
func (ds *databaseService) ImportDatabaseFromJSON(folder string,module string) error {
	err := database.ImportData(folder,module)
	return err
}
func (ds *databaseService) AvailableImports() []string {
	list := database.AvailableImports()
	return list
}

// nil2String interface
func (*databaseService) nil2String(val interface{}) string {
	if val == nil {
		return ""
	}
	return val.(string)
}
