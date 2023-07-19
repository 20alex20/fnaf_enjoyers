-- Database: stories.bmstu

-- drop database if exists "stories.bmstu";
--
-- create database "stories.bmstu"
--     with
--     owner = "admin"
--     encoding = 'UTF8'
--     lc_collate = 'en_US.utf8'
--     lc_ctype = 'en_US.utf8'
--     tablespace = pg_default
--     connection limit = -1
--     is_template = false;

create extension if not exists "uuid-ossp";

drop table if exists "user" cascade;
create table if not exists "user" (
                                      id              uuid        default uuid_generate_v4()
    ,   login           text      unique not null
    ,   password        text        not null
    ,   "name"          varchar(50) not null
    ,   rating          numeric(5,2) not null default 0.0
    ,   is_moderator    boolean not null default false

    ,   primary key (id)
);

drop table if exists category cascade;
create table if not exists category (
                                        id int generated always as identity
    ,   title varchar(50) unique not null

    ,   primary key (id)
);

drop table if exists post cascade;
create table if not exists post (
                                    id uuid default uuid_generate_v4()
    ,   user_id uuid references "user" (id)
    ,   "date" date not null default current_date
    ,   "text" text not null
    ,   views int not null default 0
    ,   likes int not null default 0
    ,   comments int not null default 0

    ,   primary key (id)
);

drop table if exists post_category;
create table if not exists post_category (
                                             post_id uuid references post (id)
    ,   category_id int references category (id)

    ,   constraint post_category_pkey
                                                 primary key (post_id, category_id)
);

drop table if exists "comment" cascade;
create table if not exists "comment" (
                                         id uuid default uuid_generate_v4()
    ,   post_id uuid references post (id)
    ,   user_id uuid references "user" (id)
    ,   "date" date not null default current_date
    ,   "text" text not null
    ,   likes int not null default 0

    ,   primary key (id)
);

drop table if exists achievement cascade;
create table if not exists achievement (
                                           id int generated always as identity
    ,   title varchar(50) not null
    ,   description varchar(50) not null

    ,   primary key (id)
);

drop table if exists user_achievement;
create table if not exists user_achievement (
                                                user_id uuid references "user" (id)
    ,   achievement_id int references achievement (id)

    ,   constraint user_achievement_pkey
                                                    primary key (user_id, achievement_id)
);
