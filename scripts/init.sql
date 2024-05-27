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


CREATE TABLE IF NOT EXISTS comments(
    id serial PRIMARY KEY,
    comment_data TEXT,
    parent_id INT,
    post_id INT NOT NULL,
    nesting_level INT,
    time_add TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON UPDATE CASCADE
);