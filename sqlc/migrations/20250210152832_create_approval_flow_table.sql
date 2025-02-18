-- +goose Up
-- +goose StatementBegin
CREATE TABLE approval_flows (
    id UUID PRIMARY KEY,
    order_number INT NOT NULL,
    is_last_order BOOLEAN NOT NULL,
    approval_id UUID,
    department_id UUID,
    flows_name_id UUID,
    approval_type approval_type NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE approval_flows;
-- +goose StatementEnd
