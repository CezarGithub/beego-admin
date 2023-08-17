package models

import (
	"quince/initialize/database"
	"quince/internal/validation"

	"github.com/beego/beego/v2/client/orm"
)

// AdminMenu struct
type AdminRoute struct {
	Name          string `orm:"column(name);size(100)" description:"name"  i18n:"app.name"`
	Module        string `orm:"column(module);size(100)" description:"module"  i18n:"app.module"`
	Controller    string `orm:"column(controller);size(255)" description:"module"  i18n:"app.controller"`
	MappingMethod string `orm:"column(mapping_method);size(255)" description:"module"  i18n:"app.mapping_method"`
	Url           string `orm:"column(url);size(100);index;unique" description:"url"  i18n:"app.url"`
	Params        string `orm:"column(params);size(100);index" description:"parameters"  i18n:"app.parameter"`
	LogMethod     string `orm:"column(log_method);size(8);default(OFF)" description:"Logging method"  i18n:"app.logging_method"`
	Status        int8   `orm:"column(status);size(1);default(1)" description:"Enabled 0：NO 1：YES" i18n:"admin.user.status"`
	IsAPI         int8   `orm:"column(isAPI);size(1);default(0)" description:"Enabled 0：NO 1：YES" i18n:"app.route.isAPI"`
	Base
}

// TableName
func (*AdminRoute) TableName() string {
	return "admin_route"
}

// init
func init() {
	database.RegisterModel("admin", new(AdminRoute))
}

// GetLogMethod
func (*AdminRoute) GetLogMethod() []string {
	return []string{"OFF", "GET", "POST", "PUT", "DELETE"}
}

// SearchField
func (*AdminRoute) SearchField() []string {
	return []string{"name", "module"}
}

func (cs *AdminRoute) WhereCondition() *orm.Condition {
	cond := orm.NewCondition()
	if cs != nil {
		if cs.Id != 0 {
			cond = cond.And("id", cs.Id)
		}
		if cs.Name != "" {
			cond = cond.And("name__icontains", cs.Name)
		}
		if cs.Module != "" {
			cond = cond.And("module__icontains", cs.Module)
		}

	}
	return cond
}

// TimeField
func (*AdminRoute) TimeField() []string {
	return []string{}
}

func (c *AdminRoute) Validate() error {
	rules := []*validation.FieldRules{}
	rules = append(rules, validation.Field(&c.Name, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Module, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.Url, validation.Required, validation.Length(1, 255)))
	rules = append(rules, validation.Field(&c.LogMethod, validation.Required, validation.Length(1, 255)))
	//rules = append(rules, validation.Field(&c.SortId, validation.Required))
	err := validation.ValidateStruct(c, rules...)
	return err
}

// NO import export for routes. are generated by app.
func (t *AdminRoute) Export() []byte {

	return nil
}
func (t *AdminRoute) Import(tx orm.TxOrmer, data []byte) error {

	return nil
}