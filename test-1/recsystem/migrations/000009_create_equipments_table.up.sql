-- Filename: migrations/000007_create_equipments_table.up.sql
CREATE TABLE IF NOT EXISTS equipments (
  equipments_id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  image bytea,
  equipment_type_id int NOT NULL REFERENCES equipment_types (equipment_types_id),
  status boolean NOT NULL,
  availability boolean NOT NULL 
);
