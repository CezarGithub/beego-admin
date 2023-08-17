package services

import (
	"quince/modules/admin/models"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type languageService struct {
}

// GetAllData
func (*languageService) GetApplicationLanguages() []*models.Language {
	var languages []*models.Language
	c, _ := web.AppConfig.String("lang::alpha4")
	n, _ := web.AppConfig.String("lang::names")
	codes := strings.Split(c, "|")
	names := strings.Split(n, "|")
	for i, code := range codes {
		l := models.Language{Code: code, Description: names[i]}
		languages = append(languages, &l)
	}
	return languages
}
func NewLanguageService() languageService {
	var cs languageService
	return cs
}
