function setButtonClass(data) {
    var button = document.getElementById("like");
    if (data["state"] === "clicked") {
        button.classList.add("clicked");
        button.classList.remove("not_clicked");
    } else {
        button.classList.add("not_clicked");
        button.classList.remove("clicked");
    }
}

function sendButtonState(state) {
    $.ajax({
        type: "POST",
        url: "json/like.json",
        data: { state: state },
        success: function (data) {
            console.log("Status successfully updated on the server:", data);
        },
        error: function (error) {
            console.error("Error occurred during the AJAX request:", error);
        }
    });
}

$.ajax({
    url: 'json/post.json',
    method: 'get',
    dataType: 'json',
    success: function (data) {
        $("#postsContainer").append('<div><p>' + data["text"] + '</p></div>' +
            '<div><p style="margin-right: auto;">' + data["date_time"] + '</p>' +
            '<img src="images/icons8-удивление-64.png" width="40px" height="40px"/>' +
            '<p>' + data["views"] + '</p>' +
            '<img src="images/icons8-палец-вверх-64.png" width="40px" height="40px"/>' +
            '<p id="likes">' + data["likes"] + '</p>' +
            '<button id="like" class="not_clicked">' +
            '<img src="images/icons8-палец-вверх-64-3.png" width="40px" height="40px"/>' +
            '</button></div>');

        $.ajax({
            type: "GET",
            url: "json/like.json",
            dataType: "json",
            success: function (data) {
                setButtonClass(data);
                document.getElementById('like').addEventListener('click', function () {
                    var element = document.getElementById('like');
                    if (element.classList.contains("not_clicked")) {
                        element.classList.remove("not_clicked");
                        element.classList.add("clicked");
                        sendButtonState("clicked");
                    }
                });
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
        const commentElement = `
                <li class="media">
                    <div class="media-left">
                        <a href="#">
                            <img class="media-object img-rounded" src="${comment["img"]}" alt="...">
                        </a>
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
                            </div>
                        </div>
                    </div>
                </li>
            `;
        container.append(commentElement);


        if (comment["replies"] && comment["replies"].length > 0) {
            const repliesContainer = $('<ul class="media-list"></ul>');
            container.append(repliesContainer);
            // Recursive call to display nested comments (replies)
            displayComments(comment["replies"], repliesContainer);
        }
    });
    var classes = document.getElementsByClassName('btn_js')
    for (var i = 0; i < classes.length; i++) {
        classes[i].addEventListener('click', function () {
            element_global = this.parentElement.firstElementChild;
            element_global.classList.add("comment_input_js_2");
            setTimeout(function () {
                element_global.style.transition = 'revert';
            }, 800);
        })
    }
}

// GET запрос для комментариев
$.ajax({
    url: 'json/comments.json',
    method: 'GET',
    dataType: 'json',
    success: function (data) {
        const commentsContainer = $('.media-list');
        displayComments(data, commentsContainer);
    },
    error: function (error) {
        console.error('Error fetching data:', error);
    }
});


//POST запрос с содержимым комментариев(нужно добавить отправку запроса после отправки каждого ответа на комментарий)
document.addEventListener('DOMContentLoaded', function () {
    var element_global;

    //Анимация главного комментария
    document.getElementById('open_comment').addEventListener('click', function () {
        element_global = this.parentElement.parentElement.lastElementChild.firstElementChild;
        let btn = this.parentElement.parentElement.lastElementChild.lastElementChild;
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

    function sendCommentWithReplies(commentData) {
        $.ajax({
            url: 'json/comments.json',
            method: 'POST',
            data: commentData,
            dataType: 'json',
            success: function (response) {
                console.log("Comment posted successfully:", response);
            },
            error: function (error) {
                console.error('Error posting comment:', error);
            }
        });
    }

//Функция для преобразования даты и времени
    function getFormattedDateTime() {
        const currentDate = new Date();
        const day = String(currentDate.getDate()).padStart(2, '0');
        const month = String(currentDate.getMonth() + 1).padStart(2, '0');
        const year = currentDate.getFullYear();
        const hours = String(currentDate.getHours()).padStart(2, '0');
        const minutes = String(currentDate.getMinutes()).padStart(2, '0');
        return `${day}.${month}.${year}, ${hours}:${minutes}`;
    }

    document.querySelector('.comment_input_js_btn_0').addEventListener('click', function () {
        const commentText = element_global.value.trim();
        if (commentText !== "") {
            const commentData = {
                text: commentText,
                author: "user id",
                datetime: getFormattedDateTime(),
                replies: []
            };

            //Ответ по умолчанию
            const replyText = "This is a reply.";
            const replyAuthor = "user id";
            const replyDatetime = getFormattedDateTime();
            const replyData = {
                text: replyText,
                author: replyAuthor,
                datetime: replyDatetime
            };

            commentData["replies"].push(replyData);

            sendCommentWithReplies(commentData);

            //исчезновение текста из textarea после отправки
            element_global.value = "";
        }
    });

});

