-- +goose Up
-- +goose StatementBegin
create table users
(
    id         serial primary key,
    email      text,
    password   varchar(150),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
