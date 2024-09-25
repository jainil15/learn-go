-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id bigint primary key generated by default as identity,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null unique,
    password_hash varchar(255),
    created_at timestamp(0) not null default (now() at time zone 'utc'),
    updated_at timestamp(0) not null default (now() at time zone 'utc')
  );
-- +goose StatementEnd

CREATE TABLE sessions (
   id bigint primary key generated by default as identity,
   user_id bigint not null references users(id),
   expires_at timestamp(0),
   created_at timestamp(0) not null default (now() at time zone 'utc'),
   updated_at timestamp(0) not null default (now() at time zone 'utc')
);

CREATE TABLE properties (
  id bigint PRIMARY KEY generated by default as identity,
  name varchar(255) not null,
  email varchar(255) not null,
  phone_number varchar(255) not null,
  address varchar(255) not null,
  about varchar(255) ,
  created_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);

CREATE TABLE propertyaccesses (
  id bigint PRIMARY KEY generated by default as identity,
  property_id bigint references properties(id),
  user_id bigint references users(id),
  created_at timestamp(0) not null default (now() at time zone 'utc'),
  updated_at timestamp(0) not null default (now() at time zone 'utc')
);


-- +goose Down
-- +goose StatementBegin
DROP table IF EXISTS users CASCADE;
DROP table IF EXISTS sessions CASCADE;
DROP table IF EXISTS properties CASCADE;
DROP table IF EXISTS propertyaccesses CASCADE;
-- +goose StatementEnd
