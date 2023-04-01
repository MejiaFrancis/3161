-- Filename: migrations/000002_insert_users_data.up.sql
INSERT INTO users(first_name, last_name, email, age, address, phone_number, roles_id, password, status)
VALUES 
('Jane', 'Doe', 'sweetgirl@gmail.com', '20', '13 Mahogany Street', '666-9999', 'student', 'iloveUB2023', 'active'),
('Starla', 'Donde', 'sweetstar@gmail.com', '21', '10 Gill Street', '615-9876', 'admin', 'UBisthebest23', 'active'),
('John', 'Smith', 'johnsmith@gmail.com', '22', '5 Oak Lane', '555-1234', 'student', 'password123', 'active'),
('Emily', 'Johnson', 'emilyjohnson@gmail.com', '19', '20 Maple Street', '555-5678', 'student', 'mypassword', 'active'),
('Michael', 'Brown', 'michaelbrown@gmail.com', '20', '15 Elm Street', '555-4321', 'student', '12345678', 'active'),
('Jessica', 'Davis', 'jessicadavis@gmail.com', '18', '3 Pine Road', '555-8765', 'student', 'password1', 'active'),
('David', 'Miller', 'davidmiller@gmail.com', '21', '12 Birch Lane', '555-9876', 'student', 'abcdefg', 'active'),
('Jennifer', 'Wilson', 'jenniferwilson@gmail.com', '19', '8 Cedar Street', '555-2468', 'student', 'mypassword2', 'active'),
('William', 'Anderson', 'williamanderson@gmail.com', '20', '19 Maple Road', '555-1357', 'student', 'password1234', 'active'),
('Ava', 'Jackson', 'avajackson@gmail.com', '18', '7 Elm Lane', '555-8642', 'student', 'qwerty', 'active'),
('James', 'Thomas', 'jamesthomas@gmail.com', '21', '17 Oak Street', '555-3698', 'student', '12345', 'active'),
('Sophia', 'Robinson', 'sophiarobinson@gmail.com', '19', '22 Pine Lane', '555-7856', 'student', 'mypassword3', 'active'),
('Daniel', 'White', 'danielwhite@gmail.com', '20', '14 Cedar Road', '555-4682', 'student', 'password12345', 'active'),
('Olivia', 'Harris', 'oliviaharris@gmail.com', '18', '6 Maple Lane', '555-9753', 'student', 'abcdef', 'active'),
('Benjamin', 'Young', 'benjaminyoung@gmail.com', '21', '11 Elm Street', '555-3579', 'student', 'mypassword4', 'active'),
('Elizabeth', 'Allen', 'elizabethallen@gmail.com', '19', '16 Oak Road', '555-6812', 'student', 'password123456', 'active'),
('Christopher', 'King', 'christopherking@gmail.com', '20', '9 Pine Lane', '555-2589', 'student', 'qwertyuiop', 'active');