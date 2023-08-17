package master

import (
	"quince/initialize/module"
	"quince/modules/admin/models"
	_ "quince/modules/master/routers"
	_ "quince/modules/master/cron"

	"github.com/beego/beego/v2/client/orm"

	"github.com/beego/beego/v2/core/logs"
)

type master struct {
	module.IModule
}

func init() {
	a := *new(master)
	module.Register("master", a)
}

func (m master) Init(tx orm.TxOrmer) error {
	logs.Info("Module master initialisation")
	var temp models.AdminMenu
	var parentId int64
	var err error

	//cnt, err := tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", "master").Count()
	//if cnt == 0 {
	m1 := models.AdminMenu{Name: "master.title", Url: "master/index", Icon: "fa-university", Module: "master", IsShow: 1, Status: 1, SortId: 30, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m1.Module).Filter("url", m1.Url).One(&temp, "id")
	parentId = temp.GetID()
	if err == orm.ErrNoRows { //insert only. admin can modify menu settings
		parentId, _ = tx.Insert(&m1)
	} else if err != nil {
		logs.Error(err.Error())
	}

	//Companies
	m2 := models.AdminMenu{Name: "master.menu.companies", Url: "master/company/manager", Icon: "fa-industry", ParentId: parentId, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m2.Module).Filter("url", m2.Url).One(&temp, "id")
	menuCompany := temp.GetID()
	if err == orm.ErrNoRows {
		menuCompany, _ = tx.Insert(&m2)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m21 := models.AdminMenu{Name: "master.menu.companies", Url: "master/company/index", Icon: "fa-industry", ParentId: menuCompany, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m21.Module).Filter("url", m21.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m21)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m22 := models.AdminMenu{Name: "master.menu.group", Url: "master/group/index", Icon: "fa-sitemap", ParentId: menuCompany, Module: "master", IsShow: 1, Status: 1, SortId: 999, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m22.Module).Filter("url", m22.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m22)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m23 := models.AdminMenu{Name: "master.menu.department", Url: "master/department/index", Icon: "fa-building-o", ParentId: menuCompany, Module: "master", IsShow: 1, Status: 1, SortId: 999, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m23.Module).Filter("url", m23.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m23)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m24 := models.AdminMenu{Name: "master.menu.smtps", Url: "master/smtp/index", Icon: "fa-server", ParentId: menuCompany, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m24.Module).Filter("url", m24.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m24)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m25 := models.AdminMenu{Name: "master.menu.bankaccount", Url: "master/bankaccount/index", Icon: "fa-briefcase", ParentId: menuCompany, Module: "master", IsShow: 1, Status: 1, SortId: 999, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m25.Module).Filter("url", m25.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m25)
	} else if err != nil {
		logs.Error(err.Error())
	}

	//Basic
	m3 := models.AdminMenu{Name: "master.menu.basic", Url: "master/basic/manager", Icon: "fa-info", ParentId: parentId, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m3.Module).Filter("url", m3.Url).One(&temp, "id")
	menuBasic := temp.GetID()
	if err == orm.ErrNoRows {
		menuBasic, _ = tx.Insert(&m3)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m31 := models.AdminMenu{Name: "master.menu.countries", Url: "master/country/index", Icon: "fa-flag-checkered", ParentId: menuBasic, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m31.Module).Filter("url", m31.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m31)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m32 := models.AdminMenu{Name: "master.menu.counties", Url: "master/county/index", Icon: "fa-shield", ParentId: menuBasic, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m32.Module).Filter("url", m32.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m32)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m33 := models.AdminMenu{Name: "master.menu.um", Url: "master/um/index", Icon: "fa-balance-scale", ParentId: menuBasic, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m33.Module).Filter("url", m33.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m33)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m34 := models.AdminMenu{Name: "master.menu.deliveryterm", Url: "master/deliveryterm/index", Icon: "fa-handshake-o", ParentId: menuBasic, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m34.Module).Filter("url", m34.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m34)
	} else if err != nil {
		logs.Error(err.Error())
	}

	m35 := models.AdminMenu{Name: "master.menu.forwarder", Url: "master/forwarder/index", Icon: "fa-truck", ParentId: menuBasic, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m35.Module).Filter("url", m35.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m35)
	} else if err != nil {
		logs.Error(err.Error())
	}
	//Taxes & money
	m4 := models.AdminMenu{Name: "master.menu.money", Url: "master/money/manager", Icon: "fa-money", ParentId: parentId, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m4.Module).Filter("url", m4.Url).One(&temp, "id")
	menuMoney := temp.GetID()
	if err == orm.ErrNoRows {
		menuMoney, _ = tx.Insert(&m4)
	} else if err != nil {
		logs.Error(err.Error())
	}
	m41 := models.AdminMenu{Name: "master.menu.currencies", Url: "master/currency/index", Icon: "fa-stack-exchange", ParentId: menuMoney, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m41.Module).Filter("url", m41.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m41)
	} else if err != nil {
		logs.Error(err.Error())
	}
	m42 := models.AdminMenu{Name: "master.menu.exchange_rate", Url: "master/exchange_rate/index", Icon: "fa-eur", ParentId: menuMoney, Module: "master", IsShow: 1, Status: 1, SortId: 1000, LogMethod: "OFF"}
	err = tx.QueryTable(new(models.AdminMenu).TableName()).Filter("module", m42.Module).Filter("url", m42.Url).One(&temp, "id")
	if err == orm.ErrNoRows {
		tx.Insert(&m42)
	} else if err != nil {
		logs.Error(err.Error())
	}
	if err == orm.ErrNoRows {
		return nil
	} else {
		return err
	}
}
