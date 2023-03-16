-- Filename: migrations/000001_create_users_table.up.sql
CREATE TABLE IF NOT EXISTS users (
  user_id bigserial PRIMARY KEY,
  email citext UNIQUE NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  age int NOT NULL,
  address VARCHAR(255) users_address_enum,
  phone_number bigserial NOT NULL,
  roles_id int NOT NULL REFERENCES "roles" ("roles_id"),
  password citext UNIQUE NOT NULL,
  status boolean,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
