-- +goose Up
-- +goose StatementBegin
CREATE TABLE cars (
     regNum text primary key,
     mark text,
     model text,
     owner_id int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cars;
-- +goose StatementEnd