-- +goose Up
-- +goose StatementBegin
CREATE TABLE people (
    id serial primary key,
    name text,
    surname text,
    patronymic text
);

CREATE TABLE cars (
     regNum text primary key,
     mark text,
     model text,
     owner_id int references people(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cars;
DROP TABLE people;
-- +goose StatementEnd