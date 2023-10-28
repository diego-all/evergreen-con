CREATE DATABASE evergreen_con CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE evergreen_con;


DROP TABLE IF EXISTS device_types;

CREATE TABLE device_types (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(100) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

DROP TABLE IF EXISTS locations;

CREATE TABLE locations (
    id int AUTO_INCREMENT PRIMARY KEY,
    identifier varchar(50) NOT NULL,
    name varchar(100) NOT NULL,
    project varchar(100) NOT NULL,
    plot varchar(100) NOT NULL,
    property varchar(100) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);


DROP TABLE IF EXISTS devices;

CREATE TABLE devices (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(100) NOT NULL,
    location_id int,
    parameters varchar(255),
    type_id int,
    model varchar(50),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    FOREIGN KEY (type_id) REFERENCES device_types(id),
    FOREIGN KEY (location_id) REFERENCES locations(id)
);


DROP TABLE IF EXISTS applications;

CREATE TABLE applications (
    id int AUTO_INCREMENT PRIMARY KEY,
    identifier varchar(50) NOT NULL,
    name varchar(100) NOT NULL,
    port varchar(20),
    status varchar(50),
    type varchar(50),
    language varchar(50),
    device_id int,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    FOREIGN KEY (device_id) REFERENCES devices(id)
);


INSERT INTO device_types (name, created_at, updated_at) VALUES
('sensor', NOW(), NOW()),
('actuator', NOW(), NOW()),
('GPS', NOW(), NOW()),
('weather station', NOW(), NOW()),
('drone', NOW(), NOW()),
('monitor', NOW(), NOW());
