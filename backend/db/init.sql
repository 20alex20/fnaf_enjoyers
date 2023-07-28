create extension if not exists "uuid-ossp";
create extension if not exists pgcrypto;

drop table if exists profile_picture cascade;
create table if not exists profile_picture
(
    id    uuid not null default uuid_generate_v4(),
    link  text not null,
    primary key (id)
);

create or replace function default_picture()
    returns uuid
    language sql as
$$
select id
from profile_picture
where link = 'default.png'
$$;

drop table if exists "info";
create table if not exists "info"
(
    total_posts int not null default 0,
    total_likes int not null default 0,
    total_views int not null default 0
);

drop table if exists "user" cascade;
create table if not exists "user"
(
    id                 uuid        not null default uuid_generate_v4(),
    profile_picture_id uuid        not null default default_picture() references profile_picture (id),
    nickname           text unique not null,
    password           text        not null,
    is_moderator       boolean     not null default false,
    user_posts         int         not null default 0,
    user_likes         int         not null default 0,
    user_views         int         not null default 0,

    primary key (id)
);

drop table if exists category cascade;
create table if not exists category
(
    id    int generated always as identity,
    title varchar(50) unique not null,
    primary key (id)
);

drop table if exists "filter" cascade;
create table if not exists "filter"
(
    id    int generated always as identity,
    title varchar(50) unique not null,
    primary key (id)
);

drop table if exists post cascade;
create table if not exists post
(
    id       uuid      not null default uuid_generate_v4(),
    user_id  uuid references "user" (id),
    "date"   timestamp not null default date_trunc('minute', now()),
    "text"   text      not null,
    views    int       not null default 0,
    likes    int       not null default 0,
    accepted boolean   not null default false,
    primary key (id)
);

drop table if exists post_category;
create table if not exists post_category
(
    post_id     uuid references post (id),
    category_id int references category (id),
    constraint post_category_pkey
        primary key (post_id, category_id)
);

drop table if exists post_filter;
create table if not exists post_filter
(
    post_id   uuid references post (id),
    filter_id int references "filter" (id),
    constraint post_filter_pkey
        primary key (post_id, filter_id)
);

drop table if exists "comment";
create table if not exists "comment"
(
    id      uuid      not null default uuid_generate_v4(),
    post_id uuid references post (id),
    user_id uuid references "user" (id),
    "date"  timestamp not null default date_trunc('minute', now()),
    "text"  text      not null,
    primary key (id)
);

drop table if exists "message";
create table if not exists "message"
(
    id      uuid not null default uuid_generate_v4(),
    post_id uuid not null references post (id),
    text    text not null,
    primary key (id)
);

drop table if exists user_post_liked;
create table if not exists user_post_liked
(
    user_id uuid references "user" (id),
    post_id uuid references post (id),
    constraint user_post_liked_pkey
        primary key (user_id, post_id)
);
