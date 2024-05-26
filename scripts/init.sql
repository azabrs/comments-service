CREATE TABLE IF NOT EXISTS users(
    id BIGINT PRIMARY KEY,
    login TEXT UNIQUE,
    password_hash TEXT NOT NULL
);