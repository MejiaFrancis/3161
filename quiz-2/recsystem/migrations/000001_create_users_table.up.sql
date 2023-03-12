-- Filename: migrations/000007_create_users_table.up.sql
CREATE TABLE IF NOT EXISTS users (
  user_id serial PRIMARY KEY,
  user_name text NOT NULL,
  email citext UNIQUE NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);