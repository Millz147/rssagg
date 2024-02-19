-- +goose Up
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY,
    name text NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;