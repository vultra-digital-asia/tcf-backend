-- +goose Up
-- +goose StatementBegin
CREATE TABLE common_requests (
    id UUID PRIMARY KEY,
    status VARCHAR NOT NULL,
    current_order INT NOT NULL,
    reply VARCHAR,
    details TEXT NOT NULL,
    approval_flows_id UUID NOT NULL,
    user_request_id UUID NOT NULL,
    current_approval_id UUID,
    settle_by UUID,
    department_id UUID,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    start_time VARCHAR,
    end_time VARCHAR,
    url VARCHAR,
    amount VARCHAR,
    request_number SERIAL UNIQUE,
    deleted_at timestamp null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE common_requests;
-- +goose StatementEnd
