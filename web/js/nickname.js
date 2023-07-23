function getAvatarAndNickname() {
    $.ajax({
        url: 'json/nickname.json',
        method: 'GET',
        dataType: 'json',
        success: function (data) {
            var nickname = data["nickname"]
            if (nickname) {
                document.getElementById('nickname').textContent = nickname;
            }
        },
        error: function (xhr, status, error) {
            console.error('Error fetching nickname:', status, error);
        }
    });

    $.ajax({
        url: 'json/img.json',
        method: 'GET',
        dataType: 'json',
        success: function (data) {
            var avatarUrl = data.img;
            if (avatarUrl) {
                document.getElementById('avatar').style.backgroundImage = 'url(' + avatarUrl + ')';
            }
        },
        error: function (xhr, status, error) {
            console.error('Error fetching avatar image:', status, error);
        }
    });

}

getAvatarAndNickname();

document.getElementById('updateData').addEventListener('click', function () {
    var newAvatar = document.getElementById('newAvatar').files[0];
    var newNickname = document.getElementById('newNickname').value;

    if (newAvatar) {
        var reader = new FileReader();
        reader.onload = function (e) {
            document.getElementById('avatar').style.backgroundImage = 'url(' + e.target.result + ')';
        };
        reader.readAsDataURL(newAvatar);
    }

    if (newNickname) {
        $.ajax({
            url: 'json/nickname.json',
            method: 'POST',
            dataType: 'json',
            data: { nickname: newNickname },
            success: function (data) {
                console.log('Nickname updated successfully:', data);
                document.getElementById('nickname').textContent = newNickname;
            },
            error: function (xhr, status, error) {
                console.error('Error updating nickname:', status, error);
            }
        });
    }

    if (newAvatar) {
        var formData = new FormData();
        formData.append('avatar', newAvatar);
        $.ajax({
            url: 'json/nickname.json',
            method: 'POST',
            data: formData,
            processData: false,
            contentType: false,
            success: function (data) {
                console.log('Avatar updated successfully:', data);
            },
            error: function (xhr, status, error) {
                console.error('Error updating avatar:', status, error);
            }
        });
    }
});

