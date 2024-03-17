CREATE TABLE todo_lists(
    id serial PRIMARY KEY,
    user_id int REFERENCES users(id),
    title text not null,
    description text,
    list_order int,
    created_at timestamp NOT NULL,
    updated_at timestamp
);

CREATE INDEX list_order_index ON todo_lists(list_order int4_ops);