-- Filename: migrations/000013_create_equipment_usage_log_table.up.sql
CREATE TABLE IF NOT EXISTS equipment_usage_log (
  equipment_usage_log_id bigserial PRIMARY KEY,
  equipments_id bigint NOT NULL REFERENCES "equipment" ("equipments_id"),
  user_id bigserial NOT NULL REFERENCES "users" ("users_id"), 
  logs_id bigserial,
  time_borrowed TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  returned_status boolean NOT NULL
);
