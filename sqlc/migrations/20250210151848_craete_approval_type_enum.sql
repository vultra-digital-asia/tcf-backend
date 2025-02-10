-- +goose Up
-- +goose StatementBegin
CREATE TYPE  approval_type AS ENUM (
    'IZIN',
    'LEMBUR',
    'CUTI',
    'REIMBURS'
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE approval_type;
-- +goose StatementEnd
