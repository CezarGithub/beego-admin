package services

import (
	"net/url"
	im "quince/internal/models"
	"quince/utils"
	"quince/utils/page"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// BaseService struct
type BaseService struct {
	//Searchable fields- if empty ,default values from model.SearchField will be used
	SearchField []string
	//Specific where condition defined in model ; receive a pointer to model
	//if nil generic search will be used
	im.IModel
	//Can be used as a field for time range query
	TimeField []string
	//Pagination
	Pagination page.Pagination
	//Service Data Query Language support
	//for transaction support set Tx to *orm.TxOrmer
	DQL orm.DQL
	//transaction
	Tx orm.TxOrmer
}
//Return a orm.DQL object
func (bs *BaseService) DataQuery() orm.DQL {
	if bs.Tx == nil {
		return orm.NewOrm()
	} else {
		return bs.Tx
	}
}
//Return a orm.DML object
func (bs *BaseService) DataManipulation() orm.DML {
	if bs.Tx == nil {
		return orm.NewOrm()
	} else {
		return bs.Tx
	}
}

// Paginate
func (bs *BaseService) Paginate(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {
	var pagination page.Pagination
	qs := pagination.Paginate(seter, listRows, parameters)
	bs.Pagination = pagination
	return qs
}

// ScopeWhere
func (bs *BaseService) ScopeWhere(seter orm.QuerySeter, parameters url.Values) orm.QuerySeter {

	// if bs.IModel == nil { // no condition at all ,leave
	// 	return seter
	// }
	cond := orm.NewCondition()
	if len(bs.SearchField) == 0 { //if none custom fields..
		bs.SearchField = append(bs.SearchField, bs.IModel.SearchField()...)
	}
	if len(bs.TimeField) == 0 { //if none custom fields..
		bs.TimeField = append(bs.TimeField, bs.IModel.TimeField()...)
	}

	if bs.IModel.WhereCondition() != nil && !bs.IModel.WhereCondition().IsEmpty() { //there is specific filter condition
		cond = bs.IModel.WhereCondition()
	} else { //no specific filter - go generic
		//Keyword like search
		keywords := parameters.Get("_keywords")

		if keywords != "" && len(bs.SearchField) > 0 {
			for _, v := range bs.SearchField {
				cond = cond.Or(v+"__icontains", keywords)
			}
		}

		//Field condition query with named parameters - faster than previous
		if len(bs.SearchField) > 0 && len(parameters) > 0 {
			for k, v := range parameters {
				if v[0] != "" && utils.InArrayForString(bs.SearchField, k) {
					cond = cond.And(k+"__icontains", v[0])
				}
			}
		}

		//Time range query
		if len(bs.TimeField) > 0 && len(parameters) > 0 {
			for key, value := range parameters {
				if value[0] != "" && utils.InArrayForString(bs.TimeField, key) {
					timeRange := strings.Split(value[0], " - ")
					startTimeStr := timeRange[0]
					endTimeStr := timeRange[1]

					loc, _ := time.LoadLocation("Local")
					startTime, err := time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, loc)

					if err == nil {
						unixStartTime := startTime.Unix()
						if len(endTimeStr) == 10 {
							endTimeStr += "23:59:59"
						}

						endTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, loc)
						if err == nil {
							unixEndTime := endTime.Unix()
							cond = cond.And(key+"__gte", unixStartTime).And(key+"__lte", unixEndTime)
						}
					}
				}
			}
		}
	}
	//Assemble the conditional statement into the main statement
	seter = seter.SetCond(cond)

	//Sort
	order := parameters.Get("_order")
	by := parameters.Get("_by")
	if order == "" {
		order = "id"
	}

	if by == "" {
		by = "" //ASC or '-' DESC
	} else {
		if by == "asc" {
			by = ""
		} else {
			by = "-"
		}
	}

	//Sort
	seter = seter.OrderBy(by + order)

	return seter
}

// PaginateAndScopeWhere Paging and query merge, mostly used for home page list display and search
// generic search based on search fields from models
func (bs *BaseService) PaginateAndScopeWhere(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {
	return bs.Paginate(bs.ScopeWhere(seter, parameters), listRows, parameters)
}

// GetAllAndScopeWhere - all records and query merge
// for search with JSON returns
// generic search based on search fields from models
func (bs *BaseService) GetAllAndScopeWhere(seter orm.QuerySeter, parameters url.Values) orm.QuerySeter {
	return bs.Paginate(bs.ScopeWhere(seter, parameters), -1, parameters)
}
