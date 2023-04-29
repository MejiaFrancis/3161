-- Filename: migrations/000005_create_announcements_table.up.sql
CREATE TABLE IF NOT EXISTS announcements (
  announcements_id bigserial PRIMARY KEY,
  subject varchar(255) NOT NULL,
  content varchar(255) NOT NULL
);

