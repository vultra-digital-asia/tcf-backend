// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: department_dto.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createDepartment = `-- name: CreateDepartment :one
INSERT INTO departments (id, name)
VALUES ($1, $2)
RETURNING id, name, deleted_at, created_at, updated_at
`

type CreateDepartmentParams struct {
	ID   uuid.UUID
	Name pgtype.Text
}

func (q *Queries) CreateDepartment(ctx context.Context, arg CreateDepartmentParams) (Department, error) {
	row := q.db.QueryRow(ctx, createDepartment, arg.ID, arg.Name)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDepartmentById = `-- name: GetDepartmentById :one
select id, name, deleted_at, created_at, updated_at from departments where id = $1
`

func (q *Queries) GetDepartmentById(ctx context.Context, id uuid.UUID) (Department, error) {
	row := q.db.QueryRow(ctx, getDepartmentById, id)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getManyDepartment = `-- name: GetManyDepartment :many
SELECT id, name, deleted_at, created_at, updated_at FROM departments
WHERE
    (name ILIKE '%' || COALESCE($3, '') || '%')
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type GetManyDepartmentParams struct {
	Limit  int32
	Offset int32
	Name   pgtype.Text
}

func (q *Queries) GetManyDepartment(ctx context.Context, arg GetManyDepartmentParams) ([]Department, error) {
	rows, err := q.db.Query(ctx, getManyDepartment, arg.Limit, arg.Offset, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Department
	for rows.Next() {
		var i Department
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDepartment = `-- name: UpdateDepartment :one
UPDATE departments
SET name = COALESCE($2, name)
WHERE id = $1
RETURNING id, name, deleted_at, created_at, updated_at
`

type UpdateDepartmentParams struct {
	ID   uuid.UUID
	Name pgtype.Text
}

func (q *Queries) UpdateDepartment(ctx context.Context, arg UpdateDepartmentParams) (Department, error) {
	row := q.db.QueryRow(ctx, updateDepartment, arg.ID, arg.Name)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
