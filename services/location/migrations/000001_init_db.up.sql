create table countries (
                           id serial primary key,
                           name varchar(30) not null,
                           code varchar(10) not null unique,
                           created_at timestamp default now(),
                           updated_at timestamp,
                           status bool
);

create table provinces (
                           id serial primary key,
                           name varchar(30) not null,
                           code varchar(10) not null unique,
                           country_id integer not null,
                           created_at timestamp default now(),
                           updated_at timestamp,
                           status bool
);

create table wards (
                       id serial primary key,
                       name varchar(30) not null,
                       code varchar(10) not null unique,
                       province_id integer not null,
                       created_at timestamp default now(),
                       updated_at timestamp,
                       status bool
);







