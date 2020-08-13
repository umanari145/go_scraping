CREATE TABLE comments (
    id  SERIAL PRIMARY KEY,
    thread_id int DEFAULT NULL,
    thread_no smallint DEFAULT NULL,
    contents text DEFAULT NULL,
    comment_date timestamp,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE threads (
    id  SERIAL PRIMARY KEY,
    thread_key varchar(100) DEFAULT NULL,
    title varchar(100) DEFAULT NULL,
    is_close boolean DEFAULT false,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);