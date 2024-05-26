CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    login TEXT UNIQUE
);