CREATE DATABASE beverages_booking;

USE beverages_booking;


CREATE TABLE admins (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);



CREATE TABLE beverages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type ENUM('hot', 'cold', 'hard') NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL
);

