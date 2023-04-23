-- Filename: migrations/000002_insert_users_data.up.sql
INSERT INTO users(first_name, last_name, email, age, address, phone_number, roles_id, password, status)
VALUES 
('Jane', 'Doe', 'sweetgirl@gmail.com', '20', '13 Mahogany Street', '666-9999', 'student', 'iloveUB2023', 'active'),
('Starla', 'Donde', 'sweetstar@gmail.com', '21', '10 Gill Street', '615-9876', 'admin', 'UBisthebest23', 'active'),