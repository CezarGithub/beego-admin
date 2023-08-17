package exceloffice

import (
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/beego/beego/v2/server/web/context"
)

// ExportData
func ExportData(head []string, body [][]string, name string, version string, title string, responseWriter *context.Response) {
	if name == "" {
		name = time.Now().Format("2006-01-02-15-04-05")
	}

	if title == "" {
		title = "Export"
	}

	if version == "" {
		version = "2007"
	}

	charIndex := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	//Handle more than 26 columns
	a := "A"

	for _, v := range charIndex {
		charIndex = append(charIndex, a+v)
	}

	//Create excel
	f := excelize.NewFile()
	f.SetActiveSheet(0)

	//Excel table header
	for key, val := range head {
		f.SetCellValue("Sheet1", charIndex[key]+"1", val)
	}

	//Excel table body part
	for key, val := range body {
		row := key + 2
		col := 0
		for _, v := range val {
			f.SetCellValue("Sheet1", charIndex[col]+strconv.Itoa(row), v)
			col++
		}
	}

	//Finally set the custom title of Sheet1
	f.SetSheetName("Sheet1", title)

	//Version difference information
	versionOpt := map[string]map[string]string{
		"2007": {
			"mime":       "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
			"ext":        ".xlsx",
			"write_type": "Xlsx",
		},
		"2003": {
			"mime":       "application/vnd.ms-excel",
			"ext":        ".xls",
			"write_type": "Xls",
		},
		"pdf": {
			"mime":       "application/pdf",
			"ext":        ".pdf",
			"write_type": "PDF",
		},
		"ods": {
			"mime":       "application/vnd.oasis.opendocument.spreadsheet",
			"ext":        ".ods",
			"write_type": "OpenDocument",
		},
	}

	responseWriter.Header().Set("Content-Type", versionOpt[version]["mime"])
	responseWriter.Header().Set("Content-Disposition", "attachment;filename=\""+name+versionOpt[version]["ext"]+"\"")
	responseWriter.Header().Set("Cache-Control", "max-age=0")

	if _, err := f.WriteTo(responseWriter); err != nil {
		logs.Warning("export data err:", err.Error())
	}
	return
}
