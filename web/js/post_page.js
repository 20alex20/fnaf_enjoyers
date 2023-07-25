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

        // Check if there are replies for the comment
        if (comment["replies"] && comment["replies"].length > 0) {
            const repliesContainer = $('<ul class="media-list"></ul>');
            container.append(repliesContainer);
            // Recursive call to display nested comments (replies)
            displayComments(comment["replies"], repliesContainer);
        }
    });
}

// Function to make the AJAX GET request
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



