function info(flag = false, message = "Отправлено") {
    var massage_box = document.getElementById("message");
    if (massage_box.style.display == "block")
        return;
    var mode_flag = flag;
    massage_box.firstElementChild.innerHTML = message;
    massage_box.style.display = "block";
    setTimeout(function () {
        if (!mode_flag)
            massage_box.style.backgroundColor = "#4cae4c";
        else
            massage_box.style.backgroundColor = "#dc3545";
        massage_box.style.color = "white";
        setTimeout(function () {
            massage_box.style.backgroundColor = "transparent";
            massage_box.style.color = "transparent";
            setTimeout(function () {
                massage_box.style.display = "none";
            }, 1100);
        }, 2000);
    }, 100);
}

$(document).ready(function () {
    urlParams = new URLSearchParams(window.location.search);
    if (!urlParams.has("id")) {
        info(true, "Не указан id поста")
    } else {
        document.getElementsByClassName("post_comments")[0].style.display = "block";
        var id_post = parseInt(urlParams.get("id"));
        $.ajax({
            url: 'json/post.json',  // 'http://localhost:3001/main/post'
            method: 'get',
            data: {id_post: id_post},
            dataType: 'json',
            success: function (data) {
                $("#postsContainer").append('<div><p>' + data["text"].replace(/\n/g, "</p><p>") + '</p></div>' +
                    '<div><p style="margin-right: auto;">' + data["date_time"] + '</p>' +
                    '<img src="images/icons8-удивление-64.png" width="40px" height="40px"/>' +
                    '<p>' + data["views"] + '</p>' +
                    '<img src="images/icons8-палец-вверх-64.png" width="40px" height="40px"/>' +
                    '<p id="likes">' + data["likes"] + '</p>' +
                    '<button id="like" class="not_clicked">' +
                    '<img src="images/icons8-палец-вверх-64-3.png" width="40px" height="40px"/>' +
                    '</button></div>');

                $.ajax({
                    url: "json/like.json",  // 'http://localhost:3001/main/like'
                    method: "get",
                    data: {id_post: id_post},
                    dataType: "json",
                    success: function (data) {
                        var button = document.getElementById("like");
                        if (data["state"] === "clicked") {
                            button.classList.add("clicked");
                            button.classList.remove("not_clicked");
                        } else {
                            button.addEventListener('click', function () {
                                if (this.classList.contains("clicked"))
                                    return;
                                this.classList.add("clicked");
                                this.classList.remove("not_clicked");
                                $.ajax({
                                    url: "json/server_accept.json",  // 'http://localhost:3001/main/set_like'
                                    method: "post",
                                    data: {id_post: id_post},
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
                url: 'json/comments.json',  // 'http://localhost:3001/main/comments'
                method: 'get',
                data: {id_post: id_post},
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
                                    url: 'json/server_accept.json',  // 'http://localhost:3001/main/comment_comment'
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
                $.ajax({
                    url: 'json/server_accept.json',  // 'http://localhost:3001/main/post_comment'
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
    }
});
