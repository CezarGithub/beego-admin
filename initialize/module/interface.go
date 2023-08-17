package module
import "github.com/beego/beego/v2/client/orm"

type IModule interface {
	Init(tx orm.TxOrmer) error
}



