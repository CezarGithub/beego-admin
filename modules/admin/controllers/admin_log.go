package controllers

import (
	"encoding/json"
	"fmt"
	"quince/global"
	"quince/modules/admin/services"
	"quince/utils/encrypter"
)

// AdminLogController struct.
type AdminLogController struct {
	BaseController
}

// Index index.
func (alc *AdminLogController) Index() {

	adminLogService := services.NewAdminLogService()
	adminUserService := services.NewAdminUserService()
	data, pagination := adminLogService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	alc.Data["admin_user_list"] = adminUserService.GetAllAdminUser()
	alc.Data["data"] = data
	alc.Data["paginate"] = pagination
	alc.TplName = "admin/views/admin_log/index.html"
}

// Index Menu home.
func (alc *AdminLogController) View() {

	logId, _ := alc.GetInt("id", -1)
	adminLogService := services.NewAdminLogService()
	adminLogData := adminLogService.GetAdminLogDataById(logId)
	cryptData := encrypter.Decrypt(adminLogData.Data, []byte(global.BA_CONFIG.Other.LogAesKey))
	var data map[string]interface{}
	err := json.Unmarshal([]byte(cryptData), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}
	var html string
	for key, element := range data {
		// fmt.Println("Key:", key, "=>", "Element:", element)
		html += fmt.Sprintf("<tr><td>%s</td><td>%v</td></tr>", key, element)
	}
	adminLogData.Data = html
	alc.Data["data"] = adminLogData
	alc.TplName = "admin/views/admin_log/view.html"
	//routesService:= services.NewAdminLogService()
	//data, _ := routesService.GetPaginatedData(admin["per_page"].(int), gQueryParams)
	//log.Printf("Params %v \n", gQueryParams)

	//amc.Data["data"] = data

}
