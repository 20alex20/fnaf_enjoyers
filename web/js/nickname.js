function info(flag=false, message="Отправлено") {
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
            url: 'json/new_nickname.json',  // 'http://localhost:3001/main/new_nickname'
            method: 'get',
            dataType: 'json',
            success: function (data) {
                if (!data["there_is"]) {
                    $.ajax({
                        url: 'json/server_accept.json',  // 'http://localhost:3001/main/update_nickname'
                        method: 'post',
                        data: {nickname: newNickname},
                        success: function (data) {
                            info();
                            document.getElementById('nickname').textContent = newNickname;
                            nickname = newNickname;
                            after_nick();
                        }
                    });
                }
                else {
                    info(true, "Этот никнейм уже занят");
                }

            }
        });
    }

    if (newAvatar) {
        var formData = new FormData();
        formData.append('avatar', newAvatar);
        $.ajax({
            url: 'json/server_accept.json',  // 'http://localhost:3001/main/update_avatar'
            method: 'post',
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
    url: 'json/nickname.json',  // 'http://localhost:3001/main/get_nickname'
    method: 'get',
    dataType: 'json',
    success: function (data) {
        nickname = data["nickname"];
        document.getElementById('nickname').textContent = nickname;
        after_nick();
    }
});

$.ajax({
    url: 'json/img.json',  // 'http://localhost:3001/main/get_nickname'
    method: 'get',
    dataType: 'json',
    success: function (data) {
        var avatarUrl = data["img"];
        if (avatarUrl) {
            document.getElementById('avatar').style.backgroundImage = 'url(' + avatarUrl + ')';
        }
    }
});

