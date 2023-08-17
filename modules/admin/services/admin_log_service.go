package services

import (
	"encoding/json"
	"quince/global"
	"quince/modules/admin/models"
	"quince/utils/encrypter"
	"quince/utils/page"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"

	"net/url"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// AdminLogService struct
type adminLogService struct {
	BaseService
}

//AdminLogService - instantiate de IModel filter
func NewAdminLogService() adminLogService {
	var cs adminLogService
	c := models.AdminLog{}

	cs.IModel = &c
	return cs
}

// CreateAdminLog
func (*adminLogService) CreateAdminLog(loginUser *models.LoginUser, menu *models.AdminMenu, url string, ctx *context.Context) {
	var adminLog models.AdminLog

	if loginUser == nil {
		adminLog.AdminUserId = 0
	} else {
		adminLog.AdminUserId = loginUser.Id
	}
	adminLog.Name = menu.Name
	adminLog.LogMethod = menu.LogMethod
	adminLog.Url = url
	adminLog.LogIp = ctx.Input.IP()
	adminLog.CreateTime = time.Now()
	adminLog.UpdateTime = time.Now()

	o := orm.NewOrm()
	//Open transaction
	to, _ := o.Begin()

	adminLogID, err := to.Insert(&adminLog)
	if err != nil {
		to.Rollback()
		logs.Error(err)
		return
	}

	//adminLogData
	jsonData, _ := json.Marshal(ctx.Request.PostForm)
	cryptData := encrypter.Encrypt(jsonData, []byte(global.BA_CONFIG.Other.LogAesKey))
	var adminLogData models.AdminLogData
	adminLogData.AdminLogId = int(adminLogID)
	adminLogData.Data = cryptData
	_, err = to.Insert(&adminLogData)
	if err != nil {
		to.Rollback()
		logs.Error(err)
		return
	}
	to.Commit()
}

// LoginLog
func (*adminLogService) LoginLog(loginUserID int64, ctx *context.Context) {
	var adminLog models.AdminLog
	adminLog.AdminUserId = loginUserID
	adminLog.Name = "login"
	adminLog.Url = "admin/auth/login"
	adminLog.LogMethod = "POST"
	adminLog.LogIp = ctx.Input.IP()
	adminLog.CreateTime = time.Now()
	adminLog.UpdateTime = time.Now()

	o := orm.NewOrm()

	//
	to, _ := o.Begin()

	adminLogID, err := o.Insert(&adminLog)
	if err != nil {
		to.Rollback()
		logs.Error(err)
		return
	}

	//adminLogData
	jsonData, _ := json.Marshal(ctx.Request.PostForm)
	cryptData := encrypter.Encrypt(jsonData, []byte(global.BA_CONFIG.Other.LogAesKey))

	var adminLogData models.AdminLogData
	adminLogData.AdminLogId = int(adminLogID)
	adminLogData.Data = cryptData
	_, err = o.Insert(&adminLogData)
	if err != nil {
		to.Rollback()
		logs.Error(err)
		return
	}
	to.Commit()
}

// GetCount
func (cs *adminLogService) GetCount() int {
	count, err := cs.DataQuery().QueryTable(new(models.AdminLog)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetPaginateData adminuser
func (als *adminLogService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminLog, page.Pagination) {
	//Search and query field assignment
	als.SearchField = append(als.SearchField, new(models.AdminLog).SearchField()...)
	//als.WhereField = append(als.WhereField, new(models.AdminLog).WhereField()...)
	als.TimeField = append(als.TimeField, new(models.AdminLog).TimeField()...)

	var adminLog []*models.AdminLog
	o := als.DataQuery().QueryTable(new(models.AdminLog))
	_, err := als.PaginateAndScopeWhere(o, listRows, params).All(&adminLog)
	if err != nil {
		return nil, als.Pagination
	}
	return adminLog, als.Pagination
}

// 
func  (als *adminLogService) GetAdminLogDataById(adminLogId int ) *models.AdminLogData {
	var adminLogData models.AdminLogData
	err := als.DataQuery().QueryTable(new(models.AdminLogData)).Filter("admin_log_id", adminLogId).One(&adminLogData)
	if err == nil {
		return &adminLogData
	}
	return nil
}