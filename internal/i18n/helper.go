package i18n

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

func LanguageFlag(language string) string {
	if len(language) < 2 {
		language = "xx"
	}
	src := fmt.Sprintf("/static/images/flags/32x32/%s.png", language[0:2])
	return src
}

// return value for Javascript - without quotes
func JavascriptLanguage(alpha4 string) template.JS {
	var t template.JS
	t = template.JS("en")
	langs := ApplicationLanguages()
	s, _ := web.AppConfig.String("lang::alpha2")
	short := strings.Split(s, "|")
	for i, lang := range langs {
		if lang == alpha4 {
			t = template.JS(short[i])
			return t
		}
	}
	return t //default
}

func ApplicationLanguages() []string {
	l, _ := web.AppConfig.String("lang::alpha4")
	langs := strings.Split(l, "|")
	if len(langs) == 0 {
		temp := []string{"en"}
		return temp
	}
	return langs
}
