<section class="content">
    <div class="row">
        <div class="col-md-12">
            <div class="box box-info">
                <div class="box-header with-border">
                    <h3 class="box-title"></h3>
					{{button "[[.Model.Name_lcase]].submit"  $.buttons  .data}}
                    <div class="box-tools pull-right">
                        <button type="button" class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i>
                        </button>
                        
                    </div>
                </div>
                <!-- box body -->
                <div class="box-body">
                    <!-- Form header -->
                    <form id="[[.Model.Name_lcase]]Form" class="form-horizontal dataForm" action="" method="post"
                        enctype="multipart/form-data">
                        <input type="hidden" name="id" value="{{.data.Id}}">
						[[- range $item := .Model.Fields]]
							[[- if eq $item.Type "int8" ]]
								<div class="form-group">
									<label for="[[$item.Name_lcase]]" class="col-sm-2 control-label">{{i18n "app.enabled"}}</label>
									<div class="col-sm-10 col-md-4">
										<input class="input-switch" id="[[$item.Name_lcase]]" value="1" {{if eq .data.[[$item.Name]] 1}}checked {{end}}type="checkbox"/>
										<input class="switch field-switch" placeholder='{{i18n  "app.enabled"}}' name="[[$item.Name_lcase]]"
											value="{{.data.[[$item.Name]] }}" hidden/>
									</div>
								</div>

								<script>
									$('#[[$item.Name_lcase]]').bootstrapSwitch({
										onText: "{{i18n  "app.yes"}}",
										offText: "{{i18n  "app.no"}}",
										onColor: "success",
										offColor: "danger",
										onSwitchChange: function (event, state) {
											$(event.target).closest('.bootstrap-switch').next().val(state ? '1' : '0').change();
										}
									});
								</script>

							[[- else if eq $item.Type "ptr" ]]
								//SearchBox
								<div class="form-group">
									<label for="[[$item.Name_lcase]]_id" class="col-sm-2 control-label">{{i18n "[[$.Model.Module]].[[$item.Name_lcase]].name"}}</label>
									<div class="col-sm-10 col-md-4">
										<input id="[[$item.Name_lcase]]_id" name="[[$item.Name_lcase]]_id"  value="{{.data.[[$item.Name]].Id}}" style="display:none"/>
										<div class="input-group iconpicker-container">
											<a  data-model="{{.searchBox[[$item.Name]]}}" 
												data-callback="search[[$item.Name]]" data-url='{{urlfor "[[$item.Name]]Controller.Search"}}' 
												class="input-group-addon btn btn-default btn-xs  SearchBox" data-title='{{i18n  "app.details"}}' title='{{i18n  "app.details"}}' data-toggle="tooltip">                                    
											</a>      
									
											<input maxlength="50" id="[[$item.Name_lcase]]" name="[[$item.Name_lcase]]" readonly="true"
												value="{{.data.[[$item.Name]].Name}}" class="form-control" placeholder='{{i18n  "app.description"}}'>
										</div>
									</div>
								</div>	
								
								//OR DropDown
									<div class="form-group">
										<label for="[[$item.Name_lcase]]_id" class="col-sm-2 control-label">{{i18n  "[[$.Model.Module]].[[$.Model.Name_lcase]].[[$item.Name_lcase]]"}}</label>
										<div class="col-sm-10 col-md-4">
											<select name="[[$.Model.Name_lcase]]_id" id="[[$item.Name_lcase]]_id" class="form-control field-select"
													data-placeholder='{{i18n  "[[$.Model.Module]].[[$.Model.Name_lcase]].[[$item.Name_lcase]]"}}'>
												<option value=""></option>
												{{range $key,$item := .[[$item.Name_lcase]]_list}}
												<option value="{{$item.Id}}" {{if compare $.data.[[$item.Name]].Id $item.Id}}selected {{end}}>
													{{$item.Name}}
												</option>
												{{end}}
											</select>
										</div>
									</div>
									<script>
										$('#[[$item.Name_lcase]]_id').select2({
											language: "{{.admin.jsLang}}"
										});
									</script>
									
							
							[[- else ]]
								<div class="form-group">
									<label for="name" class="col-sm-2 control-label">{{i18n  "[[$.Model.Module]].[[$.Model.Name_lcase]].[[$item.Name_lcase]]"}}</label>
									<div class="col-sm-10 col-md-4">
										<input maxlength="50" id="[[$item.Name_lcase]]" name="[[$item.Name_lcase]]" value="{{.data.[[$item.Name]]}}"
											class="form-control" placeholder='{{i18n  "[[$.Model.Module]].[[$.Model.Name_lcase]].[[$item.Name_lcase]]"}}'>
									</div>
								</div>
							[[- end]]
							
						[[- end]]      
                    </form>
                <!--box body-->
                </div>

            </div>
        </div>
    </div>
</section>

<script>
   function search[[.Model.Name]](data){
        $("#[[.Model.Name_lcase]]_id").val(data.id);
        $("#[[.Model.Name_lcase]]_name").val(data.name);
    }
 </script>

 

