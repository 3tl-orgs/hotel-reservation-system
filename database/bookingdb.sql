create table bookings (
    id serial primary key,
    booking_number varchar(50) not null unique,
    user_id integer,
    property_id integer not null,
    check_in date not null,
    check_out date not null,
    total_children integer not null,
    total_adult integer not null,
    guest_first_name varchar(30),
    guest_last_name varchar(30),
    guest_email varchar(100),
    guest_phone varchar(15),
    tax decimal(10,2),
    booking_status varchar(20) default 'pending',
    total_amount decimal(10,2) not null,
    final_amount decimal(10,2) not null,
    discount_amount decimal(10,2) not null,
    currency_id integer not null,
    note text,
    status bool,
    created_at timestamp default now(),
    updated_at timestamp
);

create index idx_bookings_check_in on bookings(check_in);
create index idx_bookings_check_out on bookings(check_out);
create index idx_bookings_booking_status on bookings(booking_status);

create table booking_details (
    id serial primary key,
    booking_id integer not null,
    room_type_id integer not null,
    quantity integer not null,
    price decimal(10,2) not null,
    status bool,
    created_at timestamp default now(),
    updated_at timestamp
);

create table booking_cancellations (
id serial primary key,
    booking_id integer not null,
    cancelled_by integer,
    reason text not null,
    refund_amount decimal(10,2),
    cancelled_at timestamp,
    status bool,
    created_at timestamp default now(),
    updated_at timestamp
);






