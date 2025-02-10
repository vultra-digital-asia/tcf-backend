-- +goose Up
-- +goose StatementBegin
ALTER TABLE positions ADD COLUMN hierarchy_level INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE positions DROP COLUMN hirarcy_level;
-- +goose StatementEnd
