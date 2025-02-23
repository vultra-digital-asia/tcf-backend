-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    birth_place VARCHAR,
    birth_date TIMESTAMP,
    address TEXT,
    position_id UUID,
    department_id UUID,
    role_id UUID,
    deleted_at timestamp null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
