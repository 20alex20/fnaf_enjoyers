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

$.ajax({
    url: 'json/post.json',
    method: 'get',
    dataType: 'json',
    success: function (data) {
        $("#postsContainer").append('<div><p>' + data["text"].replace(/\n/g, "</p><p>") + '</p>' +
            '<div style="display: none;">' + data["id"] + '</div></div>' +
            '<div><p style="margin-right: auto;">' + data["date_time"] + '</p>' +
            '<img src="images/icons8-удивление-64.png" width="40px" height="40px"/>' +
            '<p>' + data["views"] + '</p>' +
            '<img src="images/icons8-палец-вверх-64.png" width="40px" height="40px"/>' +
            '<p id="likes">' + data["likes"] + '</p>' +
            '<button id="like" class="not_clicked">' +
            '<img src="images/icons8-палец-вверх-64-3.png" width="40px" height="40px"/>' +
            '</button></div>');

        $.ajax({
            url: "json/like.json",
            method: "get",
            dataType: "json",
            success: function (data) {
                var button = document.getElementById("like");
                if (data["state"] === "clicked") {
                    button.classList.add("clicked");
                    button.classList.remove("not_clicked");
                } else {
                    button.addEventListener('click', function () {
                        this.classList.add("clicked");
                        this.classList.remove("not_clicked");
                        $.ajax({
                            url: "json/server_accept.json",
                            method: "post",
                            data: {state: "clicked"},
                            success: function (data) {
                                console.log("Status successfully updated on the server:", data);
                                var likes = document.getElementById("likes");
                                likes.innerHTML = String(parseInt(likes.innerHTML) + 1);
                            },
                            error: function (error) {
                                console.error("Error occurred during the AJAX request:", error);
                            }
                        });
                    });
                }
            },
            error: function (error) {
                console.error("Error occurred during the AJAX request:", error);
            }
        });
    },
    error: function (xhr, status, error) {
        console.error("AJAX Error: " + status + ", " + error);
    }
});

function displayComments(comments, container) {
    comments.forEach(comment => {
        var commentElement = `
                <li class="media">
                    <div class="media-left">
                        <div>
                            <img class="media-object img-rounded" src="${comment["img"]}" width="50px" height="50px">
                        </div>
                    </div>
                    <div class="media-body">
                        <div class="panel panel-primary">
                            <div class="panel-heading">
                                <div class="author">${comment["author"]}</div>
                                <div class="metadata">
                                    <span class="date">${comment["date"]}</span>
                                </div>
                            </div>
                            <div class="panel-body">
                                <div class="media-text text-justify">${comment["text"]}</div>
                            </div>
                            <div class="panel-footer">
                                <textarea class="comment_input comment_input_js" rows="3" placeholder="Введите ваш комментарий"></textarea>
                                <button class="btn btn-primary btn_js">Ответить</button>
                                <div style="display: none;">${comment["id"]}</div>
                            </div>
                        </div>
                    </div>
                </li>
            `;
        container.append(commentElement);


        if (comment["replies"] && comment["replies"].length > 0) {
            var repliesContainer = $('<ul class="media-list"></ul>');
            container.append(repliesContainer);
            displayComments(comment["replies"], repliesContainer);
        }
    });
}

function get_comments() {
// GET запрос для комментариев
    $.ajax({
        url: 'json/comments.json',
        method: 'get',
        dataType: 'json',
        success: function (data) {
            var commentsContainer = $('#first-media-list');
            commentsContainer.empty();
            displayComments(data, commentsContainer);

            var classes = document.getElementsByClassName('btn_js');
            for (var i = 0; i < classes.length; i++) {
                classes[i].addEventListener('click', function () {
                    var element = this.parentElement.firstElementChild;
                    if (element.classList.contains("comment_input_js_2")) {
                        var text = element.value.trim();
                        if (text.length === 0)
                            return;
                        var id_comment = parseInt(this.parentElement.lastElementChild.innerHTML);
                        $.ajax({
                            url: 'json/server_accept.json',
                            method: 'post',
                            data: {text: text, id_comment: id_comment},
                            success: function (data) {
                                info();
                                console.log("Comment posted successfully:", data);
                                get_comments();
                            },
                            error: function (error) {
                                console.error('Error posting comment:', error);
                            }
                        });
                    } else {
                        element.classList.add("comment_input_js_2");
                        setTimeout(function () {
                            element.style.transition = 'revert';
                        }, 800);
                    }
                })
            }
        },
        error: function (error) {
            console.error('Error fetching data:', error);
        }
    });
}

get_comments();

//Анимация главного комментария
document.getElementById('open_comment').addEventListener('click', function () {
    var element_global = this.parentElement.parentElement.lastElementChild.firstElementChild;
    var btn = this.parentElement.parentElement.lastElementChild.lastElementChild;
    if (element_global.classList.contains("comment_input_js_1")) {
        element_global.classList.remove("comment_input_js_1");
        element_global.classList.add("comment_input_js_0");
        btn.classList.add("comment_input_js_btn_0");
        btn.innerHTML = '';
        this.parentElement.parentElement.lastElementChild.style.height = "0";
    } else {
        element_global.classList.remove("comment_input_js_0");
        element_global.classList.add("comment_input_js_1");
        btn.classList.remove("comment_input_js_btn_0");
        btn.innerHTML = 'Отправить';
        this.parentElement.parentElement.lastElementChild.style.height = "130px";
    }
});

document.getElementById("comment_send").addEventListener('click', function () {
    var commentText = this.parentElement.firstElementChild.value.trim();
    var element_text = this.parentElement.firstElementChild;
    if (commentText !== "") {
        var id_post = parseInt(document.getElementById("postsContainer").firstElementChild.lastElementChild.innerHTML);
        $.ajax({
            url: 'json/server_accept.json',
            method: 'post',
            data: {text: commentText, id_post: id_post},
            success: function (data) {
                info();
                console.log("Comment posted successfully:", data);
                get_comments();
                //исчезновение текста из textarea после отправки
                element_text.value = "";
            },
            error: function (error) {
                console.error('Error posting comment:', error);
            }
        });
    }
});
