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

    var filters = [];
    for (i = 0; i < 3; i++) {
        var elem = document.getElementById("filter" + String(i + 1));
        if (elem.checked)
            filters.push(elem.value);
    }

    var text = document.getElementById("text").value;

    $.ajax({
        url: 'json/server_accept.json',
        method: 'post',
        dataType: 'json',
        data: {"categories[]": categories, "filters[]": filters, "text": text}
    });

    for (i = 0; i < 11; i++) {
        faculties[i].selected = false;
        for (var j = 0; j < selects[i + 1].children.length; j++)
            selects[i + 1].children[j].selected = false;
        selects[i + 1].style.display = "none";
    }
    selects[0].style.display = "none";
    for (i = 0; i < 3; i++) {
        document.getElementById("filter" + String(i + 1)).checked = false;
    }
    document.getElementById("text").value = "";

    document.getElementById("message").style.display = "block";
    setTimeout(function () {
        document.getElementById("message").style.backgroundColor = "#4cae4c";
        document.getElementById("message").style.color = "white";
        setTimeout(function () {
            document.getElementById("message").style.backgroundColor = "transparent";
            document.getElementById("message").style.color = "transparent";
            setTimeout(function () {
                document.getElementById("message").style.display = "none";
            }, 1100);
        }, 2000);
    }, 100);
});

function truncate(str, maxlength) {
    if (str.length > maxlength) {
        return str.slice(0, maxlength - 3) + '...';
    }
    return str
}

$.ajax({
    url: 'json/posts_20.json', // 'http://localhost:3001/main/my_posts',
    method: 'post',
    dataType: 'json',
    data: {id: '3490589089389489', what: "my"},
    success: function (data) {
        for (var j = 0; j < data.length; j++) {
            var obj = data[j];
            $("#content-1").append("<div>\n" +
                "                    <a href=\"http://localhost:3001/main/post_view?id=" + String(obj["id"]) + "\" class=\"b-link-stripe b-animate-go\">\n" +
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
                "                                src=\"images/icon-eye.png\" alt=\"\" style=\"top: 0px; margin-top: 5%;\"></h2>\n" +
                "                            <p class=\"b-animate b-from-right  b-delay03 \"><span class=\"m_4\">Читать дальше</span></p>\n" +
                "                        </div>\n" +
                "                    </a>\n" +
                "                </div>");
        }
    }
});

$.ajax({
    url: 'json/posts_20.json', // 'http://localhost:3001/main/my_posts',
    method: 'post',
    dataType: 'json',
    data: {id: '3490589089389489', what: "liked"},
    success: function (data) {
        for (var j = 0; j < data.length; j++) {
            var obj = data[j];
            $("#content-2").append("<div>\n" +
                "                    <a href=\"http://localhost:3001/main/post_view?id=" + String(obj["id"]) + "\" class=\"b-link-stripe b-animate-go\">\n" +
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
                "                                src=\"images/icon-eye.png\" alt=\"\" style=\"top: 0px; margin-top: 5%;\"></h2>\n" +
                "                            <p class=\"b-animate b-from-right  b-delay03 \"><span class=\"m_4\">Читать дальше</span></p>\n" +
                "                        </div>\n" +
                "                    </a>\n" +
                "                </div>");
        }
    }
});

var dict = {funny: "Смешные", instructive: "Поучительные", condemning: "Осуждающие"}
$.ajax({
    url: 'json/posts_from_moder.json', // 'http://localhost:3001/main/my_posts',
    method: 'post',
    dataType: 'json',
    data: {id: '3490589089389489', what: "rejected"},
    success: function (data) {
        for (var j = 0; j < data.length; j++) {
            let obj = data[j];
            let arr = [];
            for (var i = 0; i < obj["filters"].length; i++)
                arr.push(dict[obj["filters"][i]])
            $("#content-5").append("<div>\n" +
                "                    <div class=\"post\">\n" +
                "                        <div class=\"post-header\">\n" +
                "                            <span>Автор: Boba</span>\n" +
                "                            <span>Категории: " + obj["categories"].join(', ') + "</span>\n" +
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