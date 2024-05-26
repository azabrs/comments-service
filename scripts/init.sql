CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    login TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS posts(
    id serial PRIMARY KEY,
    author TEXT,
    time_add TIMESTAMP,
    is_comment_enable BOOLEAN,
    subject TEXT
);