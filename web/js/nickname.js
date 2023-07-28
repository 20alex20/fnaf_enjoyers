document.getElementById('updateData').addEventListener('click', function () {
    var newAvatar = document.getElementById('newAvatar').files[0];
    var newNickname = document.getElementById('newNickname').value.trim();

    if (newNickname) {
        $.ajax({
            url: 'json/server_accept.json',  // 'http://localhost:3001/main/update_nickname'
            method: 'post',
            data: {nickname: newNickname },
            success: function (data) {
                info();
                document.getElementById('nickname').textContent = newNickname;
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
        nickname = data["nickname"]
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

