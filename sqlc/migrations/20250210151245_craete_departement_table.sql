-- +goose Up
-- +goose StatementBegin
CREATE TABLE departments (
    id UUID PRIMARY KEY,
    name VARCHAR
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE departments;
-- +goose StatementEnd
