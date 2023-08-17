package page

import (
	"fmt"
	"math"
	"net/url"
	"sort"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

// Pagination struct
type Pagination struct {
	//current page
	CurrentPage int
	//the last page
	LastPage int
	//Total number of data
	Total int
	//Number per page
	ListRows int
	//Is there a next page
	HasMore bool
	//Render tag
	BootStrapRenderLink string
	//Parameters in url
	Parameters url.Values
}

//Number per page
//config parameter
//  page:current page,
//  path:url path,
//  query:url extra parameters,
//  fragment:url anchor,
//  var_page:Paging variable,
//  list_rows:Number per page ; -1 will return all records !!!!!!
func (pagination *Pagination) Paginate(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {

	pagination.Parameters = parameters

	//current page
	var page int
	pageStr := pagination.Parameters.Get("page")
	if pageStr == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
	}

	if page < 1 {
		page = 1
	}

	pagination.CurrentPage = page
	pagination.ListRows = listRows
	total, err := seter.Count()
	if err != nil {
		pagination.Total = 0
	}

	pagination.Total = int(total)
	pagination.LastPage = int(math.Ceil((float64)(pagination.Total) / (float64)(pagination.ListRows)))
	pagination.HasMore = pagination.CurrentPage < pagination.LastPage

	//Put to the last execution, the first need to be assigned
	pagination.BootStrapRenderLink = pagination.render()

	if listRows == -1 {
		return seter
	}

	return seter.Limit(pagination.ListRows, (pagination.CurrentPage-1)*pagination.ListRows)
}

//Render pagination html
func (pagination *Pagination) render() string {
	if pagination.hasPages() {
		return fmt.Sprintf("<ul class=\"pagination pagination-sm no-margin pull-right\">%s %s %s</ul>",
			pagination.getPreviousButton(),
			pagination.getLinks(),
			pagination.getNextButton(),
		)
	} else {
		return ""
	}
}

//Back button
func (pagination *Pagination) getPreviousButton() string {
	text := "&laquo;"
	if pagination.CurrentPage <= 1 {
		return pagination.getDisabledTextWrapper(text)
	}
	url := pagination.url(pagination.CurrentPage - 1)
	return pagination.getPageLinkWrapper(url, text)
}

//Next button
func (pagination *Pagination) getNextButton() string {
	text := "&raquo;"
	if !pagination.HasMore {
		return pagination.getDisabledTextWrapper(text)
	}
	url := pagination.url(pagination.CurrentPage + 1)
	return pagination.getPageLinkWrapper(url, text)
}

//Page number button
func (pagination *Pagination) getLinks() string {
	block := map[string]map[int]string{
		"first":  nil,
		"slider": nil,
		"last":   nil,
	}

	side := 3
	window := side * 2

	if pagination.LastPage < window+6 {
		block["first"] = pagination.getUrlRange(1, pagination.LastPage)
	} else if pagination.CurrentPage <= window {
		block["first"] = pagination.getUrlRange(1, window+2)
		block["last"] = pagination.getUrlRange(pagination.LastPage-1, pagination.LastPage)
	} else if pagination.CurrentPage > (pagination.LastPage - window) {
		block["first"] = pagination.getUrlRange(1, 2)
		block["last"] = pagination.getUrlRange(pagination.LastPage-(window+2), pagination.LastPage)
	} else {
		block["first"] = pagination.getUrlRange(1, 2)
		block["slider"] = pagination.getUrlRange(pagination.CurrentPage-side, pagination.CurrentPage+side)
		block["last"] = pagination.getUrlRange(pagination.LastPage-1, pagination.LastPage)
	}

	html := ""
	if len(block["first"]) > 0 {
		html += pagination.getUrlLinks(block["first"])
	}

	if len(block["slider"]) > 0 {
		html += pagination.getDots()
		html += pagination.getUrlLinks(block["slider"])
	}

	if len(block["last"]) > 0 {
		html += pagination.getDots()
		html += pagination.getUrlLinks(block["last"])
	}

	return html

}

//Create a set of page links
func (pagination *Pagination) getUrlRange(start, end int) map[int]string {
	urls := map[int]string{}
	for page := start; page <= end; page++ {
		urls[page] = pagination.url(page)
	}
	return urls
}

//Get the link corresponding to the page number
func (pagination *Pagination) url(page int) string {
	parameters := pagination.Parameters

	urlValue := parameters.Get("queryParamUrl")
	parameters.Del("queryParamUrl")
	parameters.Del("_pjax")

	if len(parameters) > 0 {
		//Copy value
		parameters.Set("page", strconv.Itoa(page))
		urlStr := parameters.Encode()
		return urlValue + "?" + urlStr
	}
	return urlValue + "?page=" + strconv.Itoa(page)
}

//Generate a clickable button
func (pagination *Pagination) getAvailablePageWrapper(url string, page string) string {
	return `<li><a href="` + url + `">` + page + `</a></li>`
}

//Generate a disabled button
func (pagination *Pagination) getDisabledTextWrapper(text string) string {
	return `<li class="disabled"><span>` + text + `</span></li>`
}

//Generate an activated button
func (pagination *Pagination) getActivePageWrapper(text string) string {
	return `<li class="active"><span>` + text + `</span></li>`
}

//Generate ellipsis button
func (pagination *Pagination) getDots() string {
	return pagination.getDisabledTextWrapper("...")
}

//Batch generate page number button
func (pagination *Pagination) getUrlLinks(urls map[int]string) string {
	html := ""
	var sortKeys []int
	for page, _ := range urls {
		sortKeys = append(sortKeys, page)
	}
	sort.Ints(sortKeys)
	for _, page := range sortKeys {
		html += pagination.getPageLinkWrapper(urls[page], page)
	}
	return html
}

//Generate normal page number button
func (pagination *Pagination) getPageLinkWrapper(url string, page interface{}) string {
	pageInt, ok := page.(int)
	if ok {
		if pagination.CurrentPage == pageInt {
			return pagination.getActivePageWrapper(strconv.Itoa(pageInt))
		}
		return pagination.getAvailablePageWrapper(url, strconv.Itoa(pageInt))
	}
	pageStr := page.(string)
	return pagination.getAvailablePageWrapper(url, pageStr)
}

//Whether the data is paged enough
func (pagination *Pagination) hasPages() bool {
	return !(1 == pagination.CurrentPage && !pagination.HasMore)
}
