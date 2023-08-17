package component

import (
	"bytes"
	"encoding/json"

	//"fmt"
	"html/template"
	"quince/internal/i18n"
	utils "quince/utils/url"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

type IButton interface {
	GetUrl() string
	GetName() string
	GetData() string
	GetUrlParams() string
	IsDropDown() bool
	AddQueryParams(key string, value string)
	SetDeniedAccess()
	SetTitle(title string)
	SetIcon(icon string)
	Translate(lang string)
	HasVariableData() bool // has variable data in params ex.: {{$item}} etc..
}

// Button - common struct for all buttons in web interface
type Button struct {
	AuthPath           string            //authorized path
	Href               string            //href for a tag
	Params             string            //inline params without ? ex:. name=xyz&age=7
	DataUrl            string            //ajax calls if set
	Class              string            //button class
	Icon               string            //icon
	Title              string            //title & tooltip message
	DataConfirmTitle   string            //title on popum message
	DataConfirmContent string            //question on popup message
	DataId             string            //ID of current operation data for ajax calls
	DataData           template.HTMLAttr //if no DataId this field will contain AJAX call params key:value json array
	DataToggle         string            //checked , tooltip ,popover?
	DataType           string            //1-For direct access，2-for the modal window
	DataConfirm        string            //1-For direct access，2-prompt confirmation message
	DataForm           string            //form id - ex:submit button form id var form=$('#' + formName);
	DataMethod         string            //GET,POST etc..
	OnClick            string            //onclick javascript function
	Name               string            //button div name
	Enabled            bool
	params             map[string]string
}

func (b *Button) GetUrl() string {
	auth, err := utils.FullPath(b.AuthPath)
	if err != nil {
		return ""
	}
	return auth
}
func (b *Button) GetName() string {
	return b.Name
}
func (b *Button) GetData() string {
	return string(b.DataData)
}
func (b *Button) GetUrlParams() string {
	return b.Params
}
func (b *Button) SetDeniedAccess() {
	b.Href = "#"
	b.DataUrl = ""
	b.Enabled = false
}
func (b *Button) SetTitle(title string) {
	b.Title = title
}
func (b *Button) SetIcon(icon string) {
	b.Icon = icon
}

func (*Button) IsDropDown() bool {
	return false
}
func (b *Button) Translate(lang string) {
	b.Title = i18n.Tr(lang, b.Title)
	b.DataConfirmContent = i18n.Tr(lang, b.DataConfirmContent)
	b.DataConfirmTitle = i18n.Tr(lang, b.DataConfirmTitle)
}
func (b *Button) AddQueryParams(key string, value string) {
	if b.params == nil {
		b.params = make(map[string]string)
	}
	b.params[key] = value
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(&b.params); err != nil {
		logs.Info(err)
	}
	ts := buf.String()
	ts = strings.ReplaceAll(ts, "\n", "")
	b.DataData = template.HTMLAttr("data-data=" + ts) //this is only functional version - change carefully
	//'ZgotmplZ' is a special value that indicates that unsafe content reached a CSS or URL context at runtime

	//inline url params
	var p string
	i := 0
	var prefix string
	for k, v := range b.params {
		if i > 0 {
			prefix = "&"
		}
		p = prefix + p + k + "=" + v
	}
	b.Params = p
}
func (b *Button) HasVariableData() bool {

	//s := fmt.Sprintf("%v", b.DataData)
	for _, v := range b.params {
		if strings.Contains(v, "{{") && strings.Contains(v, "}}") {
			return true
		}
	}
	return false
}
