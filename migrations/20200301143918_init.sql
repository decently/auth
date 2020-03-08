-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE account
(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(50) NOT NULL,
    hashed_password TEXT NOT NULL,
    admin BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE(email)
);

CREATE TABLE app
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE user_hashes
(
    user_id INT NOT NULL REFERENCES account(id),
    app_id INT NOT NULL REFERENCES app(id),
    password_hash TEXT NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table user_hashes;
drop table app;
drop table account;