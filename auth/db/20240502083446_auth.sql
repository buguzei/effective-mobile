-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id serial primary key,
    name text,
    email text,
    password text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
