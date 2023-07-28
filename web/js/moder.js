function info() {
    if (document.getElementById("message").style.display == "block")
        return;
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
}

function after_nick() {

}

var dict = {funny: "Смешные", instructive: "Поучительные", condemning: "Осуждающие"}
var dict2 = {IU: "ИУ", IBM: "ИБМ", SM: "СМ", E: "Э", MT: "МТ", RL: "РЛ", BMT: "БМТ", RK: "РК", FN: "ФН", L: "Л",
    SGN: "СГН", UR: "ЮР"}
$.ajax({
    url: 'json/posts_for_moder.json',  // 'http://localhost:3001/moder/posts',
    method: 'get',
    dataType: 'json',
    success: function (data) {
        for (var j = 0; j < data.length; j++) {
            var obj = data[j];
            var arr = [];
            for (var i = 0; i < obj["filters"].length; i++)
                arr.push(dict[obj["filters"][i]])
            var arr2 = [];
            for (i = 0; i < obj["categories"].length; i++)
                if (obj["categories"][i].includes("-"))
                    arr2.push(dict2[obj["categories"][i].split("-")[0]] + '-' + obj["categories"][i].split("-")[1]);
                else
                    arr2.push(dict2[obj["categories"][i]])
            $("#content-1").append("<div>\n" +
                "                    <div class=\"post\">\n" +
                "                        <div class=\"post-header\">\n" +
                "                            <span>Автор: " + obj["nickname"] + "</span>\n" +
                "                            <span>Категории: " + arr2.join(', ') + "</span>\n" +
                "                        </div>\n" +
                "                        <div class=\"post-header\">\n" +
                "                            <span>Время: " + obj["date_time"] + "</span>\n" +
                "                            <span>Подкатегории: " + arr.join(', ') + "</span>\n" +
                "                        </div>\n" +
                "                        <p style=\"font-size: 14px;\">" + obj["text"] + "</p>\n" +
                "                        <div style=\"justify-content: center;\">\n" +
                "                            <button class=\"accept-btn\">Принять</button>\n" +
                "                            <button class=\"reject-btn\">Отклонить</button>\n" +
                "                            <div class=\"reason-box\">\n" +
                "                                <textarea class=\"reason-form\" rows=\"3\" placeholder=\"Причина отказа\"></textarea>\n" +
                "                                <button class=\"cancel-reject-btn\">Отмена</button>\n" +
                "                                <button class=\"confirm-reject-btn\">Удалить окончательно</button>\n" +
                "                            </div>\n" +
                "                            <div style='display: none;'>" + String(obj["id"]) + "</div>\n" +
                "                        </div>\n" +
                "                    </div>\n" +
                "                </div>");
        }

        var elems = document.getElementsByClassName("accept-btn");
        for (i = 0; i < data.length; i++) {
            elems[i].addEventListener("click", function () {
                var id = parseInt(this.parentElement.lastElementChild.innerHTML);
                var element_parent = this;
                $.ajax({
                    url: 'json/server_accept.json',  // 'http://localhost:3001/moder/posts_verified'
                    method: 'post',
                    data: {id_post: id, accept: true},
                    success: function (data) {
                        info();
                        var parent = document.getElementById("content-1");
                        parent.removeChild(element_parent.parentElement.parentElement.parentElement);
                    }
                });
            });
        }

        elems = document.getElementsByClassName("confirm-reject-btn");
        for (i = 0; i < data.length; i++) {
            elems[i].addEventListener("click", function () {
                var id_post = parseInt(this.parentElement.parentElement.lastElementChild.innerHTML);
                var text = this.parentElement.firstElementChild.value.trim();
                var element_parent_2 = this;
                if (text.length === 0)
                    return;
                $.ajax({
                    url: 'json/server_accept.json',  // 'http://localhost:3001/moder/posts_verified'
                    method: 'post',
                    data: {id_post: id_post, accept: true, text: text},
                    success: function (data) {
                        info();
                        var parent = document.getElementById("content-1");
                        parent.removeChild(element_parent_2.parentElement.parentElement.parentElement.parentElement);
                    }
                });
            });
        }

        const posts = document.querySelectorAll('.post');
        posts.forEach((post) => {
            const rejectBtn = post.querySelector('.reject-btn');
            const acceptBtn = post.querySelector('.accept-btn');
            const reasonBox = post.querySelector('.reason-box');
            const confirmRejectBtn = post.querySelector('.confirm-reject-btn');

            rejectBtn.addEventListener('click', () => {
                reasonBox.style.display = 'block';
                rejectBtn.style.display = 'none';
                acceptBtn.style.display = 'none';
            });

            const cancelRejectBtn = post.querySelector('.cancel-reject-btn');
            cancelRejectBtn.addEventListener('click', () => {
                reasonBox.style.display = 'none';
                rejectBtn.style.display = 'block';
                acceptBtn.style.display = 'block';
            });
        });
    }
});