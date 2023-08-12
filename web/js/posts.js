function truncate(str, maxlength) {
    if (str.length > maxlength) {
        return str.slice(0, maxlength - 3) + '...';
    }
    return str
}

function ajax() {
    document.getElementById("post_20").innerHTML = '';
    $.ajax({
        url: 'http://localhost:3001/main/posts',  // 'http://localhost:3001/main/posts'
        method: 'get',
        dataType: 'json',
        data: params,
        success: function (data) {
            for (j = 0; j < data.length; j++) {
                var obj = data[j];
                $("#post_20").append("<div>\n" +
                    "                    <a href=\"post_view.html?id=" + obj["id"] + "\" class=\"b-link-stripe b-animate-go\">\n" +
                    "                        <div class=\"post\">\n" +
                    "                            <p>" + truncate(obj["text"], 500) + "</p>\n" +
                    "                            <div><p style=\"margin-right: auto;\">" + obj["date_time"] + "</p>\n" +
                    "                                <img src=\"images/icons8-удивление-64.png\" width=\"40px\" height=\"40px\"/>\n" +
                    "                                <p>" + String(obj["views"]) + "</p>\n" +
                    "                                <img src=\"images/icons8-палец-вверх-64.png\" width=\"40px\" height=\"40px\"/>\n" +
                    "                                <p>" + String(obj["likes"]) + "</p>\n" +
                    "                            </div>\n" +
                    "                        </div>\n" +
                    "                        <div class=\"b-wrapper\"><h2 class=\"b-animate b-from-left    b-delay03 \"><img\n" +
                    "                                src=\"images/icon-eye.png\" alt=\"\" style=\"top: 0; margin-top: 5%;\"></h2>\n" +
                    "                            <p class=\"b-animate b-from-right  b-delay03 \"><span class=\"m_4\">Читать дальше</span></p>\n" +
                    "                        </div>\n" +
                    "                    </a>\n" +
                    "                </div>");
            }
            document.getElementsByClassName("header-top")[0].scrollIntoView();
        }
    });
}

function reboot_url() {
    history.pushState({},"",window.location.href.split('?')[0]+"?category="+params["category"]+"&filter="+
        params["filter"] + "&sort=" + params["sort"] + "&page=" + params["page"]);
}

document.getElementById("filt").addEventListener("change", function () {
    var selects = document.getElementById("filt").children;
    for (var i = 0; i < 4; i++) {
        if (selects[i].selected) {
            params["filter"] = selects[i].getAttribute("value")
            break;
        }
    }
    reboot_url();
    ajax();
});

document.getElementById("sort").addEventListener("change", function () {
    var selects = document.getElementById("sort").children;
    for (var i = 0; i < 3; i++) {
        if (selects[i].selected) {
            params["sort"] = selects[i].getAttribute("value")
            break;
        }
    }
    reboot_url();
    ajax();
});

function numbers() {
    document.getElementById("numbers").innerHTML = '';
    var current = parseInt(params["page"])
    var max, min
    if (current == 1 || current == 2) {
        min = 1
        max = 5 <= max_page ? 5 : max_page;
    } else if (current == max_page - 1 || current == max_page) {
        min = (max_page - 4) >= 1 ? (max_page - 4) : 1;
        max = max_page;
    } else {
        min = current - 2
        max = current + 2
    }
    flag1 = min > 1;
    flag2 = max < max_page;

    if (flag1)
        $("#numbers").append("<li class=\"page-item\" id='first'>\n" +
            "                            <span class=\"page-link\" aria-label=\"Previous\">\n" +
            "                                <span aria-hidden=\"true\">&laquo;</span>\n" +
            "                            </span>\n" +
            "                        </li>");
    for (var j = min; j <= max; j++) {
        $("#numbers").append("<li class=\"page-item number_pages" + (j == current ? " active" : "") + "\"><span class=\"page-link\" num='" + String(j) + "'>" + String(j) + "</span></li>");
    }
    if (flag2)
        $("#numbers").append("<li class=\"page-item\" id='last'>\n" +
            "                            <span class=\"page-link\" aria-label=\"Next\">\n" +
            "                                <span aria-hidden=\"true\">&raquo;</span>\n" +
            "                            </span>\n" +
            "                        </li>");

    if (flag1) {
        document.getElementById("first").children[0].addEventListener("click", function () {
            params["page"] = String(parseInt(params["page"]) - 1);
            reboot_url();
            ajax();
            numbers();
        });
    }
    if (flag2) {
        document.getElementById("last").children[0].addEventListener("click", function () {
            params["page"] = String(parseInt(params["page"]) + 1);
            reboot_url();
            ajax();
            numbers();
        });
    }
    var elements = document.getElementsByClassName("number_pages");
    for (j = min; j <= max; j++) {
        elements[j - min].children[0].addEventListener("click", function () {
            params["page"] = this.getAttribute("num");
            reboot_url();
            ajax();
            numbers();
        });
    }
}

urlParams = new URLSearchParams(window.location.search);
var params = {};
urlParams.forEach((p, key) => {
    params[key] = p;
});
if (!("category" in params))
    params["category"] = "all";
if (!("filter" in params))
    params["filter"] = "without";
if (!("sort" in params))
    params["sort"] = "date_time";
if (!("page" in params))
    params["page"] = "1";
params["number"] = "20";

var selects = document.getElementById("filt").children;
for (var i = 0; i < 4; i++) {
    if (selects[i].getAttribute("value") == params["filter"]) {
        selects[i].selected = true;
        break;
    }
}
selects = document.getElementById("sort").children;
for (i = 0; i < 3; i++) {
    if (selects[i].getAttribute("value") == params["sort"]) {
        selects[i].selected = true;
        break;
    }
}

reboot_url();
ajax();

var max_page;
var flag1, flag2;
$.ajax({
    url: 'http://localhost:3001/main/max-page',  // 'http://localhost:3001/main/max_pages'
    method: 'get',
    dataType: 'json',
    data: params,
    success: function (data) {
        max_page = data["max_page"];
        numbers();
    }
});