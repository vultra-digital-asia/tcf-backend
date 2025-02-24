-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name VARCHAR,
    deleted_at timestamp null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roles;
-- +goose StatementEnd
