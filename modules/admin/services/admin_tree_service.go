package services

import (
	//"quince/internal/i18n"
	//"quince/modules/admin/models"
	"quince/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

// AdminTreeService struct
type AdminTreeService struct {
	Ret   string
	Html  string
	Array map[int64]orm.Params
	Text  map[string]interface{}
}

var (
	icon  []string = []string{"│", "├", "└"}
	space string   = "&nbsp;&nbsp;"
)

// GetLeftMenu
// func (adminTreeService *AdminTreeService) GetLeftMenu(requestUrl string, user models.LoginUser, language string) string {
// 	menu := user.GetShowMenu()
// 	var maxLevel int64
// 	var currentID int64
// 	currentID = 1
// 	parentIds := []int64{0}

// 	for k, v := range menu {
// 		if v["Url"].(string) == requestUrl {
// 			parentIds = adminTreeService.getMenuParent(menu, v["Id"].(int64), []int64{})
// 			currentID = v["Id"].(int64)
// 		}
// 		menu[k]["Name"] = i18n.Tr(language, string(v["Name"].(string)))

// 	}

// 	if len(parentIds) == 0 {
// 		parentIds = []int64{0}
// 	}

// 	for k, v := range menu {
// 		menu[k]["Url"] = "/" + v["Url"].(string)
// 		tempLevel := adminTreeService.GetLevel(v["Id"].(int64), menu, 0)
// 		menu[k]["Level"] = tempLevel
// 		if maxLevel <= tempLevel {
// 			maxLevel = tempLevel
// 		}
// 	}

// 	adminTreeService.initTree(menu)

// 	textBaseOne := "<li class='treeview"
// 	textHover := " active"
// 	textBaseTwo := `'><a href='javascript:void(0);'>
// 	<i class='fa $icon'></i>
// 	<span>
// 	$name
// 	</span>
// 	<span class='pull-right-container'><i class='fa fa-angle-left pull-right'></i></span>
// 	</a><ul class='treeview-menu`
// 	textOpen := " menu-open"
// 	textBaseThree := "'>"

// 	textBaseFour := "<li"
// 	textHoverLi := " class='active'"
// 	textBaseFive := `>
// 	<a href='$url'>
// 	<i class='fa $icon'></i>
// 	<span>$name</span>
// 	</a>
// 	</li>`

// 	text0 := textBaseOne + textBaseTwo + textBaseThree
// 	text1 := textBaseOne + textHover + textBaseTwo + textOpen + textBaseThree
// 	text2 := "</ul></li>"
// 	textCurrent := textBaseFour + textHoverLi + textBaseFive
// 	textOther := textBaseFour + textBaseFive

// 	adminTreeService.Text = make(map[string]interface{})
// 	var i int64
// 	for i = 0; i <= maxLevel; i++ {
// 		adminTreeService.Text[(string)(i)] = []string{text0, text1, text2}
// 	}
// 	adminTreeService.Text["current"] = textCurrent
// 	adminTreeService.Text["other"] = textOther

// 	return adminTreeService.getAuthTree(0, currentID, parentIds)

// }

// getMenuParent
func (adminTreeService *AdminTreeService) getMenuParent(menu map[int64]orm.Params, myID int64, parentIds []int64) []int64 {
	for _, a := range menu {
		if a["Id"].(int64) == myID && int(a["ParentId"].(int64)) != 0 {
			parentIds = append(parentIds, a["ParentId"].(int64))
			parentIds = adminTreeService.getMenuParent(menu, a["ParentId"].(int64), parentIds)
		}
	}
	if len(parentIds) > 0 {
		return parentIds
	}
	return []int64{}
}

// GetLevel
func (adminTreeService *AdminTreeService) GetLevel(id int64, menu map[int64]orm.Params, i int64) int64 {
	var parentID int64
	v, ok := menu[id]["ParentId"].(int64)
	if ok {
		parentID = v
	} else {
		v1, ok := menu[id]["ParentId"].(int64)
		if ok {
			parentID = v1
		}
	}

	if (parentID == 0) || id == parentID {
		return i
	}
	i++
	return adminTreeService.GetLevel(parentID, menu, i)
}

// initTree
func (adminTreeService *AdminTreeService) initTree(menu map[int64]orm.Params) {
	adminTreeService.Array = make(map[int64]orm.Params)
	adminTreeService.Array = menu
	adminTreeService.Ret = ""
	adminTreeService.Html = ""
}

// getAuthTree
func (adminTreeService *AdminTreeService) getAuthTree(myId int64, currentID int64, parentIds []int64) string {
	nStr := ""
	child := adminTreeService.getChild(myId)
	if len(child) > 0 {
		menu := make(map[string]interface{})
		_ = menu
		//Take the smallest key to prevent random selection of for range, resulting in a different menu order each time
		var sortID int64
		sortID = 99999
		for k := range child {
			if k < sortID {
				sortID = k
			}
		}
		menu = child[sortID]

		//Get the current level of html
		var textHTMLInterface interface{}
		if adminTreeService.Text[(string)(menu["Level"].(int64))] != "" {
			//[]string
			textHTMLInterface = adminTreeService.Text[(string)(menu["Level"].(int64))]
		} else {
			//string
			textHTMLInterface = adminTreeService.Text["other"]
		}

		//Child sorting to prevent the menu position from being different every time
		var childKeys []int64
		for k := range child {
			childKeys = append(childKeys, k)
		}
		//sort.Ints(childKeys) - int only
		sort.Slice(childKeys, func(i, j int) bool { return childKeys[i] < childKeys[j] })
		for _, key := range childKeys {
			k := key
			v := child[key]

			if len(adminTreeService.getChild(k)) > 0 {
				textHTMLArr := textHTMLInterface.([]string)
				if utils.InArrayForInt(parentIds, k) {
					nStr = adminTreeService.strReplace(textHTMLArr[1], v)
					adminTreeService.Html += nStr
				} else {
					nStr = adminTreeService.strReplace(textHTMLArr[0], v)
					adminTreeService.Html += nStr
				}
				adminTreeService.getAuthTree(k, currentID, parentIds)
				nStr = adminTreeService.strReplace(textHTMLArr[2], v)
				adminTreeService.Html += nStr
			} else if k == currentID {
				a := adminTreeService.Text["current"].(string)
				nStr = adminTreeService.strReplace(a, v)
				adminTreeService.Html += nStr
			} else {
				a := adminTreeService.Text["other"].(string)
				nStr = adminTreeService.strReplace(a, v)
				adminTreeService.Html += nStr
			}
		}
	}
	return adminTreeService.Html
}

// getChild Get the child array
func (adminTreeService *AdminTreeService) getChild(pid int64) map[int64]map[string]interface{} {
	result := make(map[int64]map[string]interface{})
	for k, v := range adminTreeService.Array {
		parentID, ok := v["ParentId"].(int64)
		var parentIDInt int64
		if ok {
			parentIDInt = parentID
		} else {
			parentIDInt = v["ParentId"].(int64)
		}

		if parentIDInt == pid {
			result[k] = v
		}
	}
	return result
}

// strReplace
func (adminTreeService *AdminTreeService) strReplace(str string, m map[string]interface{}) string {
	str = strings.ReplaceAll(str, "$icon", m["Icon"].(string))
	str = strings.ReplaceAll(str, "$name", m["Name"].(string))
	str = strings.ReplaceAll(str, "$url", m["Url"].(string))
	return str
}

// GetTree
// myId Means to get all children under this ID
// str Generate the basic code of the tree structure, for example："<option value=\$id \$selected>\$spacer\$name</option>"
// sid The selected ID, for example, it needs to be used when making a tree drop-down box.
func (adminTreeService *AdminTreeService) GetTree(myId int64, str string, sid int64, adds string, strGroup string) string {
	number := 1
	child := adminTreeService.getChild(myId)

	if len(child) > 0 {
		total := len(child)

		//child排序
		var ids []int64
		for id := range child {
			ids = append(ids, id)
		}
		//sort.Ints(ids)
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

		for _, id := range ids {
			value := child[id]

			j := ""
			k := ""
			if number == total {
				j += icon[2]
			} else {
				j += icon[1]
				if adds != "" {
					k = icon[0]
				} else {
					k = ""
				}
			}

			spacer := ""
			if adds != "" {
				spacer = adds + j
			}
			selected := ""
			if id == sid {
				selected = "selected"
			}
			nStr := ""
			parentIDInt, ok := value["ParentId"].(int)
			if !ok {
				parentIDInt = int(value["ParentId"].(int64))
			}
			if parentIDInt == 0 && strGroup != "" {
				nStr = strGroup
			} else {
				nStr = str
			}

			//id Orm conversion may be int or int64 type, compatible
			idInt, ok := value["Id"].(int)
			if !ok {
				idInt64, ok := value["Id"].(int64)
				if ok {
					idInt = int(idInt64)
					nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(idInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$id", strconv.Itoa(idInt))
			}

			levelInt, ok := value["Level"].(int)
			if !ok {
				levelInt64, ok := value["Level"].(int64)
				if ok {
					levelInt = int(levelInt64)
					nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(levelInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$level", strconv.Itoa(levelInt))
			}

			sortIDInt, ok := value["SortId"].(int)
			if !ok {
				sortIDInt64, ok := value["SortId"].(int64)
				if ok {
					sortIDInt = int(sortIDInt64)
					nStr = strings.ReplaceAll(nStr, "$sort_id", strconv.Itoa(sortIDInt))
				}
			} else {
				nStr = strings.ReplaceAll(nStr, "$sort_id", strconv.Itoa(sortIDInt))
			}

			parentIDNodeStringValue, ok := value["ParentIdNode"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$parent_id_node", parentIDNodeStringValue)
			}

			nStr = strings.ReplaceAll(nStr, "$spacer", spacer)
			nStr = strings.ReplaceAll(nStr, "$selected", selected)

			nameValue, ok := value["Name"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$name", nameValue)
			}

			urlValue, ok := value["Url"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$url", urlValue)
			}

			nStr = strings.ReplaceAll(nStr, "$parent_id", strconv.Itoa(parentIDInt))

			iconValue, ok := value["Icon"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$icon", iconValue)
			}

			isShowValue, ok := value["IsShow"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$is_show", isShowValue)
			}

			logMethodValue, ok := value["LogMethod"].(string)
			if ok {
				nStr = strings.ReplaceAll(nStr, "$log_method", logMethodValue)
			}

			strManageValue, ok := value["StrManage"].(string)
			if ok {
				strManageValue = strings.ReplaceAll(strManageValue, "\\", "")
				nStr = strings.ReplaceAll(nStr, "$str_manage", strManageValue)
			}

			adminTreeService.Ret += nStr
			adminTreeService.GetTree(id, str, sid, adds+k+space, strGroup)

			number++
		}

	}

	return adminTreeService.Ret
}

// Menu Menu selection select tree selection
func (adminTreeService *AdminTreeService) Menu(selected int64, currentId int64) string {
	adminMenuService := NewAdminMenuService()
	result := adminMenuService.Menu(currentId)
	resultKey := make(map[int64]orm.Params)
	if result != nil {
		for _, r := range result {
			idInt, ok := r["Id"].(int64)
			if !ok {
				idInt = r["Id"].(int64)
			}
			resultKey[idInt] = r
			if idInt == selected {
				resultKey[idInt]["selected"] = "selected"
			} else {
				resultKey[idInt]["selected"] = ""
			}
		}

		str := `<option value='$id' $selected >$spacer $name</option>`
		adminTreeService.initTree(resultKey)
		return adminTreeService.GetTree(0, str, selected, "", "")
	}
	return ""
}

// AdminMenuTree
func (adminTreeService *AdminTreeService) AdminMenuTree() string {
	adminMenuService := NewAdminMenuService()
	adminMenus := adminMenuService.AllMenu()
	if adminMenus != nil {
		result := make(map[int64]orm.Params)
		var adminTreeService AdminTreeService
		for _, adminMenu := range adminMenus {
			n := adminMenu.Id
			//Initialize the orm.Params map type
			if result[n] == nil {
				result[n] = make(orm.Params)
			}

			result[n]["Id"] = adminMenu.Id
			result[n]["ParentId"] = adminMenu.ParentId
			result[n]["Name"] = adminMenu.Name
			result[n]["Url"] = adminMenu.Url
			result[n]["Icon"] = adminMenu.Icon
			result[n]["IsShow"] = adminMenu.IsShow
			result[n]["SortId"] = adminMenu.SortId

			result[n]["Level"] = adminTreeService.GetLevel(adminMenu.Id, result, 0)
			if adminMenu.ParentId > 0 {
				result[n]["ParentIdNode"] = ` class="child-of-node-` + strconv.FormatInt(adminMenu.ParentId, 10) + `"`
			} else {
				result[n]["ParentIdNode"] = ""
			}
			id := strconv.FormatInt(adminMenu.Id, 10)
			result[n]["StrManage"] = `<a href="/admin/admin_menu/edit?id=` + id + `" class="btn btn-primary btn-xs" title="[EDIT]" data-toggle="tooltip"><i class="fa fa-pencil"></i></a> <a class="btn btn-danger btn-xs AjaxButton" data-id="` + id + `" data-url="del"  data-confirm-title="[DELETE]" data-confirm-content=\'[DELETE_QUESTION] ID <span class="text-red"> ` + id + ` </span> \'  data-toggle="tooltip" title="[DELETE]"><i class="fa fa-trash"></i></a>`
			if adminMenu.IsShow == 1 {
				result[n]["IsShow"] = "<span style='width: 100%; margin:auto; display:table;' class='label label-success'>O</span>" //"显示"
			} else {
				result[n]["IsShow"] = "<span style='width: 100%; margin:auto; display:table;' class='label label-danger'>X</span>" //"隐藏"
			}
			result[n]["LogMethod"] = adminMenu.LogMethod
		}
		str := `<tr id='node-$id' data-level='$level' $parent_id_node><td><input type='checkbox' onclick='checkThis(this)'
                     name='data-checkbox' data-id='$id' class='checkbox data-list-check' value='$id' placeholder=''>
                    </td><td>$id</td><td>$spacer$name</td><td>$url</td>
                    <td>$parent_id</td><td><i class='fa $icon'></i><span>($icon)</span></td>
                    <td>$sort_id</td><td>$is_show</td><td>$log_method</td><td class='td-do'>$str_manage</td></tr>`

		adminTreeService.initTree(result)

		return adminTreeService.GetTree(0, str, 0, "", "")

	}
	return ""
}

// AuthorizeHtml
func (adminTreeService *AdminTreeService) AuthorizeHtml(menu map[int64]orm.Params, authMenus []string) string {
	for id := range menu {
		if utils.InArrayForString(authMenus, (string)(id)) {
			menu[id]["Checked"] = " checked"
		} else {
			menu[id]["Checked"] = ""
		}
		levelInt := adminTreeService.GetLevel(id, menu, 0)
		menu[id]["Level"] = levelInt
		menu[id]["Width"] = 100 - levelInt
	}

	adminTreeService.initTree(menu)

	adminTreeService.Text = make(map[string]interface{})
	adminTreeService.Text["other"] = `<label class='checkbox'  >
                        <input $checked  name='url[]' value='$id' level='$level'
                        onclick='javascript:checkNode(this);' type='checkbox'>
                       $name
                   </label>`

	adminTreeService.Text["0"] = []string{
		`<dl class='checkMod'>
                    <dt class='hd'>
                        <label class='checkbox'>
                            <input $checked name='url[]' value='$id' level='$level'
                             onclick='javascript:checkNode(this);'
                             type='checkbox'>
                            $name
                        </label>
                    </dt>
                    <dd class='bd'>`,
		`</dd></dl>`,
	}

	adminTreeService.Text["1"] = []string{
		`
                        <div class='menu_parent'>
                            <label class='checkbox'>
                                <input $checked  name='url[]' value='$id' level='$level'
                                onclick='javascript:checkNode(this);' type='checkbox'>
                               $name
                            </label>
                        </div>
                        <div class='rule_check' style='width: $width%;'>`,
		`</div><span class='child_row'></span>`,
	}

	return adminTreeService.getAuthTreeAccess(0)
}

// getAuthTreeAccess
func (adminTreeService *AdminTreeService) getAuthTreeAccess(myID int64) string {
	nStr := ""
	child := adminTreeService.getChild(myID)

	if len(child) > 0 {
		//Take the smallest key to prevent random selection of for range, resulting in a different menu order each time
		var sortID int64
		sortID = 99999
		for k := range child {
			if k < sortID {
				sortID = k
			}
		}
		level := make(map[string]interface{})
		_ = level
		level = child[sortID]

		//child
		var ids []int64
		for id := range child {
			ids = append(ids, id)
		}
		//sort.Ints(ids)
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
		var text []string
		//lvl, _ := strconv.ParseInt(level["Level"].(string), 10, 64)
		//if v, ok := adminTreeService.Text[strconv.Itoa(level["Level"].(int))]; ok {
		if v, ok := adminTreeService.Text[strconv.FormatInt(level["Level"].(int64), 10)]; ok {
			text = v.([]string)
		} else {
			text = adminTreeService.Text["1"].([]string)
		}

		for _, id := range ids {
			k := id
			v := child[id]
			if len(adminTreeService.getChild(k)) > 0 {
				nStr = text[0]
				nStr = strings.ReplaceAll(nStr, "$id", strconv.FormatInt(v["Id"].(int64), 10))
				nStr = strings.ReplaceAll(nStr, "$checked", v["Checked"].(string))
				nStr = strings.ReplaceAll(nStr, "$level", strconv.FormatInt(v["Level"].(int64), 10))
				nStr = strings.ReplaceAll(nStr, "$name", v["Name"].(string))
				nStr = strings.ReplaceAll(nStr, "$width", strconv.FormatInt(v["Width"].(int64), 10))
				adminTreeService.Html += nStr

				adminTreeService.getAuthTreeAccess(k)

				nStr = text[1]
				nStr = strings.ReplaceAll(nStr, "$id", strconv.FormatInt(v["Id"].(int64), 10))
				nStr = strings.ReplaceAll(nStr, "$checked", v["Checked"].(string))
				nStr = strings.ReplaceAll(nStr, "$level", strconv.FormatInt(v["Level"].(int64), 10))
				nStr = strings.ReplaceAll(nStr, "$name", v["Name"].(string))
				nStr = strings.ReplaceAll(nStr, "$width", strconv.FormatInt(v["Width"].(int64), 10)) //strconv.Itoa(v["Width"].(int))
				adminTreeService.Html += nStr

			} else {
				nStr = adminTreeService.Text["other"].(string)
				nStr = strings.ReplaceAll(nStr, "$id", strconv.FormatInt(v["Id"].(int64), 10)) //strconv.Itoa(v["Id"].(int))
				nStr = strings.ReplaceAll(nStr, "$checked", v["Checked"].(string))
				nStr = strings.ReplaceAll(nStr, "$level", strconv.FormatInt(v["Level"].(int64), 10)) //strconv.Itoa(v["Level"].(int))
				nStr = strings.ReplaceAll(nStr, "$name", v["Name"].(string))
				adminTreeService.Html += nStr
			}
		}

		return adminTreeService.Html
	}

	return ""
}
