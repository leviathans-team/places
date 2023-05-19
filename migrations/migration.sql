CREATE table users_info
(
    user_id       bigserial not null primary key,
    isAdmin bool default false,
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


