CREATE TABLE auths (
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN DEFAULT TRUE
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
   auth_id INT REFERENCES auths(id) ON DELETE CASCADE,
   first_name VARCHAR(30),
   last_name VARCHAR(30),
   phone VARCHAR(15),
   birthday DATE,
   avatar JSONB,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   status BOOLEAN DEFAULT TRUE
);

CREATE TABLE roles (
   id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL UNIQUE,
   description VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   status BOOLEAN DEFAULT TRUE
);

CREATE TABLE user_roles (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN DEFAULT TRUE
);

CREATE TABLE countries (
   id SERIAL PRIMARY KEY,
   name VARCHAR(50) UNIQUE NOT NULL,
   code VARCHAR(30) UNIQUE NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   status BOOLEAN DEFAULT TRUE
);

CREATE TABLE provinces (
   id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   code VARCHAR(30) NOT NULL,
   country_id INT REFERENCES countries(id) ON DELETE CASCADE,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   status BOOLEAN DEFAULT TRUE
);

create table districts (
   id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   code VARCHAR(30) NOT NULL,
   country_id INT REFERENCES countries(id) ON DELETE CASCADE,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   status BOOLEAN DEFAULT TRUE
)

CREATE TABLE wards (
   id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   code VARCHAR(30) NOT NULL,
   province_id INT REFERENCES provinces(id) ON DELETE CASCADE,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   status BOOLEAN DEFAULT TRUE
);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    content TEXT,
    user_id INT REFERENCES users(id),
    property_id INT REFERENCES properties(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN DEFAULT TRUE
);

CREATE TABLE user_ratings (
    user_id INT REFERENCES users(id),
    hotel_id INT REFERENCES properties(id),
    score DECIMAL(2,1) CHECK (score BETWEEN 1 AND 5),
    PRIMARY KEY (user_id, property_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN DEFAULT TRUE
);

create table property_types (
    id serial primary key,
    name varchar(50) not null unique,
    desc text,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN DEFAULT TRUE
)

CREATE TABLE properties (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    address VARCHAR(255),
    property_type_id int,
    country_id INT REFERENCES countries(id),
    province_id INT REFERENCES provinces(id),
    ward_id INT REFERENCES wards(id),
    capacity INT DEFAULT 1,
    rating DECIMAL(2,1),
    images JSONB,
    policies jsonb,
    latitude DECIMAL(10,6),
    longitude DECIMAL(10,6),
    host_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN DEFAULT TRUE
);

CREATE TABLE room_types (
    id SERIAL PRIMARY KEY,
    property_id INT REFERENCES properties(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    capacity INT DEFAULT 1,
    base_price DECIMAL(10,2) NOT NULL,
    total_rooms INT DEFAULT 1,
    images JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status BOOLEAN DEFAULT TRUE
);

CREATE TABLE rooms (
   id SERIAL PRIMARY KEY,
   room_type_id INT REFERENCES room_types(id) ON DELETE CASCADE,
   name VARCHAR(50) NOT NULL,
   room_number VARCHAR(50),
   status bool true,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bookings (
      id SERIAL PRIMARY KEY,
      user_id INT NOT NULL REFERENCES users(id),
      hotel_id INT REFERENCES properties(id),
      room_type_id INT REFERENCES room_types(id),
      check_in DATE NOT NULL,
      check_out DATE NOT NULL,
      total_adults INT DEFAULT 1,
      total_children INT DEFAULT 0,
      num_rooms INT DEFAULT 1,
      total_amount DECIMAL(10,2),
      discount DECIMAL(10,2) DEFAULT 0,
      final_amount decimal(10,2),
      note TEXT,
      booking_status VARCHAR(20) DEFAULT 'pending',
      payment_status VARCHAR(20) DEFAULT 'unpaid',  -- unpaid | paid | refunded
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      status BOOLEAN DEFAULT TRUE
);

CREATE TABLE booking_details (
      id SERIAL PRIMARY KEY,
      booking_id INT REFERENCES bookings(id) ON DELETE CASCADE,
      room_id INT REFERENCES rooms(id),
      price DECIMAL(10,2),
      note TEXT,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      status BOOLEAN DEFAULT TRUE
);

CREATE TABLE payment_methods (
      id SERIAL PRIMARY KEY,
      name VARCHAR(50) NOT NULL UNIQUE,
      method_type VARCHAR(30) DEFAULT 'card', -- card | wallet | bank_transfer
      provider VARCHAR(50),
      supported_currencies JSONB,
      config JSONB,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      status BOOLEAN DEFAULT TRUE
);

CREATE TABLE transactions (
      id SERIAL PRIMARY KEY,
      booking_id INT REFERENCES bookings(id),
      payment_method_id INT REFERENCES payment_methods(id),
      amount DECIMAL(10,2),
      currency VARCHAR(20) DEFAULT 'VND',
      metadata JSONB,
      transaction_status VARCHAR(20) DEFAULT 'pending',
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      status BOOLEAN DEFAULT TRUE
);

CREATE TABLE refunds (
      id SERIAL PRIMARY KEY,
      transaction_id INT REFERENCES transactions(id),
      amount DECIMAL(10,2),
      currency VARCHAR(20),
      reason TEXT,
      metadata JSONB,
      refund_status VARCHAR(20) DEFAULT 'pending',
      processed_at TIMESTAMP,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      status BOOLEAN DEFAULT TRUE
);
