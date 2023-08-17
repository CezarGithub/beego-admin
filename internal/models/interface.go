package models

import "github.com/beego/beego/v2/client/orm"

//IModel - updatable models
type IModel interface {
	GetID() int64
	SetUser(u int64)
	TableName() string
	Export() []byte
	Import(tx orm.TxOrmer, data []byte) error
	InitData(tx orm.TxOrmer) error //initial data if new database
	WhereCondition() *orm.Condition
	SearchField() []string
	TimeField() []string
}
