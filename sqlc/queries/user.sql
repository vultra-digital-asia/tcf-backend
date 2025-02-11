-- name: GetAllUser :many
select *
from users;

-- name: CreateUser :one
insert into users (id, email, password, username, full_name, phone)
VALUES ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetUserByEmail :one
select users.*, roles.name as role_name, positions.name as position_name, departments.name as department_name
from users
         left join roles on users.role_id = roles.id
         left join positions on users.position_id = positions.id
         left join departments on users.department_id = departments.id
where email = $1;