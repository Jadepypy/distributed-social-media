create database dsm;

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255),
                       nickname VARCHAR(100),
                       birthday DATE,
                       intro TEXT,
                       created_at BIGINT,
                       updated_at BIGINT
);
