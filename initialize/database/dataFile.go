package database

import (
	"context"
	"errors"
	"io"
	"os"
	"path"
	"quince/global"
	"quince/internal/models"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type DataModel struct {
	Items map[string]models.IModel
	Path  string
}

func NewDataModelFile() (t *DataModel) {
	t = new(DataModel)
	t.Items = make(map[string]models.IModel)
	return t
}
func (t *DataModel) Register(module string, a models.IModel) {
	var key string
	o := reflect.TypeOf(a)
	if o.Kind() == reflect.Ptr {
		key = module + "." + o.Elem().Name()
	} else {
		key = module + "." + o.Name()
	}
	t.Items[key] = a
}

func (t *DataModel) Export() []error {
	var err []error
	n := time.Now()
	day := n.Format("20060102150405")
	folder := path.Join(global.BA_CONFIG.DataFile.Path, day)
	e := os.Remove(folder)
	if e != nil && !os.IsNotExist(e) {
		err = append(err, e)
	}
	e = os.Mkdir(folder, 0755)
	if e != nil {
		err = append(err, e)
	}
	for m, v := range t.Items {
		module := strings.Split(m, ".")
		fullPath := path.Join(folder, module[0])
		if stat, ero := os.Stat(fullPath); ero == nil && stat.IsDir() {
			//directory exists - don't do anything
		} else {
			e = os.Mkdir(fullPath, 0755) //create folders for each module,  one json file for each table
			if e != nil {
				err = append(err, e)
			}
		}
		data := v.Export()
		if data != nil {
			tableName := v.TableName() //.(models.IModel)
			file := path.Join(fullPath, tableName+".json")
			f, _ := os.Create(file)
			_, e = f.Write(data)
		}
		if e != nil {
			err = append(err, e)
		}
	}
	return err
}

func (t *DataModel) Import(folder string, module string) error {

	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		var er error
		for m, v := range t.Items {
			tableName := v.TableName() //(models.IModel).
			moduleName := strings.Split(m, ".")[0]
			if module == moduleName || module == "" {
				fullPath := path.Join(global.BA_CONFIG.DataFile.Path, folder, moduleName, tableName)
				jsonFile, e := os.Open(fullPath + ".json")
				if !errors.Is(e, os.ErrNotExist) {
					// handle the case where the file doesn't exist
					byteValue, _ := io.ReadAll(jsonFile)
					er = v.Import(txOrm, byteValue)
				} else {
					s := "[]" //send empty array to avoid JSON decode error
					er = v.Import(txOrm, []byte(s))
				} //otherwise send nil data to allow delete the table content
			}
		}
		return er
	})
	return err
}
func (t *DataModel) InitData() error {
	o := orm.NewOrm()
	var er error
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		for _, v := range t.Items {
			er = v.InitData(txOrm)
		}
		return er
	})
	return err
}
func (t *DataModel) AvailableImports() []string {
	var dir []string
	//files, err := ioutil.ReadDir(global.BA_CONFIG.DataFile.Path)
	files, err := os.ReadDir(global.BA_CONFIG.DataFile.Path)
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			dir = append(dir, fileInfo.Name())
		}
	}
	return dir
}
