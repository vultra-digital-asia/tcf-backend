-- +goose Up
-- +goose StatementBegin
CREATE TABLE permissions (
    id UUID PRIMARY KEY,
    name VARCHAR,
    role_id UUID,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE permissions;
-- +goose StatementEnd
