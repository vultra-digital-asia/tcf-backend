-- name: GetManyDepartment :many
SELECT * FROM departments
WHERE
    (name ILIKE '%' || COALESCE(@Name, '') || '%')
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountAllDepartments :one
SELECT COUNT(*) AS total_count
FROM departments
WHERE
    (name ILIKE '%' || COALESCE(@Name, '') || '%');

-- name: GetDepartmentById :one
select * from departments where id = $1;

-- name: CreateDepartment :one
INSERT INTO departments (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateDepartment :one
UPDATE departments
SET name = COALESCE($2, name)
WHERE id = $1
RETURNING *;

-- name: GetByName :one
select *
from departments
where name = $1 AND deleted_at is null;
