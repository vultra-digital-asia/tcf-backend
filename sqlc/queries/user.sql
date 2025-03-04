-- name: GetAllUser :many
SELECT * FROM users
WHERE
    (full_name ILIKE '%' || COALESCE(@FullName, '') || '%')
  AND (username ILIKE '%' || COALESCE(@UserName, '') || '%')
  AND (email ILIKE '%' || COALESCE(@Email, '') || '%')
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountAllUsers :one
SELECT COUNT(*) AS total_count
FROM users
WHERE
    (full_name ILIKE '%' || COALESCE(@FullName, '') || '%')
  AND (username ILIKE '%' || COALESCE(@UserName, '') || '%')
  AND (email ILIKE '%' || COALESCE(@Email, '') || '%');

-- name: GetUserById :one
select * from users us where us.id = $1;

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

-- name: UpdateUser :one
UPDATE users
SET email = COALESCE($2, email),
    password = COALESCE($3, password),
    username = COALESCE($4, username),
    full_name = COALESCE($5, full_name),
    phone = COALESCE($6, phone),
    role_id = COALESCE($7, role_id),
    department_id = COALESCE($8, department_id),
    position_id = COALESCE($9, position_id)
WHERE id = $1
RETURNING *;


-- name: GetUserByEmail :one
select users.*, roles.name as role_name, positions.name as position_name, departments.name as department_name
from users
         left join roles on users.role_id = roles.id
         left join positions on users.position_id = positions.id
         left join departments on users.department_id = departments.id
where email = $1 AND users.deleted_at is null;

-- name: GetUserByUserName :one
select *
from users
where username = $1 AND deleted_at is null;