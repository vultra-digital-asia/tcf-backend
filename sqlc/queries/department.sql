-- name: GetManyDepartment :many
SELECT * FROM departments
WHERE
    (name ILIKE '%' || COALESCE(@Name, '') || '%')
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetDepartmentById :one
select * from departments where id = $1;

-- name: CreateDepartment :one
INSERT INTO departments (id, name)
VALUES ($1, $2);

-- name: UpdateDepartment :one
UPDATE departments
SET name = COALESCE($2, name)
WHERE id = $1
RETURNING *;
