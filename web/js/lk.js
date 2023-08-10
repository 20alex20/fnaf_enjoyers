var status
$.ajax({
    url: 'json/user_status.json',  // 'http://localhost:3001/main/get_nickname'
    method: 'get',
    dataType: 'json',
    success: function (data) {
        status = data["status"];
        if (status == "user") {
            $("#user-block").append("<div class=\"tabs\">\n" +
            "            <input type=\"radio\" name=\"tab-btn\" id=\"tab-btn-1\" value=\"\" checked>\n" +
            "            <label for=\"tab-btn-1\">Понравившиеся</label>\n" +
            "            <input type=\"radio\" name=\"tab-btn\" id=\"tab-btn-2\" value=\"\">\n" +
            "            <label for=\"tab-btn-2\">Мои посты</label>\n" +
            "            <input type=\"radio\" name=\"tab-btn\" id=\"tab-btn-5\" value=\"\">\n" +
            "            <label for=\"tab-btn-5\">Отклоненные</label>\n" +
            "            <input type=\"radio\" name=\"tab-btn\" id=\"tab-btn-3\" value=\"\">\n" +
            "            <label for=\"tab-btn-3\">Создать пост</label>\n" +
            "            <input type=\"radio\" name=\"tab-btn\" id=\"tab-btn-4\" value=\"\">\n" +
            "            <label for=\"tab-btn-4\">Мои данные</label>\n" +
            "\n" +
            "            <div class=\"read\" id=\"content-1\">\n" +
            "\n" +
            "            </div>\n" +
            "            <div class=\"read\" id=\"content-2\">\n" +
            "\n" +
            "            </div>\n" +
            "            <div class=\"read\" id=\"content-5\">\n" +
            "\n" +
            "\n" +
            "            </div>\n" +
            "            <div class=\"comment\" id=\"content-3\">\n" +
            "                <div id=\"categories\">\n" +
            "                    <div>\n" +
            "                        <p>Выберите факультеты:</p>\n" +
            "                        <select id=\"faculties\" multiple size=\"4\">\n" +
            "                            <option value=\"IU\">ИУ</option>\n" +
            "                            <option value=\"IBM\">ИБМ</option>\n" +
            "                            <option value=\"SM\">СМ</option>\n" +
            "                            <option value=\"E\">Э</option>\n" +
            "                            <option value=\"MT\">МТ</option>\n" +
            "                            <option value=\"RL\">РЛ</option>\n" +
            "                            <option value=\"BMT\">БМТ</option>\n" +
            "                            <option value=\"RK\">РК</option>\n" +
            "                            <option value=\"FN\">ФН</option>\n" +
            "                            <option value=\"L\">Л</option>\n" +
            "                            <option value=\"SGN\">СГН</option>\n" +
            "                            <option value=\"UR\">ЮР</option>\n" +
            "                        </select>\n" +
            "                    </div>\n" +
            "                    <div id=\"selects\">\n" +
            "                        <p style=\"display: none;\">Выберите кафедру: </p>\n" +
            "                        <select id=\"IU\" multiple size=\"4\">\n" +
            "                            <option value=\"IU-1\">ИУ-1</option>\n" +
            "                            <option value=\"IU-2\">ИУ-2</option>\n" +
            "                            <option value=\"IU-3\">ИУ-3</option>\n" +
            "                            <option value=\"IU-4\">ИУ-4</option>\n" +
            "                            <option value=\"IU-5\">ИУ-5</option>\n" +
            "                            <option value=\"IU-6\">ИУ-6</option>\n" +
            "                            <option value=\"IU-7\">ИУ-7</option>\n" +
            "                            <option value=\"IU-8\">ИУ-8</option>\n" +
            "                            <option value=\"IU-9\">ИУ-9</option>\n" +
            "                            <option value=\"IU-10\">ИУ-10</option>\n" +
            "                            <option value=\"IU-11\">ИУ-11</option>\n" +
            "                            <option value=\"IU-12\">ИУ-12</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"IBM\" multiple size=\"4\">\n" +
            "                            <option value=\"IBM-1\">ИБМ-1</option>\n" +
            "                            <option value=\"IBM-2\">ИБМ-2</option>\n" +
            "                            <option value=\"IBM-3\">ИБМ-3</option>\n" +
            "                            <option value=\"IBM-4\">ИБМ-4</option>\n" +
            "                            <option value=\"IBM-5\">ИБМ-5</option>\n" +
            "                            <option value=\"IBM-6\">ИБМ-6</option>\n" +
            "                            <option value=\"IBM-7\">ИБМ-7</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"SM\" multiple size=\"4\">\n" +
            "                            <option value=\"SM-1\">СМ-1</option>\n" +
            "                            <option value=\"SM-2\">СМ-2</option>\n" +
            "                            <option value=\"SM-3\">СМ-3</option>\n" +
            "                            <option value=\"SM-4\">СМ-4</option>\n" +
            "                            <option value=\"SM-5\">СМ-5</option>\n" +
            "                            <option value=\"SM-6\">СМ-6</option>\n" +
            "                            <option value=\"SM-7\">СМ-7</option>\n" +
            "                            <option value=\"SM-8\">СМ-8</option>\n" +
            "                            <option value=\"SM-9\">СМ-9</option>\n" +
            "                            <option value=\"SM-10\">СМ-10</option>\n" +
            "                            <option value=\"SM-11\">СМ-11</option>\n" +
            "                            <option value=\"SM-12\">СМ-12</option>\n" +
            "                            <option value=\"SM-13\">СМ-13</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"E\" multiple size=\"4\">\n" +
            "                            <option value=\"E-1\">Э-1</option>\n" +
            "                            <option value=\"E-2\">Э-2</option>\n" +
            "                            <option value=\"E-3\">Э-3</option>\n" +
            "                            <option value=\"E-4\">Э-4</option>\n" +
            "                            <option value=\"E-5\">Э-5</option>\n" +
            "                            <option value=\"E-6\">Э-6</option>\n" +
            "                            <option value=\"E-7\">Э-7</option>\n" +
            "                            <option value=\"E-8\">Э-8</option>\n" +
            "                            <option value=\"E-9\">Э-9</option>\n" +
            "                            <option value=\"E-10\">Э-10</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"MT\" multiple size=\"4\">\n" +
            "                            <option value=\"MT-1\">МТ-1</option>\n" +
            "                            <option value=\"MT-2\">МТ-2</option>\n" +
            "                            <option value=\"MT-3\">МТ-3</option>\n" +
            "                            <option value=\"MT-4\">МТ-4</option>\n" +
            "                            <option value=\"MT-5\">МТ-5</option>\n" +
            "                            <option value=\"MT-6\">МТ-6</option>\n" +
            "                            <option value=\"MT-7\">МТ-7</option>\n" +
            "                            <option value=\"MT-8\">МТ-8</option>\n" +
            "                            <option value=\"MT-9\">МТ-9</option>\n" +
            "                            <option value=\"MT-10\">МТ-10</option>\n" +
            "                            <option value=\"MT-11\">МТ-11</option>\n" +
            "                            <option value=\"MT-12\">МТ-12</option>\n" +
            "                            <option value=\"MT-13\">МТ-13</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"RL\" multiple size=\"4\">\n" +
            "                            <option value=\"RL-1\">РЛ-1</option>\n" +
            "                            <option value=\"RL-2\">РЛ-2</option>\n" +
            "                            <option value=\"RL-3\">РЛ-3</option>\n" +
            "                            <option value=\"RL-4\">РЛ-4</option>\n" +
            "                            <option value=\"RL-5\">РЛ-5</option>\n" +
            "                            <option value=\"RL-6\">РЛ-6</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"BMT\" multiple size=\"4\">\n" +
            "                            <option value=\"BMT-1\">БМТ-1</option>\n" +
            "                            <option value=\"BMT-2\">БМТ-2</option>\n" +
            "                            <option value=\"BMT-3\">БМТ-3</option>\n" +
            "                            <option value=\"BMT-4\">БМТ-4</option>\n" +
            "                            <option value=\"BMT-5\">БМТ-5</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"RK\" multiple size=\"4\">\n" +
            "                            <option value=\"RK-1\">РК-1</option>\n" +
            "                            <option value=\"RK-2\">РК-2</option>\n" +
            "                            <option value=\"RK-3\">ИУ-3</option>\n" +
            "                            <option value=\"RK-4\">ИУ-4</option>\n" +
            "                            <option value=\"RK-5\">ИУ-5</option>\n" +
            "                            <option value=\"RK-6\">ИУ-6</option>\n" +
            "                            <option value=\"RK-9\">ИУ-9</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"FN\" multiple size=\"4\">\n" +
            "                            <option value=\"FN-1\">ФН-1</option>\n" +
            "                            <option value=\"FN-2\">ФН-2</option>\n" +
            "                            <option value=\"FN-3\">ФН-3</option>\n" +
            "                            <option value=\"FN-4\">ФН-4</option>\n" +
            "                            <option value=\"FN-5\">ФН-5</option>\n" +
            "                            <option value=\"FN-6\">ФН-6</option>\n" +
            "                            <option value=\"FN-7\">ФН-7</option>\n" +
            "                            <option value=\"FN-11\">ФН-11</option>\n" +
            "                            <option value=\"FN-12\">ФН-12</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"L\" multiple size=\"4\">\n" +
            "                            <option value=\"L-1\">Л-1</option>\n" +
            "                            <option value=\"L-2\">Л-2</option>\n" +
            "                            <option value=\"L-3\">Л-3</option>\n" +
            "                            <option value=\"L-4\">Л-4</option>\n" +
            "                        </select>\n" +
            "                        <select id=\"SGN\" multiple size=\"4\">\n" +
            "                            <option value=\"SGN-1\">СГН-1</option>\n" +
            "                            <option value=\"SGN-2\">СГН-2</option>\n" +
            "                            <option value=\"SGN-3\">СГН-3</option>\n" +
            "                            <option value=\"SGN-4\">СГН-4</option>\n" +
            "                        </select>\n" +
            "                    </div>\n" +
            "                </div>\n" +
            "                <div class=\"filters2\">\n" +
            "                    <p>Выберите подкатегории:</p>\n" +
            "                    <label><input type=\"checkbox\" id=\"filter1\" value=\"funny\">\n" +
            "                        <p>Смешные</p></label>\n" +
            "                    <label><input type=\"checkbox\" id=\"filter2\" value=\"instructive\">\n" +
            "                        <p>Поучительные</p></label>\n" +
            "                    <label><input type=\"checkbox\" id=\"filter3\" value=\"condemning\">\n" +
            "                        <p>Осуждающие</p></label>\n" +
            "                </div>\n" +
            "                <textarea id=\"text\" class=\"comment_input\" rows=\"6\" placeholder=\"Введите ваш комментарий\"></textarea>\n" +
            "                <button id=\"btn\" class=\"btn btn-primary\">Отправить</button>\n" +
            "            </div>\n" +
            "            <script type=\"text/javascript\" src=\"js/personal_area.js\"></script>\n" +
            "            <script>\n" +
            "                document.getElementById(\"faculties\").addEventListener(\"change\", function () {\n" +
            "                    var selects = document.getElementById(\"selects\").children;\n" +
            "                    var faculties = this.children;\n" +
            "                    var flag = false;\n" +
            "                    for (var i = 0; i < 11; i++) {\n" +
            "                        if (faculties[i].selected) {\n" +
            "                            selects[i + 1].style.display = \"block\";\n" +
            "                            for (var j = 0; j < selects[i + 1].children.length; j++)\n" +
            "                                selects[i + 1].children[j].selected = true;\n" +
            "                            flag = true;\n" +
            "                        } else\n" +
            "                            selects[i + 1].style.display = \"none\";\n" +
            "                    }\n" +
            "                    if (flag)\n" +
            "                        selects[0].style.display = \"block\";\n" +
            "                    else\n" +
            "                        selects[0].style.display = \"none\";\n" +
            "                })\n" +
            "            </script>\n" +
            "            \n" +
            "            <div class=\"about_me\" id=\"content-4\">\n" +
            "                <div class=\"avatar\" id=\"avatar\" style=\"background-image: url('images/icons8-аватар-100.png')\"></div>\n" +
            "                <div class=\"nickname\" id=\"nickname\">JohnDoe</div>\n" +
            "                <div class=\"level\" id=\"level\">&#9734 Уровень: 1205</div>\n" +
            "\n" +
            "                <div class=\"edit-form\">\n" +
            "                    <div>\n" +
            "                        <div>\n" +
            "                            <div><label for=\"newAvatar\" style=\"font-size: 18px\">Сменить аватарку:</label></div>\n" +
            "                            <div><input type=\"file\" style=\"font-size: 18px; font-weight: 100;\" id=\"newAvatar\"\n" +
            "                                        accept=\"image/*\"></div>\n" +
            "                        </div>\n" +
            "                        <div>\n" +
            "                            <div><label for=\"newNickname\" style=\"font-size: 18px\">Сменить ник:</label></div>\n" +
            "                            <div><input type=\"text\" id=\"newNickname\"></div>\n" +
            "                        </div>\n" +
            "                    </div>\n" +
            "                    <button id=\"updateData\" class=\"btn btn-primary\" style=\"margin:0 0 0 auto;\">Сохранить изменения</button>\n" +
            "                    <button id=\"logout\" class=\"btn btn-primary2\" style=\"border: 1px solid;margin-left: 7px;\">Выйти</button>\n" +
            "                </div>\n" +
            "\n" +
            "\n" +
            "            </div>\n" +
            "            <script type=\"text/javascript\" src=\"js/nickname.js\"></script>\n" +
            "        </div>")

        }
    else{
            $("#user-block").append("        <div class=\"tabs\">\n" +
                "            <input type=\"radio\" name=\"tab-btn\" id=\"tab-btn-1\" value=\"\" checked>\n" +
                "            <label for=\"tab-btn-1\">Новые посты</label>\n" +
                "            <input type=\"radio\" name=\"tab-btn\" id=\"tab-btn-4\" value=\"\">\n" +
                "            <label for=\"tab-btn-4\">Мои данные</label>\n" +
                "\n" +
                "            <div class=\"read\" id=\"content-1\">\n" +
                "\n" +
                "            </div>\n" +
                "            <script type=\"text/javascript\" src=\"js/moder.js\"></script>\n" +
                "            <div class=\"about_me\" id=\"content-4\">\n" +
                "                <div class=\"avatar\" id=\"avatar\" style=\"background-image: url('images/icons8-аватар-100.png')\"></div>\n" +
                "                <div class=\"nickname\" id=\"nickname\">JohnDoe</div>\n" +
                "                <div class=\"level\" id=\"level\">&#9734 Уровень: 1205</div>\n" +
                "\n" +
                "                <div class=\"edit-form\">\n" +
                "                    <div>\n" +
                "                        <div>\n" +
                "                            <div><label for=\"newAvatar\" style=\"font-size: 18px\">Сменить аватарку:</label></div>\n" +
                "                            <div><input type=\"file\" style=\"font-size: 18px; font-weight: 100;\" id=\"newAvatar\"\n" +
                "                                        accept=\"image/*\"></div>\n" +
                "                        </div>\n" +
                "                        <div>\n" +
                "                            <div><label for=\"newNickname\" style=\"font-size: 18px\">Сменить ник:</label></div>\n" +
                "                            <div><input type=\"text\" id=\"newNickname\"></div>\n" +
                "                        </div>\n" +
                "                    </div>\n" +
                "                    <button id=\"updateData\" class=\"btn btn-primary\" style=\"margin:0 0 0 auto;\">Сохранить изменения\n" +
                "                    <button id=\"logout\" class=\"btn btn-primary2\" style=\"border: 1px solid;margin-left: 7px;\">Выйти</button>\n" +
                "                    </button>\n" +
                "                </div>\n" +
                "\n" +
                "\n" +
                "            </div>\n" +
                "        </div>\n" +
                "        <script type=\"text/javascript\" src=\"js/nickname.js\"></script>\n" +
                "    </div>")

        }
    },
    error: function (xhr, textStatus, errorThrown) {
        console.error('AJAX Error:', textStatus, errorThrown);
    }
});
