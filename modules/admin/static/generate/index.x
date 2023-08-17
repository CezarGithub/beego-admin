<section class="content">
    <div class="row">
        <div class="col-md-12">
            <div class="box">
                <div class="box-header ">
                    {{button "[[.Model.Name_lcase]].add"  $.buttons  nil}}
                    {{button "[[.Model.Name_lcase]].toggle"  $.buttons  nil}}
                    {{button "[[.Model.Name_lcase]].delete"  $.buttons  nil}}
                    {{button "[[.Model.Name_lcase]].export"  $.buttons  nil}}
                </div>
                <div class="box-body">
                    <form class="form-inline searchForm" id="searchForm" action='{{urlfor "[[$.Model.Name]]Controller.Index"}}' method="GET">
                        <div class="form-group">
                            <input value="{{._keywords}}"
                                   name="_keywords" id="_keywords" class="form-control input-sm" placeholder='{{i18n "app.name"}}'>
                        </div>
                        <div class="form-group">
                            <button class="btn btn-sm btn-primary" type="submit"><i class="fa fa-search"></i> {{i18n "app.search"}}
                            </button>
                        </div>
                        <div class="form-group">
                            <button onclick="clearSearchForm()" class="btn btn-sm btn-default" type="button"><i
                                    class="fa  fa-eraser"></i> {{i18n "app.reset"}}
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>

    <div class="row">
        <div class="col-md-12">
            <div class="box">
                <div class="box-body table-wrap">
                    <table class="table table-hover table-bordered " id="tbl[[$.Model.Name]]" width="100%">
                        <thead>
                        <tr>
                            <th>
                                <input id="dataCheckAll" type="checkbox" onclick="checkAll(this)" class="checkbox" placeholder='{{i18n "app.toggle"}}'>
                            </th>
                            <th>ID</th>
						[[- range $item := .Model.Fields]]
                            <th>{{i18n "[[$.Model.Module]].[[$.Model.Name_lcase]].[[$item.Name_lcase]]"}}</th>
						[[- end]]
                            <th>{{i18n "db.actions"}}</th>							
                        </tr>
                        </thead>
                        <tbody>
                        {{range $key,$item := .data}}
                        <tr>
                            <td>
                                <input type="checkbox" onclick="checkThis(this)" name="data-checkbox"
                                       data-id="{{$item.Id}}" class="checkbox data-list-check" value="{{$item.Id}}">
                            </td>
							<td>{{$item.Id}}</td>
						[[- range $field := .Model.Fields]] 				
							[[- if eq $field.Type "int8" ]]
								<td>{{if eq $item.[[$field.Name]] 1}} <span class="label label-success">{{i18n "app.yes"}}</span> {{else}} <span class="label label-warning">{{i18n "app.no"}}</span> {{end}}</td>
							[[- else]]
								<td>{{$item.[[$field.Name]]}}</td>
							[[- end]]                            
						[[- end]]
                            <td class="td-do">
                                {{toolbar "[[$.Model.Name_lcase]].grid" $.toolbars $item.Id }}                       
                            </td>
                        </tr>
                        {{end}}
                        </tbody>
                    </table>
                </div>
                <!-- Bottom of data list -->

            </div>
        </div>
    </div>
</section>

