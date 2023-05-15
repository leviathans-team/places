-- Миграции нужны для быстрого подъёма и настройки БД в любом окружениию. По сути просто SQL запросы со схемой БД.
create table event (
    id serial primary key,
    name varchar,
    chlen
);