create table notification_channels (
    id serial primary key,
    code varchar(30) unique not null,
    name varchar(50) not null,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create table notification_templates (
    id serial primary key,
    code varchar(30) unique not null,
    channel_id integer not null references notification_channels(id),
    subject varchar(100) not null,
    content text not null,
    is_html bool default false,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create table notifications (
    id serial primary key,
    template_id integer not null references notification_templates(id),
    user_id integer not null,
    recipient varchar(100) not null,
    subject varchar(100) not null,
    content text not null,
    notification_status varchar(20) default 'pending',
    retry_count integer default 0,
    sent_at timestamp,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);