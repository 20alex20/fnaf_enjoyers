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

document.getElementById('updateData').addEventListener('click', function () {
    var newAvatar = document.getElementById('newAvatar').files[0];
    var newNickname = document.getElementById('newNickname').value.trim();

    if (newNickname) {
        $.ajax({
            url: 'http://localhost:3002/user/exist',  // 'http://localhost:3001/main/new_nickname'
            method: 'get',
            crossDomain: true,
            xhrFields: {
                withCredentials: true
            },
            dataType: 'json',
            data: {nickname: newNickname},
            success: function (data) {
                if (!data["there_is"]) {
                    $.ajax({
                        url: 'http://localhost:3002/user/nickname',  // 'http://localhost:3001/main/update_nickname'
                        method: 'post',
                        crossDomain: true,
                        xhrFields: {
                            withCredentials: true
                        },
                        data: {nickname: newNickname},
                        success: function (data) {
                            info();
                            document.getElementById('nickname').textContent = newNickname;
                            nickname = newNickname;
                            after_nick();
                        }
                    });
                } else {
                    info(true, "Этот никнейм уже занят");
                }

            }
        });
    }

    if (newAvatar) {
        var formData = new FormData();
        formData.append('avatar', newAvatar);
        $.ajax({
            url: 'http://localhost:3002/user/profile-pic',  // 'http://localhost:3001/main/update_avatar'
            method: 'post',
            crossDomain: true,
            xhrFields: {
                withCredentials: true
            },
            data: formData,
            processData: false,
            contentType: false,
            success: function (data) {
                info();
                var reader = new FileReader();
                reader.onload = function (e) {
                    document.getElementById('avatar').style.backgroundImage = 'url(' + e.target.result + ')';
                };
                reader.readAsDataURL(newAvatar);
            }
        });
    }
});


var nickname
$.ajax({
    url: 'http://localhost:3002/user/nickname',  // 'http://localhost:3001/main/get_nickname'
    method: 'get',
    crossDomain: true,
    xhrFields: {
        withCredentials: true
    },
    dataType: 'json',
    success: function (data) {
        nickname = data["nickname"];
        document.getElementById('nickname').textContent = nickname;
        after_nick();
    }
});

$.ajax({
    url: 'http://localhost:3002/user/profile-pic',  // 'http://localhost:3001/main/get_nickname'
    method: 'get',
    crossDomain: true,
    xhrFields: {
        withCredentials: true
    },
    dataType: 'json',
    success: function (data) {
        var avatarUrl = data["img"];
        if (avatarUrl) {
            document.getElementById('avatar').style.backgroundImage = 'url(' + avatarUrl + ')';
        }
    }
});

function CookiesDelete() {
    var cookies = document.cookie.split(";");
    for (var i = 0; i < cookies.length; i++) {
        var cookie = cookies[i];
        var eqPos = cookie.indexOf("=");
        var name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT;";
        document.cookie = name + '=; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
    }
}
$(document).ready(function () {
    // Add a click event listener to the button with the ID "logout"
    $("#logout").click(function () {
        // Send an AJAX POST request to the server
        $.ajax({
            url: 'http://localhost:3002/user/logout',  // 'http://localhost:3001/main/exit'
            method: 'post',
            crossDomain: true,
            xhrFields: {
                withCredentials: true
            },
            success: function (data) {
                console.log("Exit was requested successfully:", data);
                CookiesDelete();
                window.location.href = "authorization.html";
            },
            error: function (xhr, status, error) {
                console.error("Error occurred:", error);
            }
        });
    });
});