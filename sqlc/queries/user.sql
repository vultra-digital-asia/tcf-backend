-- name: GetAllUser :many
select * from users;

-- name: CreateUser :one
insert into users (email, password) VALUES ($1, $2) returning *;