function truncate(str, maxlength) {
    if (str.length > maxlength) {
        return str.slice(0, maxlength - 3) + '...';
    }
    return str
}

dict = {"date_time": "card", "likes": "logo", "views": "app"}
array = ["date_time", "likes", "views"]
var my_index = 0
for (var i = 0; i < 3; i++) {
    var key = array[i]
    $.ajax({
        url: 'json/posts_' + key + '.json',         /* Куда пойдет запрос */
        method: 'get',             /* Метод передачи (post или get) */
        dataType: 'json',          /* Тип данных в ответе (xml, json, script, html). */
        data: {category: 'all', filter: "without", sort: key, number: 2, page: 1},     /* Параметры передаваемые в запросе. */
        success: function (data) {   /* функция, которая будет выполнена после успешного запроса.  */
            for (var j = 0; j < data.length; j++) {
                var obj = data[j];
                $("#portfoliolist").append("<div class='portfolio " + dict[array[my_index]] + " mix_all' data-cat='" + dict[array[my_index]] + "' style=' display: inline-block; opacity: 1;'>\n" +
                    "                    <div class='portfolio-wrapper'>\n" +
                    "                        <a href='#small-dialog' class='b-link-stripe b-animate-go  thickbox '>\n" +
                    "                            <div class='post'>\n" +
                    "                                <p>" + truncate(obj["text"], 500) + "</p>\n" +
                    "                                <div><p style='margin-right: auto;'>" + obj["date_time"] + "</p>\n" +
                    "                                    <img src='images/icons8-удивление-64.png' width='40px' height='40px'/>\n" +
                    "                                    <p>" + String(obj["views"]) + "</p>\n" +
                    "                                    <img src='images/icons8-палец-вверх-64.png' width='40px' height='40px'/>\n" +
                    "                                    <p>" + String(obj["likes"]) + "</p>\n" +
                    "                                </div>\n" +
                    "                            </div>\n" +
                    "                            <div class='b-wrapper'><h2 class='b-animate b-from-left    b-delay03 '><img\n" +
                    "                                    src='images/icon-eye.png' alt='' style='top: 0; margin-top: 5%;'></h2>\n" +
                    "                                <p class='b-animate b-from-right  b-delay03 '><span class='m_4'>Читать дальше</span></p>\n" +
                    "                            </div>\n" +
                    "                        </a>\n" +
                    "                    </div>\n" +
                    "                </div>");
            }
            my_index += 1;
        }
    });
}