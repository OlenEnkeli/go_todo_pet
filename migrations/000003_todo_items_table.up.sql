CREATE TABLE todo_items(
    id serial PRIMARY KEY,
    list_id int REFERENCES todo_lists(id),
    title text not null,
    description text,
    item_order int,
    is_done bool DEFAULT false,
    done_until timestamp,
    created_at timestamp NOT NULL,
    updated_at timestamp
);

CREATE INDEX item_order_index ON todo_items(item_order int4_ops);
CREATE INDEX item_list_id_index ON todo_items(list_id int4_ops);
