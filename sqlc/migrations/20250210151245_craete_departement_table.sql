-- +goose Up
-- +goose StatementBegin
CREATE TABLE departments (
    id UUID PRIMARY KEY,
    name VARCHAR,
    deleted_at timestamp null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE departments;
-- +goose StatementEnd
