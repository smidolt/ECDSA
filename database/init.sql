CREATE DATABASE status_db;

\c status_db

CREATE TABLE statuses (
    id SERIAL PRIMARY KEY,
    data BYTEA NOT NULL
);
