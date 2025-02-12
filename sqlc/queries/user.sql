-- name: GetAllUser :many
select *
from users;

-- name: CreateUser :one
WITH inserted_user AS (
    INSERT INTO users (id, email, password, username, full_name, phone, role_id, department_id, position_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING *
)
SELECT
    inserted_user.*,
    roles.name AS role_name -- Get the role name from roles table
FROM inserted_user
         LEFT JOIN roles ON inserted_user.role_id = roles.id;


-- name: GetUserByEmail :one
select users.*, roles.name as role_name, positions.name as position_name, departments.name as department_name
from users
         left join roles on users.role_id = roles.id
         left join positions on users.position_id = positions.id
         left join departments on users.department_id = departments.id
where email = $1;