drop database if exists "stories.bmstu";

create database "stories.bmstu"
    with
    owner = "admin"
    encoding = 'UTF8'
    lc_collate = 'en_US.utf8'
    lc_ctype = 'en_US.utf8'
    tablespace = pg_default
    connection limit = -1
    is_template = false;
