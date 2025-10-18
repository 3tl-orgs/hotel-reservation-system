create table currencies (
    id serial primary key,
    code varchar(3) unique not null,
    name varchar(50) not null,
    icon varchar(255),
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

create table payment_events (
    id serial primary key,
    property_id integer not null,
    type varchar(20) default 'pay_in', -- pay_in, refund, pay_out
    is_payment_done bool default false,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

-- Bảng thanh toán chính (ghi nhận thanh toán của người dùng)
create table payments (
    id serial primary key,
    booking_id integer not null,
    pay_in_payment_event_id integer not null,
    user_id integer not null,
    property_id integer not null,
    payment_method varchar(20) default 'stripe',
    currency_id integer not null references currencies(id),
    amount decimal(10,2) not null,
    amount_in_base_currency decimal(10,2) not null,
    exchange_rate decimal(10,4) not null,
    provider_tnx_id varchar(255) not null,
    payment_status varchar(20) default 'pending',
    paid_at timestamp not null,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

-- Bảng hoàn tiền
create table payment_refunds (
    id serial primary key,
    user_id integer not null,
    booking_id integer not null,
    payment_id integer not null references payments(id),
    refund_reason text not null,
    property_id integer not null,
    pay_in_payment_event_id integer not null,
    refund_payment_event_id integer not null,
    payment_method varchar(20) default 'stripe',
    provider_tnx_id varchar(255) not null,
    refund_status varchar(20) default 'pending',
    currency_id integer not null references currencies(id),
    amount decimal(10,2) not null,
    amount_in_base_currency decimal(10,2) not null,
    exchange_rate decimal(10,4) not null,
    refund_at timestamp not null,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

-- Bảng bút toán kế toán (ledger)
create table ledgers (
    id serial primary key,
    transaction_id varchar(255) not null,
    account_id varchar(255) not null,
    debit decimal(10,2) not null,
    credit decimal(10,2) not null,
    description text,
    currency_id integer not null references currencies(id),
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);

-- Bảng ghi nhận việc giải ngân (trả tiền cho chủ khách sạn)
create table property_payouts (
    id serial primary key,
    booking_id integer not null,
    pay_in_payment_event_id integer not null,
    refund_payment_event_id integer,
    pay_out_payment_event_id integer not null,
    payout_date timestamp not null,
    property_id integer not null,
    status bool default true,
    created_at timestamp default now(),
    updated_at timestamp
);