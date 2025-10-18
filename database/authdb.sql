create table users (
    id serial primary key,
    first_name varchar(30),
    last_name varchar(30),
    phone varchar(15) unique,
    birthday date,
    email varchar(100) unique,
    gender varchar(20) default 'other',
    avatar varchar(255),
    is_verified bool,
    status bool,
    created_at timestamp default now(),
    updated_at timestamp
);

create table accounts (
    id serial primary key,
    email varchar(100) not null unique,
    password_hash varchar(255) not null,
    salt varchar(255) not null unique,
    user_id integer not null unique,
    status bool,
    created_at timestamp default now(),
    updated_at timestamp
);

create table oauth_providers (
    id serial primary key,
    name varchar(50) not null unique,
    status bool,
    created_at timestamp,
    updated_at timestamp
);

create table user_oauths (
    user_id integer not null,
    provider_id integer not null,
    code varchar(255) not null unique,
    status bool,
    created_at timestamp default now(),
    updated_at timestamp
);

create table roles (
    id serial primary key,
    name varchar(30) not null unique,
    status bool,
    created_at timestamp default now(),
    updated_at timestamp
);

create table user_roles (
    user_id integer,
    role_id integer,
    status bool,
    created_at timestamp default now(),
    updated_at timestamp
);