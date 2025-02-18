-- +goose Up
-- +goose StatementBegin
CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id UUID,
    notification_type VARCHAR(20),
    title VARCHAR(50),
    message TEXT,
    status VARCHAR,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at timestamp with time zone default now(),
    read_at timestamp null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE notifications;
-- +goose StatementEnd
