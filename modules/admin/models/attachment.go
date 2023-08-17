package models

import (
	"encoding/json"
	"math"
	"quince/global"
	"quince/initialize/database"
	"quince/internal/validation"
	"quince/utils"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// Attachment struct
type Attachment struct {
	LoginUser    *LoginUser `orm:"rel(fk)"`
	User         *User      `orm:"rel(fk)"`
	OriginalName string     `orm:"column(original_name);size(200)" description:"Original File Name" i18n:"admin.attachment.original_name"`
	SaveName     string     `orm:"column(save_name);size(200)" description:"Save file name" i18n:"admin.attachment.save_name"`
	SavePath     string     `orm:"column(save_path);size(255)" description:"System full path" i18n:"admin.attachment.save_path"`
	Url          string     `orm:"column(url);size(255)" description:"Image access path" i18n:"admin.attachment.url"`
	Extension    string     `orm:"column(extension);size(100)" description:"suffix" i18n:"admin.attachment.extension"`
	Mime         string     `orm:"column(mime);size(100)" description:"type" i18n:"admin.attachment.mime"`
	Size         int64      `orm:"column(size);size(20);default(0)" description:"size" i18n:"admin.attachment.size"`
	Md5          string     `orm:"column(md5);size(32);default(\"\")" description:"MD5" i18n:"admin.attachment.md5"`
	Sha1         string     `orm:"column(sha1);size(40);default(\"\")" description:"SHA1" i18n:"admin.attachment.sha1"`
	Base
}

// TableName
func (*Attachment) TableName() string {
	return "admin_attachment"
}

func (cs *Attachment) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.SaveName != "" {
			cond = cond.And("save_name__icontains", cs.SaveName)
		}
		if cs.SavePath != "" {
			cond = cond.And("save_path__icontains", cs.SavePath)
		}
		if cs.OriginalName != "" {
			cond = cond.And("original_name__icontains", cs.OriginalName)
		}
	}
	return cond
}

// SearchField
func (*Attachment) SearchField() []string {
	return []string{"save_name", "original_name", "save_path"}
}

// TimeField
func (*Attachment) TimeField() []string {
	return []string{}
}

// init
func init() {
	database.RegisterModel("admin", new(Attachment))
}

func (c *Attachment) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.User.Id, validation.Required, validation.Min(0)))
	rules = append(rules, validation.Field(&c.OriginalName, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.SaveName, validation.Required, validation.Length(1, 300)))
	err := validation.ValidateStruct(c, rules...)
	return err
}

// FileType
func (*Attachment) FileType() map[string][]string {
	return map[string][]string{
		"Img":    {"jpg", "bmp", "png", "jpeg", "gif", "svg"},
		"Doc":    {"txt", "doc", "docx", "xls", "xlsx", "pdf"},
		"Arhive": {"rar", "zip", "7z", "tar"},
		"Audio":  {"mp3", "ogg", "flac", "wma", "ape"},
		"Video":  {"mp4", "wmv", "avi", "rmvb", "mov", "mpg"},
	}
}

// FileThumb 属性定义
func (*Attachment) FileThumb() map[string][]string {
	return map[string][]string{
		"picture":      {"jpg", "bmp", "png", "jpeg", "gif", "svg"},
		"txt.svg":      {"txt", "pdf"},
		"pdf.svg":      {"pdf"},
		"word.svg":     {"doc", "docx"},
		"excel.svg":    {"xls", "xlsx"},
		"archives.svg": {"rar", "zip", "7z", "tar"},
		"audio.svg":    {"mp3", "ogg", "flac", "wma", "ape"},
		"video.svg":    {"mp4", "wmv", "avi", "rmvb", "mov", "mpg"},
	}
}

// GetSize
func (attachment *Attachment) GetSize() string {
	size := float64(attachment.Size)
	units := []string{" B", " KB", " MB", " GB", " TB"}
	var i int
	for i = 0; size >= 1024 && i < 4; i++ {
		size /= 1024
	}
	return strconv.FormatFloat(math.Round(size), 'f', -1, 64) + units[i]
}

// GetFileType
func (attachment *Attachment) GetFileType() string {
	typeName := "other"
	extension := attachment.Extension
	for name, arr := range attachment.FileType() {
		if utils.InArrayForString(arr, extension) {
			typeName = name
			break
		}
	}
	return typeName
}

// GetThumbnail
func (attachment *Attachment) GetThumbnail() string {
	thumbnail := global.BA_CONFIG.Attachment.ThumbPath + "unknown.svg"
	extension := attachment.Extension
	thumbPath := global.BA_CONFIG.Attachment.ThumbPath

	fileThumb := map[string][]string{
		"picture":                  {"jpg", "bmp", "png", "jpeg", "gif", "svg"},
		thumbPath + "txt.svg":      {"txt", "pdf"},
		thumbPath + "pdf.svg":      {"pdf"},
		thumbPath + "word.svg":     {"doc", "docx"},
		thumbPath + "excel.svg":    {"xls", "xlsx"},
		thumbPath + "archives.svg": {"rar", "zip", "7z", "tar"},
		thumbPath + "audio.svg":    {"mp3", "ogg", "flac", "wma", "ape"},
		thumbPath + "video.svg":    {"mp4", "wmv", "avi", "rmvb", "mov", "mpg"},
	}

	for name, arr := range fileThumb {
		if utils.InArrayForString(arr, extension) {
			if name == "picture" {
				thumbnail = attachment.Url
			} else {
				thumbnail = name
			}
			break
		}
	}

	return thumbnail
}

func (t *Attachment) Export() []byte {
	var items []*Attachment
	_, err := orm.NewOrm().QueryTable(t.TableName()).All(&items)
	data, _ := json.Marshal(items)
	if err != nil {
		logs.Error(err.Error())
	}
	return data
}
func (t *Attachment) Import(tx orm.TxOrmer, data []byte) error {
	var list []*Attachment
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		return err
	}
	tx.QueryTable(t.TableName()).Filter("id__gt", 0).Delete()
	for _, item := range list {
		_, err := tx.Insert(item)
		if err != nil {
			return err
		}
	}
	return nil
}
