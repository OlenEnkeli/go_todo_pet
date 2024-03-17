CREATE TABLE users (
    id serial PRIMARY KEY,
    login text NOT NULL,
    password text NOT NULL,
    name text,
    created_at timestamp NOT NULL,
    updated_at timestamp
);

CREATE INDEX user_login_index ON users(login text_ops);
