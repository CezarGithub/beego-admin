package controllers

import (
	"errors"
	"quince/initialize/database"
	"quince/initialize/module"
	"quince/modules/admin/services"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// DatabaseController struct
type DatabaseController struct {
	BaseController
}

// Table show
func (dc *DatabaseController) Table() {
	var list []map[string]string
	modules:=module.GetModulesNames()
	reqModule := dc.GetString("_module")
	databaseService := services.NewDatabaseService()
	data, affectRows := databaseService.GetTableStatus()
	models := database.RegisteredModels()

	for k, v := range models {
		m := strings.Split(k, ".")
		module := m[0]
		table:=v.TableName()
		for _, t := range data {
			if reqModule == module {
				if(t["name"]==table){
					list =append(list,t) 
				}
			}
		}
	}
	dc.Data["data"] = list
	dc.Data["total"] = affectRows
	dc.Data["module"] = reqModule
	dc.Data["modules"] = modules

	dc.TplName = "admin/views/database/table.html"
}

// Optimize
func (dc *DatabaseController) Optimize() {
	name := dc.GetString("name")

	if name == "" {
		dc.ResponseErrorWithMessage(errors.New("admin.select_table"), dc.Ctx)
	}
	databaseService := services.NewDatabaseService()
	ok := databaseService.OptimizeTable(name)
	if ok {
		dc.ResponseSuccessWithMessage("ok.message", dc.Ctx)
	} else {
		dc.ResponseErrorWithMessage(errors.New("error.failed"), dc.Ctx)
	}
}

// Repair
func (dc *DatabaseController) Repair() {
	name := dc.GetString("name")

	if name == "" {
		dc.ResponseErrorWithMessage(errors.New("admin.select_table"), dc.Ctx)
	}
	databaseService := services.NewDatabaseService()
	ok := databaseService.OptimizeTable(name)
	if ok {
		dc.ResponseSuccessWithMessage("ok.message", dc.Ctx)
	} else {
		dc.ResponseErrorWithMessage(errors.New("error.failed"), dc.Ctx)
	}
}

// View
func (dc *DatabaseController) View() {
	name := dc.GetString("name")

	if name == "" {
		dc.ResponseErrorWithMessage(errors.New("admin.select_table"), dc.Ctx)
	}

	databaseService := services.NewDatabaseService()
	data := databaseService.GetFullColumnsFromTable(name)

	dc.Data["data"] = data
	dc.Layout = "" //no base layout
	dc.TplName = "admin/views/database/view.html"
}

// Export to JSON
func (dc *DatabaseController) ExportDatabaseToJSON() {
	databaseService := services.NewDatabaseService()
	err := databaseService.ExportDatabaseToJSON()
	if err == nil {
		dc.ResponseSuccess(dc.Ctx)
	} else {
		logs.Error("Export DB error %v", err)
		dc.ResponseError(dc.Ctx)
	}
}
func (dc *DatabaseController) ImportDatabaseFromJSON() {
	databaseService := services.NewDatabaseService()
	list := databaseService.AvailableImports()
	layout := "20060102150405"
	var recent time.Time
	var folder string
	for _, str := range list {
		t, _ := time.Parse(layout, str)
		if recent.Before(t) {
			recent = t
			folder = str
		}
	}
	err := dc.ImportJSON(folder,"")
	if err == nil {
		dc.ResponseSuccess(dc.Ctx)
	} else {
		dc.ResponseError(dc.Ctx)
	}
}
func (dc *DatabaseController) InitDatabaseFromJSON() {
	reqModule := dc.GetString("_module")
	folder := "init"
	err := dc.ImportJSON(folder,reqModule)
	if err == nil {
		dc.ResponseSuccess(dc.Ctx)
	} else {
		dc.ResponseError(dc.Ctx)
	}
}
func (dc *DatabaseController) ImportJSON(folder string,module string) error {
	databaseService := services.NewDatabaseService()
	err := databaseService.ImportDatabaseFromJSON(folder,module)
	return err
}
