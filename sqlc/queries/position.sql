-- name: GetManyPosition :many
SELECT * FROM positions
WHERE
    (name ILIKE '%' || COALESCE(@Name, '') || '%')
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetPositionById :one
select * from positions where id = $1;

-- name: CreatePosition :one
INSERT INTO positions (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: UpdatePosition :one
UPDATE positions
SET name = COALESCE($2, name)
WHERE id = $1
RETURNING *;
