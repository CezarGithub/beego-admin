/**
 * 
 */
(function ($) {
    'use strict';

    $.fn.sidebarMenu = function (options) {
        options = $.extend({}, $.fn.sidebarMenu.defaults, options || {});
        var $menu_ul = $(this);
        var level = 0;
        if (options.data) {
            var source = buildData(options.data);
            buildUL($menu_ul, source);
        }
        else {
            if (!options.url) return;
            $.getJSON(options.url, options.param, function (data) {
                const obj = JSON.parse(data);
                var source = buildData(obj);
                buildUL($menu_ul, source);
            });
        }

        function buildData(data) {
            var source = [];
            var items = [];
            // build hierarchical source.
            for (let i = 0; i < data.length; i++) {
                var item = data[i];
                var label = item["Name"];
                var parentid = item["ParentId"];
                var id = item["Id"];
                var icon = "fa " + item["Icon"];
                var url ="/" + item["Url"];
                if (items[parentid]) {
                    var item = { parentid: parentid, label: label,url:url,icon:icon, item: item };
                    if (!items[parentid].items) {
                        items[parentid].items = [];
                    }
                    items[parentid].items[items[parentid].items.length] = item;
                    items[id] = item;
                }
                else {
                    items[id] = { parentid: parentid, label: label,url:url,icon:icon, item: item };
                    source[id] = items[id];
                }
            }
            return source;
        } 
        
        function buildUL(parent, items) {
            //console.log(items);
            $.each(items, function (index) {
                if(this){//skip empty slots
                if (this.label) {
                    var $a = $('<a href="javascript:void(0);" ></a>');
                    //icon
                    var $icon = $('<i></i>');
                    $icon.addClass(this.icon);
                    //title
                    var $title = $('<span>' + this.label + '</span>');
 
                    $a.append($icon);
                    $a.append($title);               
                    // create LI element and append it to the parent element.
                    var li = $('<li class="treeview"></li>');
                    var isOpen = this.isOpen;
                    if (isOpen === true) {
                        li.addClass("active");
                    }
 
                    li.appendTo(parent);
                    // if there are sub items, call the buildUL function.
                    if (this.items && this.items.length > 0) {
                        var pullSpan = $('<span class="pull-right-container"></span>');
                        var pullIcon = $('<i class="fa fa-angle-left pull-right"></i>');
                        pullSpan.append(pullIcon);
                        $a.append(pullSpan);
                        li.append($a);
                        var ul = $("<ul></ul>");
                        ul.addClass('treeview-menu');
                        if (isOpen === true) {
                            ul.css("display", "block");
                            ul.addClass("menu-open");
                        } else {
                            ul.css("display", "none");
                        }
                        ul.appendTo(li);
                        buildUL(ul, this.items);
                    }else{
                        $a.attr("href", this.url);
                        $a.on('click', function () { //for unknown reasons  href is not working   
                            window.location.href = this; //fix
                        });    
                        li.append($a);
                    }
                }
                }
            });

        }      

        // function init($menu_ul, data, level) {
        //     $.each(data, function (i, item) {
        //         //isHeader
        //         var $header = $('<li class="header"></li>');
        //         if (item.isHeader !== null && item.isHeader === true) {
        //             $header.append(item.text);
        //             $menu_ul.append($header);
        //             return;
        //         }

        //         //header li
        //         var li = $('<li class="treeview menu-open" data-level="' + level + '"></li>');

        //         //a
        //         var $a;
        //         if (level > 0) {
        //             $a = $('<a style="padding-left:' + (level * 20) + 'px"></a>');
        //         } else {
        //             $a = $('<a href="javascript:void(0);"></a>');
        //         }

        //         //icon
        //         var $icon = $('<i></i>');
        //         $icon.addClass(item.icon);

        //         //title
        //         var $title = $('<span class="title"></span>');
        //         $title.addClass('menu-text').text(item.text);

        //         $a.append($icon);
        //         $a.append($title);
        //         //$a.addClass("nav-link");

        //         var isOpen = item.isOpen;

        //         if (isOpen === true) {
        //             li.addClass("active");
        //         }
        //         if (item.children && item.children.length > 0) {
        //             var pullSpan = $('<span class="pull-right-container"></span>');
        //             var pullIcon = $('<i class="fa fa-angle-left pull-right"></i>');
        //             pullSpan.append(pullIcon);
        //             $a.append(pullSpan);
        //             li.append($a);

        //             var menus = $('<ul></ul>');
        //             menus.addClass('treeview-menu');
        //             if (isOpen === true) {
        //                 menus.css("display", "block");
        //                 menus.addClass("menu-open");
        //             } else {
        //                 menus.css("display", "none");
        //             }
        //             init(menus, item.children, level + 1);
        //             li.append(menus);
        //         }
        //         else {

        //             if (item.targetType != null && item.targetType === "blank") //new page
        //             {
        //                 $a.attr("href", item.url);
        //                // $a.attr("target", "_blank");
        //             }
        //             else if (item.targetType != null && item.targetType === "ajax") { //ajax
        //                 $a.attr("href", item.url);
        //                 $a.addClass("ajaxify");
        //             }
        //             else if (item.targetType != null && item.targetType === "iframe-tab") {
        //                 item.urlType = item.urlType ? item.urlType : 'relative';
        //                 var href = 'addTabs({id:\'' + item.id + '\',title: \'' + item.text + '\',close: true,url: \'' + item.url + '\',urlType: \'' + item.urlType + '\'});';
        //                 $a.attr('onclick', href);
        //             }
        //             else if (item.targetType != null && item.targetType === "iframe") { //iframe
        //                 $a.attr("href", item.url);
        //                 $a.addClass("iframeOpen");
        //                 $("#iframe-main").addClass("tab_iframe");
        //             } else {
        //                 $a.attr("href", item.url);
        //                 $a.addClass("iframeOpen");
        //                 $("#iframe-main").addClass("tab_iframe");
        //             }
        //             //$a.addClass("nav-link");
        //             var badge = $("<span></span>");
        //             // <span class="badge badge-success">1</span>
        //             if (item.tip != null && item.tip > 0) {
        //                 badge.addClass("label").addClass("label-success").text(item.tip);
        //             }
        //             $a.append(badge);
        //             li.append($a);
        //         }
        //         $menu_ul.append(li);
        //     });
        //}

        // for Iframe version
        // $menu_ul.on("click", "li.treeview a", function () {
        //     var $a = $(this);

        //     if ($a.next().size() == 0) {//size>0
        //         if ($(window).width() < $.AdminLTE.options.screenSizes.sm) {//
        //             //
        //             $($.AdminLTE.options.sidebarToggleSelector).click();
        //         }
        //     }
        // });
    };

    $.fn.sidebarMenu.defaults = {
        url: null,
        param: null,
        data: null,
        isHeader: false
    };
})(jQuery);
