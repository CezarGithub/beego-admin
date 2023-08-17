package error

import (
	"fmt"
	"net/http"
	"quince/internal/i18n"
	"strings"
	"text/template"

	"github.com/beego/beego/v2/core/logs"
)

var tr i18n.Locale

type Message struct {
	Text string
	Flag string
	Back string
}

func PageNotFound(rw http.ResponseWriter, r *http.Request) {
	msg := "error.page_not_found"
	loadError(rw, r, msg)
}
func ServerError(rw http.ResponseWriter, r *http.Request) {
	msg := "error.internal_error"
	loadError(rw, r, msg)
}
func AccessDenied(rw http.ResponseWriter, r *http.Request) {
	msg := "error.access_denied"
	loadError(rw, r, msg)
}

func loadError(rw http.ResponseWriter, r *http.Request, msg string) {
	//t, _ := template.ParseFiles(beego.ViewsPath + path)
	var items []Message
	t, _ := template.New("errortpl").Parse(tpl)
	//l, _ := web.AppConfig.String("lang::alpha4")
	//langs := strings.Split(l, "|")
	langs := i18n.ApplicationLanguages()
	for _, lang := range langs {
		tr.Lang = lang
		c := tr.Tr(msg)
		m := Message{}
		m.Text = c
		m.Back = tr.Tr("app.click_to_return")
		alpha := strings.Split(lang, "-")
		alpha2 := "xx"
		if len(alpha) > 1 {
			alpha2 = strings.ToLower(alpha[1])
		}
		m.Flag = fmt.Sprintf("/static/images/flags/32x32/%s.png", alpha2)
		items = append(items, m)
	}

	//tr.Lang = "ro-RO"
	//c := tr.Tr(msg)
	logs.Warn("Error page %s", items)
	data := make(map[string]interface{})
	data["error"] = items
	t.Execute(rw, data)
}
