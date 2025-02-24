-- +goose Up
-- +goose StatementBegin
CREATE TABLE permissions (
    id UUID PRIMARY KEY,
    name VARCHAR,
    role_id UUID,
    deleted_at timestamp null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permissions;
-- +goose StatementEnd
