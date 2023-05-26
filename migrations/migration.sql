CREATE table users_info
(
    user_id       bigserial not null primary key,
    name       text      not null,
    surname    text      not null,
    patronymic text,
    email      text,
    phone text

);

create table users_login (
                             login_id bigint references users_info(user_id) primary key,
                             password_hash text
);



create table admins (
                        user_id bigint references users_info(user_id) primary key not null,
                        admin_level int
);


create table landlords (
                           user_id bigint references users_info(user_id) primary key,
                           post text not null,
                           places bigint[],
                           legal_entity text not null,
                           inn text not null,
                           industry int
); -- нужно от лехи взять ид места, созданного пользователем



