-- +goose Up
-- +goose StatementBegin
ALTER TABLE notifications ADD FOREIGN KEY (user_id) REFERENCES users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE notifications DROP CONSTRAINT notifications_user_id_fkey;
-- +goose StatementEnd
