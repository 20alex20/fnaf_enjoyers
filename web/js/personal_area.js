document.getElementById("btn").addEventListener("click", function () {
    var selects = document.getElementById("selects").children;
    var faculties = document.getElementById("faculties").children;
    var categories = [];
    for (var i = 0; i < 11; i++) {
        if (faculties[i].selected) {
            categories.push(faculties[i].value)
            var flag = false;
            var small_categories = [];
            for (var j = 0; j < selects[i + 1].children.length; j++)
                if (selects[i + 1].children[j].selected)
                    small_categories.push(selects[i + 1].children[j].value);
                else
                    flag = true;
            if (flag)
                while (small_categories.length > 0)
                    categories.push(small_categories.pop())
        }
    }
    if (faculties[11].selected)
        categories.push(faculties[11].value)

    var filters = [];
    for (i = 0; i < 3; i++) {
        var elem = document.getElementById("filter" + String(i + 1));
        if (elem.checked)
            filters.push(elem.value);
    }

    var text = document.getElementById("text").value.trim();

    if (categories.length == 0 || text == "")
        return;

    $.ajax({
        url: 'http://localhost:3002/post/create',  // 'http://localhost:3001/main/create_post'
        method: 'post',
        crossDomain: true,
        xhrFields: {
            withCredentials: true
        },
        data: {"categories": categories, "filters": filters, "text": text},
        success: function (data) {
            info();
            for (i = 0; i < 11; i++) {
                faculties[i].selected = false;
                for (var j = 0; j < selects[i + 1].children.length; j++)
                    selects[i + 1].children[j].selected = false;
                selects[i + 1].style.display = "none";
            }
            faculties[11].selected = false;
            selects[0].style.display = "none";
            for (i = 0; i < 3; i++) {
                document.getElementById("filter" + String(i + 1)).checked = false;
            }
            document.getElementById("text").value = "";
        }
    });
});

function truncate(str, maxlength) {
    if (str.length > maxlength) {
        return str.slice(0, maxlength - 3) + '...';
    }
    return str
}

$.ajax({
    url: 'http://localhost:3002/posts/liked',  // 'http://localhost:3001/main/my_posts',
    method: 'get',
    crossDomain: true,
    xhrFields: {
        withCredentials: true
    },
    success: function (data) {
        for (var j = 0; j < data.length; j++) {
            var obj = data[j];
            $("#content-1").append("<div>\n" +
                "                    <a href=\"post_view.html?id=" + obj["id"] + "\" class=\"b-link-stripe b-animate-go\" style=\"width: 100%;\">\n" +
                "                        <div class=\"post\">\n" +
                "                            <p>" + truncate(obj["text"], 500) + "</p>\n" +
                "                            <div><p style=\"margin-right: auto; padding-right: 10px;\">" + obj["date_time"] + "</p>\n" +
                "                                <img src=\"images/icons8-удивление-64.png\" width=\"40px\" height=\"40px\"/>\n" +
                "                                <p>" + String(obj["views"]) + "</p>\n" +
                "                                <img src=\"images/icons8-палец-вверх-64.png\" width=\"40px\" height=\"40px\"/>\n" +
                "                                <p>" + String(obj["likes"]) + "</p>\n" +
                "                            </div>\n" +
                "                        </div>\n" +
                "                        <div class=\"b-wrapper\"><h2 class=\"b-animate b-from-left    b-delay03 \"><img\n" +
                "                                src=\"images/icon-eye.png\" alt=\"\" style=\"top: 0px; margin-top: 5%;\"></h2>\n" +
                "                            <p class=\"b-animate b-from-right  b-delay03 \"><span class=\"m_4\">Читать дальше</span></p>\n" +
                "                        </div>\n" +
                "                    </a>\n" +
                "                </div>");
        }
    }
});

$.ajax({
    url: 'http://localhost:3002/posts',  // 'http://localhost:3001/main/my_posts'
    method: 'get',
    crossDomain: true,
    xhrFields: {
        withCredentials: true
    },
    success: function (data) {
        for (var j = 0; j < data.length; j++) {
            var obj = data[j];
            $("#content-2").append("<div>\n" +
                "                    <a href=\"post_view.html?id=" + obj["id"] + "\" class=\"b-link-stripe b-animate-go\" style=\"width: 100%;\">\n" +
                "                        <div class=\"post\">\n" +
                "                            <p>" + truncate(obj["text"], 500) + "</p>\n" +
                "                            <div><p style=\"margin-right: auto; padding-right: 10px;\">" + obj["date_time"] + "</p>\n" +
                "                                <img src=\"images/icons8-удивление-64.png\" width=\"40px\" height=\"40px\"/>\n" +
                "                                <p>" + String(obj["views"]) + "</p>\n" +
                "                                <img src=\"images/icons8-палец-вверх-64.png\" width=\"40px\" height=\"40px\"/>\n" +
                "                                <p>" + String(obj["likes"]) + "</p>\n" +
                "                            </div>\n" +
                "                        </div>\n" +
                "                        <div class=\"b-wrapper\"><h2 class=\"b-animate b-from-left    b-delay03 \"><img\n" +
                "                                src=\"images/icon-eye.png\" alt=\"\" style=\"top: 0px; margin-top: 5%;\"></h2>\n" +
                "                            <p class=\"b-animate b-from-right  b-delay03 \"><span class=\"m_4\">Читать дальше</span></p>\n" +
                "                        </div>\n" +
                "                    </a>\n" +
                "                </div>");
        }
    }
});
function after_nick() {
    var dict = {funny: "Смешные", instructive: "Поучительные", condemning: "Осуждающие"}
    var dict2 = {IU: "ИУ", IBM: "ИБМ", SM: "СМ", E: "Э", MT: "МТ", RL: "РЛ", BMT: "БМТ", RK: "РК", FN: "ФН", L: "Л",
        SGN: "СГН", UR: "ЮР"}
    $.ajax({
        url: 'http://localhost:3002/posts/rejected',  // 'http://localhost:3001/main/my_posts'
        method: 'get',
        crossDomain: true,
        xhrFields: {
            withCredentials: true
        },
        success: function (data) {
            $("#content-5").empty();
            for (var j = 0; j < data.length; j++) {
                var obj = data[j];
                var arr = [];
                for (var i = 0; i < obj["filters"].length; i++)
                    arr.push(dict[obj["filters"][i]]);
                var arr2 = [];
                for (i = 0; i < obj["categories"].length; i++)
                    if (obj["categories"][i].includes("-"))
                        arr2.push(dict2[obj["categories"][i].split("-")[0]] + '-' + obj["categories"][i].split("-")[1]);
                    else
                        arr2.push(dict2[obj["categories"][i]])
                $("#content-5").append("<div>\n" +
                    "                    <div class=\"post\">\n" +
                    "                        <div class=\"post-header\">\n" +
                    "                            <span>Автор: " + nickname + "</span>\n" +
                    "                            <span>Категории: " + arr2.join(', ') + "</span>\n" +
                    "                        </div>\n" +
                    "                        <div class=\"post-header\">\n" +
                    "                            <span>Время: " + obj["date_time"] + "</span>\n" +
                    "                            <span>Подкатегории: " + arr.join(', ') + "</span>\n" +
                    "                        </div>\n" +
                    "                        <p style=\"font-size: 14px;\">" + obj["text"] + "</p>\n" +
                    "                        <div><p style=\"font-size: 16px;font-weight: bold;margin: 0;\">Сообщение модератора:</p></div>\n" +
                    "                        <div id=\"moder_message\"><p>" + obj["moder_text"] + "</p></div>\n" +
                    "                    </div>\n" +
                    "                </div>");
            }
        }
    });
}