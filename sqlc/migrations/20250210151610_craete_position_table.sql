-- +goose Up
-- +goose StatementBegin
CREATE TABLE positions (
    id UUID PRIMARY KEY,
    name VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE positions;
-- +goose StatementEnd
