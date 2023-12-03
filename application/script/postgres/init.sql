create database dsm;

CREATE TABLE users (
                       user_id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       name VARCHAR(100),
                       birthday DATE,
                       intro TEXT
);