-- MLM Application - Database Initialization Script
-- Run this script to create the database and user

-- Create database
CREATE DATABASE IF NOT EXISTS mlm_app 
CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;

-- Create user (change password in production!)
CREATE USER IF NOT EXISTS 'mlm_user'@'localhost' IDENTIFIED BY 'mlm_password_123';

-- Grant privileges
GRANT ALL PRIVILEGES ON mlm_app.* TO 'mlm_user'@'localhost';

-- Apply privileges
FLUSH PRIVILEGES;

-- Use the database
USE mlm_app;

-- Show success message
SELECT 'Database mlm_app created successfully!' AS Status;
SELECT 'User mlm_user created with all privileges' AS Status;
SELECT 'You can now run the backend server to create tables' AS Status;

-- Display connection info
SELECT 
    'Connection Info:' AS Info,
    'Host: localhost' AS Detail
UNION ALL
SELECT '', 'Database: mlm_app'
UNION ALL
SELECT '', 'User: mlm_user'
UNION ALL
SELECT '', 'Password: mlm_password_123'
UNION ALL
SELECT '', 'Port: 3306';
