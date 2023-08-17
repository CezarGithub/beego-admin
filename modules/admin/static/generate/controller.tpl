package controllers


import (
	"errors"
	"quince/internal/toolbar"
	"quince/modules/admin/controllers"
	"quince/modules/{{.Model.Module}}/form"
	"quince/modules/{{.Model.Module}}/models"
	"quince/modules/{{.Model.Module}}/services"
	"quince/modules/{{.Model.Module}}/transactions"
	tx "quince/modules/{{.Model.Module}}/transactions/{{.Model.Name_lcase}}"
	"github.com/beego/beego/v2/server/web"

)

// {{.Model.Name}}Controller struct
type {{.Model.Name}}Controller struct {
	controllers.BaseController
}

// NestPrepare - Init controller
func (mc *{{.Model.Name}}Controller) NestPrepare() {}

func (mc *{{.Model.Name}}Controller) Index() {
	{{.Model.Name_lcase}}Service := services.New{{.Model.Name}}Service()
	data := {{.Model.Name_lcase}}Service.GetAll(*mc.GetQueryParams())
	mc.Data["data"] = data
	add := toolbar.Html("{{.Model.Name_lcase}}.add").Add(web.URLFor("{{.Model.Name}}Controller.Add"))
	delete := toolbar.Ajax("{{.Model.Name_lcase}}.delete").DeleteSelectedRows(web.URLFor("{{.Model.Name}}Controller.Delete"))
	export := toolbar.Ajax("{{.Model.Name_lcase}}.export").Export(web.URLFor("{{.Model.Name}}Controller.Export"))
	edit := toolbar.Html("{{.Model.Name_lcase}}.edit").Edit(web.URLFor("{{.Model.Name}}Controller.Edit"))
	mc.AddButtons(edit,add, delete, export)
	mc.AddToolbars(t, e)
	//mc.ViewPath = "modules"
	//mc.Layout = "admin/views/public/base.html"
	mc.TplName = "master/views/{{.Model.Name_lcase}}/index.html"
}
func (mc *{{.Model.Name}}Controller) Add() {
	//item := models.{{.Model.Name}}{}
	{{- range $item := .Model.Fields}}
		{{- if eq $item.Type "ptr" }}
	//	item.{{$item.Name}} = new(models.{{$item.Name}})
		{{- end}}
	{{- end}}
	item := models.New{{.Model.Name}}()
	mc.show(&item)
}
func (mc *{{.Model.Name}}Controller) Edit() {
	id, _ := mc.GetInt64("id", -1)
	{{.Model.Name_lcase}}Service := services.New{{.Model.Name}}Service()
	item := {{.Model.Name_lcase}}Service.Get{{.Model.Name}}ById(id)
	if item == nil {
		mc.ResponseErrorWithMessage(errors.New("error.info_not_found"), mc.Ctx)
	}
	mc.show(item)
}
func (mc *{{.Model.Name}}Controller) show(item *models.{{.Model.Name}}) {
	mc.Data["data"] = item
{{- range $item := .Model.Fields}}
  {{- if eq $item.Type "ptr" }}
	//for SearchBox widget
	//mc.DataSearchBox["searchBox{{$item.Name}}"] = &models.{{$item.Name}}{}
	//OR for DropDown
	//service{{$item.Name}}:=services.New{{$item.Name}}Service()
	//{{$item.Name_lcase}}_list := countyService.GetAll(*mc.GetQueryParams())
	//mc.Data["{{$item.Name_lcase}}_list"] = {{$item.Name_lcase}}_list
  {{- end}}
{{- end}}
	save := toolbar.Ajax("{{.Model.Name_lcase}}.submit").Submit("{{.Model.Name_lcase}}Form", web.URLFor("{{.Model.Name}}Controller.Update"))
	mc.AddButtons(save)
	mc.TplName = "master/views/{{.Model.Name_lcase}}/edit.html"
}
func (mc *{{.Model.Name}}Controller) Search() {
	{{.Model.Name_lcase}}Service := services.New{{.Model.Name}}Service()
	data := {{.Model.Name_lcase}}Service.GetAll(*mc.GetQueryParams())
	mc.Data["json"] = &data
	mc.ServeJSON()
}
// Delete
func (mc *{{.Model.Name}}Controller) Delete() {
	idArr := mc.GetSelectedIDs()
	t1 := tx.{{.Model.Name}}Delete(idArr)
	mc.Transaction.Add(t1)
	err := mc.Transaction.Execute()

	if err == nil {
		uc.ResponseSuccess(mc.Ctx)
	} else {
		uc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}
func (mc *{{.Model.Name}}Controller) Update() {
	{{.Model.Name_lcase}}Form := form.{{.Model.Name}}Form{}
	if err := mc.ParseForm(&{{.Model.Name_lcase}}Form); err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}

	c, err := {{.Model.Name_lcase}}Form.Validate()
	if err != nil {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
	t1 := transactions.{{.Model.Name}}Update(c)
	mc.Transaction.Add(t1)
	err = mc.Transaction.Execute()

	if err == nil {
		mc.ResponseSuccess(mc.Ctx)
	} else {
		mc.ResponseErrorWithMessage(err, mc.Ctx)
	}
}

