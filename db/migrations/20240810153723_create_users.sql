-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
  id UUID primary key default gen_random_uuid(),
  first_name varchar(255) not null,
  last_name varchar(255) not null,
  email varchar(255) not null unique,
  password_hash varchar(255),
  created_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);
-- +goose StatementEnd

CREATE TABLE sessions (
  id UUID primary key default  gen_random_uuid(),
  user_id UUID not null references users(id),
  expires_at timestamp(0),
  created_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);

CREATE TABLE properties (
  id UUID PRIMARY KEY default gen_random_uuid(), 
  name varchar(255) not null,
  email varchar(255) not null,
  phone_number varchar(255) not null,
  address varchar(255) not null,
  about varchar(255) ,
  created_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);

CREATE TABLE propertyaccesses (
  id UUID PRIMARY KEY default gen_random_uuid(),
  property_id UUID references properties(id),
  user_id UUID references users(id),
  created_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc'),
  unique(user_id, property_id)
);


-- +goose Down
-- +goose StatementBegin
DROP table IF EXISTS users CASCADE;
DROP table IF EXISTS sessions CASCADE;
DROP table IF EXISTS properties CASCADE;
DROP table IF EXISTS propertyaccesses CASCADE;
-- +goose StatementEnd
