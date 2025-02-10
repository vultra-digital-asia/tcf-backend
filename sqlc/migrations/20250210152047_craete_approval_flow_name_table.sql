-- +goose Up
-- +goose StatementBegin
CREATE TABLE approval_flow_names (
    id UUID PRIMARY KEY,
    name VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE approval_flow_names;
-- +goose StatementEnd
