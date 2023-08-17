package controllers

import (
	"bufio"
	"encoding/base64"
	"io"
	"os"
	"quince/global"
	"quince/modules/admin/services"
	"quince/utils"
	"runtime"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/server/web"
)

// IndexController struct
type IndexController struct {
	BaseController
}

// PackageLib Loaded package information
type PackageLib struct {
	Name    string
	Version string
}

// Index
func (ic *IndexController) Index() {

	flash := web.NewFlash()
	flash.Success("Mesaj succes")
	flash.Store(&ic.Controller)

	ic.Log.Info("Index controller")
	ic.Data["login_user"] = loginUser

	//Default password modification detection
	ic.Data["password_danger"] = 0

	//Whether to display prompt information on the homepage
	ic.Data["show_notice"] = global.BA_CONFIG.Base.ShowNotice
	//Prompt content
	ic.Data["notice_content"] = global.BA_CONFIG.Base.NoticeContent

	//Default password modification detection
	loginUserPassword, _ := base64.StdEncoding.DecodeString(loginUser.Password)
	if global.BA_CONFIG.Base.PasswordWarning == "1" && utils.PasswordVerify("123456", string(loginUserPassword)) {
		ic.Data["password_danger"] = 1
	}

	//Number of background users
	adminUserService := services.NewAdminUserService()
	ic.Data["admin_user_count"] = adminUserService.GetCount()
	//Number of background roles
	adminRoleService := services.NewAdminRoleService()
	ic.Data["admin_role_count"] = adminRoleService.GetCount()
	//Number of background menus
	adminMenuService := services.NewAdminMenuService()
	ic.Data["admin_menu_count"] = adminMenuService.GetCount()
	//Number of background logs
	adminLogService := services.NewAdminLogService()
	ic.Data["admin_log_count"] = adminLogService.GetCount()
	//system message
	ic.Data["system_info"] = ic.getSystemInfo()

	ic.TplName = "admin/views/index/index.html"
}

// getSystemInfo
func (ic *IndexController) getSystemInfo() map[string]interface{} {
	systemInfo := make(map[string]interface{})
	//Server system
	systemInfo["server_os"] = runtime.GOOS
	//Go
	systemInfo["go_version"] = runtime.Version()
	//File upload default memory cache size
	systemInfo["upload_file_max_memory"] = int(web.BConfig.MaxMemory / 1024 / 1024)
	//beego
	systemInfo["beego_version"] = "N/A" // web..VERSION
	//Current background version
	systemInfo["admin_version"] = global.BA_CONFIG.Base.Version
	//mysql
	databaseService := services.NewDatabaseService()
	systemInfo["db_version"] = databaseService.GetMysqlVersion()
	//go
	systemInfo["timezone"] = time.UTC
	//current time
	systemInfo["date_time"] = time.Now().Format("2006-01-02 15:04:05")
	//IP
	systemInfo["user_ip"] = ic.Ctx.Input.IP()

	userAgent := ic.Ctx.Input.Header("user-agent")

	userOs := "Other"
	if strings.Contains(userAgent, "win") {
		userOs = "Windows"
	} else if strings.Contains(userAgent, "mac") {
		userOs = "MAC"
	} else if strings.Contains(userAgent, "linux") {
		userOs = "Linux"
	} else if strings.Contains(userAgent, "unix") {
		userOs = "Unix"
	} else if strings.Contains(userAgent, "bsd") {
		userOs = "BSD"
	} else if strings.Contains(userAgent, "iPad") || strings.Contains(userAgent, "iPhone") {
		userOs = "IOS"
	} else if strings.Contains(userAgent, "android") {
		userOs = "Android"
	}

	userBrowser := "Other"
	if strings.Contains(userAgent, "MSIE") {
		userBrowser = "MSIE"
	} else if strings.Contains(userAgent, "Firefox") {
		userBrowser = "Firefox"
	} else if strings.Contains(userAgent, "Chrome") {
		userBrowser = "Chrome"
	} else if strings.Contains(userAgent, "Safari") {
		userBrowser = "Safari"
	} else if strings.Contains(userAgent, "Opera") {
		userBrowser = "Opera"
	}

	//
	systemInfo["user_os"] = userOs
	//
	systemInfo["user_browser"] = userBrowser

	//go.mod
	var requireList []*PackageLib
	srcFile, err := os.Open("go.mod")
	if err != nil {
		logs.Error(err)
	} else {
		defer srcFile.Close()
		reader := bufio.NewReader(srcFile)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				continue
			}
			if string(line) != "" {
				strArr := strings.Split(strings.TrimSpace(string(line)), " ")
				lenStrArr := len(strArr)
				//require
				if strArr[0] == "require" && lenStrArr >= 3 {
					packageLib := PackageLib{
						Name:    strArr[1],
						Version: strArr[2],
					}
					requireList = append(requireList, &packageLib)
				} else {
					//require
					if lenStrArr >= 2 && strings.Contains(strArr[0], "/") {
						packageLib := PackageLib{
							Name:    strArr[0],
							Version: strArr[1],
						}
						requireList = append(requireList, &packageLib)
					}
				}
			}
		}
	}

	systemInfo["require_list"] = requireList

	return systemInfo
}
