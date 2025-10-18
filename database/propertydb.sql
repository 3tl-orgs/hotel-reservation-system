create table property_types (
    id serial primary key,
    name varchar(100) not null unique,
    icon varchar(255),
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create index idx_property_types on property_types(status);

create table properties (
    id serial primary key,
    name varchar(255) not null,
    star_rating decimal(2,1),
    address varchar(255),
    country_id integer,
    province_id integer,
    ward_id integer,
    thumbnail varchar(255) not null,
    lat decimal(10,8),
    lng decimal(10,8),
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp,
    property_type_id integer references property_types(id)
);

create index idx_properties_country_id on properties(country_id);
create index idx_properties_province_id on properties(province_id);
create index idx_properties_ward_id on properties(ward_id);
create index idx_properties_status on properties(status);
create index idx_properties_lat_lng on properties using gist(point(lng, lat));

create table property_details (
    id serial primary key,
    property_id integer not null references properties(id) on delete cascade,
    description text,
    check_in_time time,
    check_out_time time,
    hotline varchar(15) not null,
    host_id integer,
    website_url varchar(255),
    email varchar(100),
    images jsonb,
    best_amenities jsonb,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create index idx_property_details_property_id on property_details(property_id);

create table facilities (
    id serial primary key,
    name varchar(100) not null unique,
    icon varchar(255),
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create table facility_properties (
    id serial primary key,
    property_id integer not null references properties(id) on delete cascade,
    facility_id integer not null references facilities(id) on delete cascade,
    amenity_ids jsonb,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create index idx_facility_properties_property_id on facility_properties(property_id);
create index idx_facility_properties_facility_id on facility_properties(facility_id);
create index idx_facility_properties_amenity_ids on facility_properties using gin (amenity_ids);

create table amenities (
    id serial primary key,
    name varchar(100) not null unique,
    icon varchar(255),
    type varchar(50) default 'general',
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create table room_types (
    id serial primary key,
    property_id integer not null references properties(id) on delete cascade,
    name varchar(100) not null,
    description text,
    images jsonb,
    total_rooms integer not null default 1,
    max_adults integer default 2,
    max_children integer default 0,
    base_price decimal(10,2),
    pay_before bool default false,
    bed_type varchar(100) default 'one_large_bed',
    pay_type varchar(30) default 'pay_in_place',
    best_amenities jsonb,
    amenities jsonb,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp,
    unique (property_id, name)
);

create index idx_room_types_property_id on room_types(property_id);
create index idx_room_types_status on room_types(status);

create table room_type_pricings (
    id serial primary key,
    room_type_id integer not null references room_types(id) on delete cascade,
    date date not null,
    price decimal(10,2),
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp,
    unique (room_type_id, date)
);

create index idx_room_type_pricings_room_type_id on room_type_pricings(room_type_id);
create index idx_room_type_pricings_date on room_type_pricings(date);
create unique index uq_room_type_pricings_room_type_date on room_type_pricings(room_type_id, date);

create table room_type_inventories (
    id serial primary key,
    property_id integer not null references properties(id) on delete cascade,
    room_type_id integer not null references room_types(id) on delete cascade,
    date date not null,
    reserved_rooms integer not null default 0,
    total_rooms integer not null default 0,
    available_rooms integer not null default 0,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp,
    unique (room_type_id, date)
);

create index idx_room_type_inventories_room_type_id on room_type_inventories(room_type_id);
create index idx_room_type_inventories_date on room_type_inventories(date);
create index idx_room_type_inventories_property_id on room_type_inventories(property_id);

create table refund_policy_rules (
    id serial primary key,
    property_id integer not null references properties(id) on delete cascade,
    room_type_id integer references room_types(id) on delete cascade,
    conditions jsonb,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create index idx_refund_policy_rules_property_id on refund_policy_rules(property_id);
create index idx_refund_policy_rules_room_type_id on refund_policy_rules(room_type_id);
create index idx_refund_policy_rules_status on refund_policy_rules(status);

create table hotel_wallets (
    id serial primary key,
    property_id integer not null references properties(id) on delete cascade,
    balance decimal(15,2) default 0,
    currency_id integer not null,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create unique index uq_hotel_wallets_property_id_currency_id on hotel_wallets(property_id, currency_id);
create index idx_hotel_wallets_status on hotel_wallets(status);