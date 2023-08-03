function info(message) {
    var massage_box = document.getElementById("message");
    if (massage_box.style.display == "block")
        return;
    massage_box.firstElementChild.innerHTML = message;
    massage_box.style.display = "block";
    setTimeout(function () {
        massage_box.style.backgroundColor = "#dc3545";
        massage_box.style.color = "white";
        setTimeout(function () {
            massage_box.style.backgroundColor = "transparent";
            massage_box.style.color = "transparent";
            setTimeout(function () {
                massage_box.style.display = "none";
            }, 1100);
        }, 3000);
    }, 100);
}


function check(nickname, password, password_2) {
    if (nickname.length === 0 || password.length === 0) {
        info("Все поля должны быть заполнены");
        return false;
    }
    if (nickname.length < 3 || nickname != nickname.replace(/[^A-Za-z0-9_-]/g,'')) {
        info("Никнейм должен быть длинной не менее 3 символов и<br/>может включать только латинские символы, цифры и символы \"_\" и \"-\"");
        return false;
    }
    if ("" != password.replace(/(?=.*[0-9])(?=.*[!@#$%^&*_-])(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z!@#$%^&*_-]{8,}/g,'')) {
        info("Пароль должен быть длинной не менее 8 символов и должен включать латинские символы<br/>в верхнем и нижнем регистре, цифры и символы из набора \"!@#$%^&*-_\"");
        return false;
    }
    if (password != password_2) {
        info("Пароли не совпадают");
        return false;
    }
    return true;
}

function login() {
    var nickname = document.getElementById("login").children[0].value.trim();
    var password = document.getElementById("login").children[1].value.trim();

    if (check(nickname, password, password)){
        $.ajax({
            url: 'json/new_nickname.json',  // 'http://localhost:3001/main/new_nickname'
            method: 'post',
            dataType: 'json',
            data: {nickname: nickname, password:password},
            success: function (data) {
                console.log("login worked successfully");
            }
        });

    }
}

var ajax = false;

function register() {
    var regist = document.getElementById("register");
    var nickname = regist.children[0].value.trim();
    var password = regist.children[1].value.trim();
    var password_2 = regist.children[2].value.trim();
    var answer = check(nickname, password, password_2);
    if (!answer)
        return answer;
    if (ajax)
        return ajax;

    $.ajax({
        url: 'json/new_nickname.json',  // 'http://localhost:3001/main/new_nickname'
        method: 'get',
        dataType: 'json',
        data: {nickname: nickname},
        success: function (data) {
            if (!data["there_is"]) {
                ajax = true;
                regist.submit();
            }
            else {
                info("Этот никнейм уже занят");
            }
        }
    });
    $.ajax({
        url: 'json/new_nickname.json',  // 'http://localhost:3001/main/new_nickname'
        method: 'post',
        dataType: 'json',
        data: {nickname:nickname, password:password},
        success: function (data) {
            console.log("register worked successfully");
        }
    });
    return false;
}