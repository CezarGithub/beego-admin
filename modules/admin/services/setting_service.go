package services

import (
	"encoding/json"
	"quince/global"
	"quince/internal/i18n"
	"quince/modules/admin/models"

	"github.com/beego/beego/v2/client/orm"
)

// SettingService struct
type SettingService struct {
	BaseService
}

// Show
func (settingService *SettingService) Show(id int64, lang string) []*models.Setting {
	data := settingService.getDataBySettingGroupId(id)

	settingFormService := NewSettingFormService()

	for key, value := range data {
		//contentNew := ""
		//value.Contentjson
		var contents []*models.Content
		if value.Content == "" {
			continue
		}
		err := json.Unmarshal([]byte(value.Content), &contents)

		if err != nil {
			continue
		}

		var i18tr i18n.Locale
		i18tr.Lang = lang

		data[key].Name = i18tr.Tr("admin.settings." + value.Name)

		var contentNew []*models.Content
		for _, content := range contents {
			//content.Form = settingFormService.GetFieldForm(lang,content.Type, content.Name, content.Field, content.Content, content.Option)
			translation := i18tr.Tr("admin.settings." + content.Field)
			content.Form = settingFormService.GetFieldForm(lang, content.Type, translation, content.Field, content.Content, content.Option)
			contentNew = append(contentNew, content)
		}
		data[key].ContentStrut = contentNew
	}

	return data
}

// getDataBySettingGroupId
func (cs *SettingService) getDataBySettingGroupId(settingGroupId int64) []*models.Setting {
	var settings []*models.Setting
	_, err := cs.DataQuery().QueryTable(new(models.Setting)).Filter("setting_group_id", settingGroupId).All(&settings)
	if err != nil {
		return nil
	}
	return settings
}

// GetSettingInfoById
func (cs *SettingService) GetSettingInfoById(id int64) *models.Setting {
	b := models.Base{}
	b.Id = id
	setting := models.Setting{Base: b}
	cs.DataQuery().Read(&setting)
	return &setting
}

// UpdateSettingInfoToContent
func (cs *SettingService) UpdateSettingInfoToContent(id int64, content string) int {
	affectRow, err := cs.DataQuery().QueryTable(new(models.Setting)).Filter("id", id).Update(orm.Params{
		"content": content,
	})
	if err == nil {
		return int(affectRow)
	}
	return 0
}

// LoadOrUpdateGlobalBaseConfig
func (cs *SettingService) LoadOrUpdateGlobalBaseConfig(setting *models.Setting) bool {
	if setting == nil {
		return false
	}

	if setting.Code == "base" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "name":
				global.BA_CONFIG.Base.Name = content.Content
			case "short_name":
				global.BA_CONFIG.Base.ShortName = content.Content
			case "author":
				global.BA_CONFIG.Base.Author = content.Content
			case "version":
				global.BA_CONFIG.Base.Version = content.Content
			case "link":
				global.BA_CONFIG.Base.Link = content.Content
			}
		}
	} else if setting.Code == "login" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "token":
				global.BA_CONFIG.Login.Token = content.Content
			case "captcha":
				global.BA_CONFIG.Login.Captcha = content.Content
			case "background":
				global.BA_CONFIG.Login.Background = content.Content
			}
		}
	} else if setting.Code == "index" {
		for _, content := range setting.ContentStrut {
			switch content.Field {
			case "password_warning":
				global.BA_CONFIG.Base.PasswordWarning = content.Content
			case "show_notice":
				global.BA_CONFIG.Base.ShowNotice = content.Content
			case "notice_content":
				global.BA_CONFIG.Base.NoticeContent = content.Content
			}
		}
	} else {
		return false
	}

	return true
}
