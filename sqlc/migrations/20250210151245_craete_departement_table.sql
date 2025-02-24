-- +goose Up
-- +goose StatementBegin
CREATE TABLE departments (
    id UUID PRIMARY KEY,
    name VARCHAR,
    deleted_at timestamp null

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE departments;
-- +goose StatementEnd
