-- +goose Up
-- +goose StatementBegin
CREATE TABLE departments (
    id UUID PRIMARY KEY,
    name VARCHARZ,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE departments;
-- +goose StatementEnd
