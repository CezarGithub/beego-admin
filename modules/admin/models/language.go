package models

type Language struct {
	Code        string `json:"code" i18n:"admin.settingsgroup.code"`
	Description string `json:"description" i18n:"admin.settingsgroup.description"`
}
