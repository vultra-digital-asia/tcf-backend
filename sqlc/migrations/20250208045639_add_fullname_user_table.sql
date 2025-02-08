-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN fullname TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN fullname;
-- +goose StatementEnd
