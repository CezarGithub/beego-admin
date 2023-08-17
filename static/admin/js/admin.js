try {
    /** pjax */
    //Timeout 3 seconds (optional)
    $.pjax.defaults.timeout = 3000;
    $.pjax.defaults.type = 'GET';
    //Storage container id
    $.pjax.defaults.container = '#pjax-container';
    //Target id
    $.pjax.defaults.fragment = '#pjax-container';
    //Maximum buffer length (optional)
    $.pjax.defaults.maxCacheLength = 0;
    $(document).pjax('a:not(a[target="_blank"])', {
        //Storage container id
        container: '#pjax-container',
        //Target id
        fragment: '#pjax-container'
    });
    //ajax Request
    $(document).ajaxStart(function () {
        //Startup progress bar
        NProgress.start();
    }).ajaxStop(function () {
        //Execute after the end of ajax request
        //Close the progress bar
        NProgress.done();
    });
} catch (e) {
    if (adminDebug) {
        console.log(e.message);
    }
}

$(document).on('pjax:timeout', function (event) {
    event.preventDefault();
});
$(document).on('pjax:send', function (xhr) {
    NProgress.start();
});
$(document).on('pjax:complete', function (xhr) {
    $('[data-toggle="tooltip"]').tooltip();
    NProgress.done();
});
//List page search pjax
$(document).on('submit', '.searchForm', function (event) {
    $.pjax.submit(event, '#pjax-container');
});
//document ready events goes here
$(document).ready(function ($) {
		//alert('ready');
        var $pushMenu = $('[data-toggle="push-menu"]').data('lte.pushmenu');//toggle sidebar
        $pushMenu.expandOnHover();//expand on hover
        
})

//Menu search
$(function () {
    $('#sidebar-form').on('submit', function (e) {
        e.preventDefault();
    });

    $('.sidebar-menu li.active').data('lte.pushmenu.active', true);

    $('#search-input').on('keyup', function () {
        var term = $('#search-input').val().trim();

        if (term.length === 0) {
            $('.sidebar-menu li').each(function () {
                $(this).show(0);
                $(this).removeClass('active');
                if ($(this).data('lte.pushmenu.active')) {
                    $(this).addClass('active');
                }
            });
            return;
        }

        $('.sidebar-menu li').each(function () {
            if ($(this).text().toLowerCase().indexOf(term.toLowerCase()) === -1) {
                $(this).hide(0);
                $(this).removeClass('pushmenu-search-found', false);

                if ($(this).is('.treeview')) {
                    $(this).removeClass('active');
                }
            } else {
                $(this).show(0);
                $(this).addClass('pushmenu-search-found');

                if ($(this).is('.treeview')) {
                    $(this).addClass('active');
                }

                var parent = $(this).parents('li').first();
                if (parent.is('.treeview')) {
                    parent.show(0);
                }
            }

            if ($(this).is('.header')) {
                $(this).show();
            }
        });

        $('.sidebar-menu li.pushmenu-search-found.treeview').each(function () {
            $(this).find('.pushmenu-search-found').show(0);
        });
    });
});

//Click the menu to highlight
$(function () {
    $('.sidebar-menu li:not(.treeview) > a').on('click', function () {      
        var $parent = $(this).parent().addClass('active');
        $parent.siblings('.treeview.active').find('> a').trigger('click');
        $parent.siblings().removeClass('active').find('li').removeClass('active');
    });

    $('[data-toggle="popover"]').popover();
});

//bootstrap prompt
$(function () {
    $('[data-toggle="tooltip"]').tooltip();
});

/** Form validation error display - DEPRECATED - submit form repleced bellow
$.validator.setDefaults({
    errorElement: "span",
    errorClass: "help-block error",
    submitHandler: function (form) {
        formSubmit(form);
        return false;
    }
});*/ 

/** Submit function for dataForm class **/
$(function() {
	$('.dataFormSubmit').on('click',function(e){
		 e.preventDefault();
         var formName=$(this).data('form')
		 //var form=$('.dataForm');//old version by class
         var form=$('#' + formName);//select by id
		 formSubmit(form,e);
	});
});
/** Reset function for dataForm class **/
$(function() {
	$('.dataFormReset').on('click',function(e){
		 e.preventDefault();
		 var form=$('.dataForm');//select by class
		 form.trigger("reset");;
	});
});
/**Sidebar menu */
try {
    $('.sidebar-menu').tree();
} catch (e) {
    if (adminDebug) {
        console.log(e.message);
    }
}


/* Clear search form */
function clearSearchForm() {
    let url_all = window.location.href;
    let arr = url_all.split('?');
    let url = arr[0];
    $.pjax({url: url, container: '#pjax-container'});
}


$(function () {
    /* Back button*/
    $('body').on('click', '.BackButton', function (event) {
        event.preventDefault();
        history.back(1);
    });

    /* Refresh button */
    $('body').on('click', '.ReloadButton', function (event) {
        event.preventDefault();
        $.pjax.reload();
    });

});


/*Single selection and cancellation in the list*/
function checkThis(obj) {
    var id = $(obj).attr('value');
    if ($(obj).is(':checked')) {
        if ($.inArray(id, dataSelectIds) < 0) {
            dataSelectIds.push(id);
        }
    } else {
        if ($.inArray(id, dataSelectIds) > -1) {
            dataSelectIds.splice($.inArray(id, dataSelectIds), 1);
        }
    }

    var all_length = $("input[name='data-checkbox']").length;
    var checked_length = $("input[name='data-checkbox']:checked").length;
    if (all_length === checked_length) {
        $("#dataCheckAll").prop("checked", true);
    } else {
        $("#dataCheckAll").prop("checked", false);
    }
    console.log(dataSelectIds);
}

/*Select all/cancel*/
function checkAll(obj) {
    dataSelectIds = [];
    var all_check = $("input[name='data-checkbox']");
    if ($(obj).is(':checked')) {
        all_check.prop("checked", true);
        $(all_check).each(function () {
            dataSelectIds.push(this.value);
        });
    } else {
        all_check.prop("checked", false);
    }
}


/* Form submission */
function formSubmit(form,e) {
    console.log(e);
    let loadT = layer.msg('…', {icon: 16, time: 0, shade: [0.3, "#000"],scrollbar: false,});// i18n message
    //get action url from submit button data attribute
    let action = e.currentTarget.getAttribute("data-url");// $(form).attr('action');-old version get from form attr
    const urlParams = new URLSearchParams(window.location.search);//current windows url
    const myParam = urlParams.get('request_type');//check if was a modal (layer_open) request
    var queryParam="";
    if(myParam){
        queryParam="?request_type=layer_open" //add layer_open to request
    }
    //console.log(myParam);
    let method = $(form).attr('method');
    let data = new FormData($(form)[0]);
    //Set the global general _xsrf verification token
    data.set("_xsrf",$('meta[name="_xsrf"]').attr('content'));
   
    $.ajax({
            url: action + queryParam,
            dataType: 'json',
            type: method,
            data: data,
            contentType: false,
            processData: false,
            success: function (result) {
                layer.close(loadT);
                layer.msg(result.msg, {
                    icon: result.code ? 1 : 2,
                    scrollbar: false,
                });
                goUrl(result.url);
               
            },
            error: function (xhr, type, errorThrown) {
                //Exception handling；
                layer.msg('Error code : ' + xhr.status, {icon: 2,scrollbar: false,});// i18n message
            }
        }
    );
    return false;
}


/** Jump to specified url */
function goUrl(url = 1) {
    //console.log(url);
    //Clear the selected ID on the list page
    if (url !== 'url://current' && url !== 1) {
        dataSelectIds = [];
    }
    if (url === 'url://current' || url === 1) {
        console.log('Stay current page.');
    } else if (url === 'url://reload' || url === 2) {
        console.log('Reload current page.');
        $.pjax.reload();
    } else if (url === 'url://back' || url === 3) {
        console.log('Return to the last page.');
        //history.back(1); // back without refresh
        window.location=document.referrer;// back with refresh
    }else if (url === 4 || url === 'url://close-refresh') {
        console.log('Close this layer page and refresh parent page.');
        let indexWindow = parent.layer.getFrameIndex(window.name);
        //Refresh the parent page first
        parent.goUrl(2);
        //Then close the current layer pop-up window
        parent.layer.close(indexWindow);
    } else if (url === 5 || url === 'url://close-layer') {
        console.log('Close this layer page.');
        let indexWindow = parent.layer.getFrameIndex(window.name);
        parent.layer.close(indexWindow);
    } else {
        console.log('Go to ' + url);
        try {
            $.pjax({
                url: url,
                container: '#pjax-container'
            });
        } catch (e) {
            window.location.href = url;
        }
    }
}

/**
 * ajax access button
 * For example, the element is <a class="AjaxButton" data-confirm="1" data-type="1" data-url="disable" data-id="2" data-go="" ></a>
 * data-confirm Whether to pop up a prompt，1 yes，2 no。For example, delete a piece of data，data-confirm="1"Will pop up to prompt
 * data-type Access method，1 is direct ajax access，For example, delete operation。2 is to open the layer window to display data, such as viewing the operation log details
 * data-url Is the url to be visited
 * data-id ID of the data to be operated，You can fill in the normal data ID，E.g data-id="2"，
 * Or fill in checked Represents the ID of the current data list selection，That is to take the value of the variable dataSelectIds
 * data-go Jump after the operation is completed url，If this parameter is not set, the default will be redirected according to the url returned in the background
 * data-confirm-title The title of the pop-up window for confirmation prompt E.g data-confirm-title="Delete warning"
 * data-confirm-contentTo confirm the content of the prompt E.g data-confirm-content="Are you sure you want to delete this data？"
 * data-title The title of the window display
 *
 */
$(function () {
    $('body').on('click', '.AjaxButton', function (event) {
        event.preventDefault();

        if (adminDebug) {
            console.log('AjaxButton clicked.');
        }
        //Whether to pop up a prompt
        var layerConfirm = $(this).data("confirm") || 1;
        //1-For direct access，2-Display for the layer window
        var layerType = parseInt($(this).data("type") || 1);
        //Visited url
        var url = $(this).data("url");
        //Access method, default post
        var layerMethod = $(this).data("method") || 'post';
        //The page that jumps to after a successful visit，If this parameter is not set, the default will be redirected according to the url returned in the background
        var go = $(this).data("go") || 'url://reload';

        //When displaying for a window, you can define the width and height
        var layerWith = $(this).data("width") || '80%';
        var layerHeight = $(this).data("height") || '60%';

        //The title of the window
        var layerTitle = $(this).data('title');
		

        //ID of current operation data
        var dataId = $(this).data("id");
        var formName = $(this).data("form");
        var dataData ={};

        //If no ID is defined to query the data-data attribute
        if (dataId === undefined) {
            if (formName == undefined){
                var dataData = $(this).data("data") || {};
            }else {
                
                var form=$('#' + formName);//select by id
                console.log(form);
                dataData = convertFormToJSON(form);
            }
        } else {
            if (dataId === 'checked') {
                if (dataSelectIds.length === 0) {
                    layer.msg('Please select the data to be operated', {icon: 2,scrollbar: false,});
                    return false;
                }
                dataId = dataSelectIds;
            }

            dataData = {"id": dataId};
        }

        //Ajax sets Beego's xsrf token verification // not used ?
        $.ajaxSetup({
            headers: {
                'X-XSRFToken': $('meta[name="_xsrf"]').attr('content')
            }
        });

        if (typeof (dataData) != 'object') {
            dataData = JSON.parse(dataData);
        }

        /*Need to confirm operation*/
        if (parseInt(layerConfirm) === 1) {
            //Title of the prompt window
            var confirmTitle = $(this).data("confirmTitle") || 'Operation confirmation';
            //The content of the prompt window
            var confirmContent = $(this).data("confirmContent") || 'Are you sure you want to perform this operation?';
            layer.confirm(confirmContent, {title: confirmTitle, closeBtn: 1, icon: 3}, function () {
                //If it is direct access
                if (layerType === 1) {
                    ajaxRequest(url, layerMethod, dataData, go);
                } else if (layerType === 2) {
                    //If it is an open window
                    //Permission query first
                    if (checkAuth(url)) {
                        layer.open({
                            type: 1,
                            area: [layerWith, layerHeight],
                            title: layerTitle,
                            closeBtn: 1,
                            shift: 0,
                            content: url + "?request_type=layer_open&" + parseParam(dataData),
                            scrollbar: false,
                        });
                    }
                }
            });
        } else {
            //No operation confirmation required
            if (layerType === 1) {
                //Direct request
                ajaxRequest(url, layerMethod, dataData, go);
            } else if (layerType === 2) {
                //pop up
                //Check permissions
                if (checkAuth(url)) {
                    //Open with window
                    layer.open({
                        type: 2,
                        area: [layerWith, layerHeight],
                        title: layerTitle,
                        closeBtn: 1,
                        shift: 0,
                        content: url + "?request_type=layer_open&" + parseParam(dataData),
                        scrollbar: false,
                    });
                }
            }
        }
    });
});



/**
 * SearchBox 
 * params : header : columns names for search result  -| separated
 *          fields : serach fields for models - | separated
 * 
 *
 */
 $(function () {
		//add search icon
 		$('.SearchBox').wrapInner("<span><i class='fa fa-search'></i></span>");

});
$(function () {
    $('body').on('click', '.SearchBox', function (event) {
        event.preventDefault();

        if (adminDebug) {
            console.log('SearchBox clicked.');
        }
		
		var settings={
			title: "Search..",
			url: "",
			imgLoader: '<img src="images/ajax-loader.gif">',
			notFoundMessage: "Data not found!",
			requestErrorMessage: "Request error!",
			header: null,
			searchBoxId: "lookupbox-search-box",
			searchTextId: "lookupbox-search-key",
			searchResultId: "lookupbox-search-result",
			searchFilterId: "lookupbox-filter",
			htmlForm: '<div>error</div>',
			onItemSelected: null,
			onSearch: null,
			item: null,
			fields: [],
		}
        //Visited url
        settings.url = $(this).data("url");

        //When displaying for a window, you can define the width and height
        var layerWith = $(this).data("width") || '80%';
        var layerHeight = $(this).data("height") || '60%';

        //The title of the window
        settings.title = $(this).data('title');
        
		var model=$(this).data("model").split('\t');       
        settings.header=model[0].split('|');
        settings.fields=model[1].split('|');
     
		//settings.header=$(this).data("header").split('|');
		//settings.fields=$(this).data("fields").split('|');
		
		//The callback
		settings.onItemSelected = $(this).data('callback');
		
/* 		console.log(settings.onItemSelected);
		if (typeof window[settings.onItemSelected] === "function"){
			console.log('is function');
			window[settings.onItemSelected].call(this,'bau');
		} */

        //Ajax sets Beego's xsrf token verification // not used ?
        $.ajaxSetup({
            headers: {
                'X-XSRFToken': $('meta[name="_xsrf"]').attr('content')
            }
        });
		settings.htmlForm= '<section class="content"> <div class="row"><div class="box"><div class="box-body"><div class="form-inline searchForm">'
		settings.htmlForm+='<select class="form-control select2" id="lookupbox-filter"></select>'
		settings.htmlForm+='<input class="form-control" id="lookupbox-search-key" type="text" name="key" placeholder="..." />'
		settings.htmlForm+='<input class="btn flat btn-info dataFormSubmit" type="button" id="lookupbox-search-box" value="Search" />'
		settings.htmlForm+='<span id="loading1"></span>'
		settings.htmlForm+='<div id="lookupbox-search-result" class="box-body table-wrap"></div>'
		settings.htmlForm+='</div></div></div></div></section>'
  	    if (checkAuth(settings.url)) {
		   //Open with window
			layer.open({
				type: 1,
				area: [layerWith, layerHeight],
				title: settings.title,
				closeBtn: 1,
				shift: 0,
				content: settings.htmlForm,
				scrollbar: false,
			});
		}

		  //add filter options
		  $.each(settings.fields, function (i, item) {
			  if(settings.fields[i]!="id"){
				$('#lookupbox-filter').append($('<option>', {
				  value: settings.fields[i],
				  text: settings.header[i]
				}));
			  }
		  });
		
		$("#lookupbox-search-box").click(function(){
			  if (settings.onSearch == null) {
			  $.ajax({
				beforeSend: function(){
				  $('#' + settings.loadingDivId).html(settings.imgLoader);
				},
				url: settings.url + "?" + $("#" + settings.searchFilterId).val() + "=" + $("#" + settings.searchTextId).val(),
				success: function(result) {
				  try {
					var data = null;

					if (typeof result == 'string')
					  data = $.parseJSON(result);
					else if (typeof result == 'object')
					  data = result;

					settings.item = data;

					if (data.length > 0) {;
					  var table = "<table  id='lookupbox-result' class='table table-hover table-bordered'>";
					  table=table + "<thead>";
					  table = table + "<tr id='lookupbox-result-header' class='lookupbox-result-header'>";
					  i = 0;
					  for (var key in data[0]){
						if ($.inArray(key, settings.fields) >=0) {
						  if ($.type(settings.header) == "object")
							idx = key;
						  else
							idx = i;

						  colName = key;
						  if (settings.header != null) {
							if (typeof(settings.header[idx]) != "undefined") {
							  colName = settings.header[idx];
							}
						  }

						  if ($.type(settings.colWidth) == "object")
							idx = key;
						  else
							idx = i;

						  colWidth = "";
					/**	  if (settings.colWidth != null) {
							if (typeof(settings.colWidth[idx]) != "undefined") {
							  if (settings.colWidth[idx] != null) {
								colWidth = " style='width: " + settings.colWidth[idx] + "'";
							  }
							}
						  }**/
						  var hidden ="" ;
						  if (key=="id"){hidden="display:none;"};//hide id column
						  table = table + "<th style='" +hidden + "'"+ colWidth + "data-title='" + colName + "' >" + colName + "</th>";

						  i++;
						}
					  }
					  table = table + "</tr>";
					  table=table + "</thead>";
					  table=table + "<tbody>";
					  for(i=0;i<data.length;i++){
						var rowClass = 'lookupbox-result-row odd';
						if (i % 2 == 0) {
						  rowClass = 'lookupbox-result-row even';
						}
						table = table + "<tr id='lookupbox-result-row-" + i + "' class='" + rowClass + "' >";
						k=0;
						for (var key in data[i]){
						  if ($.inArray(key, settings.fields) >= 0) {
							var hidden ="" ;
							if (key=="id"){hidden=" display:none"};
							table = table + "<td style='cursor:pointer;" +  hidden + "' data-title='" + settings.header[k] + "'>" + data[i][key] + "</td>";
							k++;
						  }
						}
						table = table + "</tr>";
					  }
					  table = table + "</table>";
					  

					  $("#" + settings.searchResultId).html(table);

					  if (settings.onItemSelected != null) {
						$("#lookupbox-result tr").click(function(){
						  var arr = $(this).attr("id").split("-");
                          console.log(arr);
						  if (typeof  window[settings.onItemSelected] === "function") {
							window[settings.onItemSelected].call(this, settings.item[arr[arr.length - 1]]);
						  }
						  //$dialog.dialog("close");
						  layer.closeAll();
						});
					  }
					}
					else {
					  $("#" + settings.searchResultId).html(settings.notFoundMessage);
					}
				  }
				  catch(e) {
					$("#" + settings.searchResultId).html(settings.requestErrorMessage);
				  }
				  $('#' + settings.loadingDivId).html('');
						},
				error: function(xhr, status, ex) {
				  $("#" + settings.searchResultId).html(settings.requestErrorMessage);
				}
			  });
			}
			else{
			  if (typeof settings.onSearch === "function") {
				settings.onSearch.call();
			  }
			}
		});
		
		$("#" + settings.searchTextId).keyup(function(e){
        if(e.keyCode == 13){
          $("#" + settings.searchBoxId).trigger('click');
        }
      });
        
    });
});


//ajax Request package
/**
 *
 * @param url Visited url
 * @param method  interview method
 * @param data  data
 * @param go Url to be redirected
 */
function ajaxRequest(url, method, data, go) {
    var loadT = layer.msg('Requesting, please wait...', {icon: 16, time: 0, shade: [0.3, '#000'],scrollbar: false,});
    $.ajax({
            url: url,
            dataType: 'json',
            type: method,
            data: data,
            success: function (result) {
                layer.close(loadT);
                layer.msg(result.msg, {
                    icon: result.code ? 1 : 2,
                    scrollbar: false,
                });

                goUrl(go);
            },
            error: function (xhr, type, errorThrown) {
                //Exception handling
                layer.msg('Access error, code' + xhr.status, {icon: 2,scrollbar: false,});
            }
        }
    );
}

//Change the number of pages
function changePerPage(obj) {
    Cookies.set(cookiePrefix + 'admin_per_page', obj.value, {expires:30});
    $.pjax.reload();
}


/**
 * Check authorization
 */
function checkAuth(url) {
    var hasAuth = false;
    var loadT = layer.msg('Requesting, please wait...', {icon: 16, time: 0, shade: [0.3, '#000'],scrollbar: false,});
    $.post({
        url: url,
        data: {"check_auth": 1},
        dataType: 'json',
        async: false,
        success: function (result) {
            layer.close(loadT);
            if (result.code === 1) {
                hasAuth = true;
            } else {
                layer.msg(result.msg, {
                    icon: 2,
                    scrollbar: false,
                });
            }
        },
        error: function (xhr, type, errorThrown) {
            layer.msg('Access error, code' + xhr.status, {icon: 2,scrollbar: false,});
        }
    });
    return hasAuth;
}

/** Processing url parameters **/
function parseParam(param, key) {
    var paramStr = "";
    if (param instanceof String || param instanceof Number || param instanceof Boolean) {
        paramStr += "&" + key + "=" + encodeURIComponent(param);
    } else {
        $.each(param, function (i) {
            var k = key == null ? i : key + (param instanceof Array ? "[" + i + "]" : "." + i);
            paramStr += '&' + parseParam(this, k);
        });
    }
    return paramStr.substr(1);
}

/** Export excel **/
function exportData(url) {
    var exportUrl = url || 'index';
    var openUrl = exportUrl + '?export_data=1&' + $("#searchForm").serialize();
    window.open(openUrl);

}
/** Parse Form to json */
function convertFormToJSON(form) {
    const array = $(form).serializeArray(); // Encodes the set of form elements as an array of names and values.
    const json = {};
    $.each(array, function () {
      json[this.name] = this.value || "";
    });
    return json;
  }
