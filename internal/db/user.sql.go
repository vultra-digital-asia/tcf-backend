// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user_dto.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
WITH inserted_user AS (
    INSERT INTO users (id, email, password, username, full_name, phone, role_id, department_id, position_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, username, password, full_name, email, phone, birth_place, birth_date, address, position_id, department_id, role_id, deleted_at, created_at, updated_at
)
SELECT
    inserted_user.id, inserted_user.username, inserted_user.password, inserted_user.full_name, inserted_user.email, inserted_user.phone, inserted_user.birth_place, inserted_user.birth_date, inserted_user.address, inserted_user.position_id, inserted_user.department_id, inserted_user.role_id, inserted_user.deleted_at, inserted_user.created_at, inserted_user.updated_at,
    roles.name AS role_name -- Get the role name from roles table
FROM inserted_user
         LEFT JOIN roles ON inserted_user.role_id = roles.id
`

type CreateUserParams struct {
	ID           uuid.UUID
	Email        string
	Password     string
	Username     string
	FullName     string
	Phone        string
	RoleID       uuid.UUID
	DepartmentID uuid.UUID
	PositionID   uuid.UUID
}

type CreateUserRow struct {
	ID           uuid.UUID
	Username     string
	Password     string
	FullName     string
	Email        string
	Phone        string
	BirthPlace   pgtype.Text
	BirthDate    pgtype.Timestamp
	Address      pgtype.Text
	PositionID   uuid.UUID
	DepartmentID uuid.UUID
	RoleID       uuid.UUID
	DeletedAt    pgtype.Timestamp
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
	RoleName     pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.Username,
		arg.FullName,
		arg.Phone,
		arg.RoleID,
		arg.DepartmentID,
		arg.PositionID,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.BirthPlace,
		&i.BirthDate,
		&i.Address,
		&i.PositionID,
		&i.DepartmentID,
		&i.RoleID,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoleName,
	)
	return i, err
}

const getAllUser = `-- name: GetAllUser :many
SELECT id, username, password, full_name, email, phone, birth_place, birth_date, address, position_id, department_id, role_id, deleted_at, created_at, updated_at FROM users
WHERE
    (full_name ILIKE '%' || COALESCE($3, '') || '%')
  AND (username ILIKE '%' || COALESCE($4, '') || '%')
  AND (email ILIKE '%' || COALESCE($5, '') || '%')
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type GetAllUserParams struct {
	Limit    int32
	Offset   int32
	Fullname pgtype.Text
	Username pgtype.Text
	Email    pgtype.Text
}

func (q *Queries) GetAllUser(ctx context.Context, arg GetAllUserParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUser,
		arg.Limit,
		arg.Offset,
		arg.Fullname,
		arg.Username,
		arg.Email,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.FullName,
			&i.Email,
			&i.Phone,
			&i.BirthPlace,
			&i.BirthDate,
			&i.Address,
			&i.PositionID,
			&i.DepartmentID,
			&i.RoleID,
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

const getUserByEmail = `-- name: GetUserByEmail :one
select users.id, users.username, users.password, users.full_name, users.email, users.phone, users.birth_place, users.birth_date, users.address, users.position_id, users.department_id, users.role_id, users.deleted_at, users.created_at, users.updated_at, roles.name as role_name, positions.name as position_name, departments.name as department_name
from users
         left join roles on users.role_id = roles.id
         left join positions on users.position_id = positions.id
         left join departments on users.department_id = departments.id
where email = $1
`

type GetUserByEmailRow struct {
	ID             uuid.UUID
	Username       string
	Password       string
	FullName       string
	Email          string
	Phone          string
	BirthPlace     pgtype.Text
	BirthDate      pgtype.Timestamp
	Address        pgtype.Text
	PositionID     uuid.UUID
	DepartmentID   uuid.UUID
	RoleID         uuid.UUID
	DeletedAt      pgtype.Timestamp
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	RoleName       pgtype.Text
	PositionName   pgtype.Text
	DepartmentName pgtype.Text
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.BirthPlace,
		&i.BirthDate,
		&i.Address,
		&i.PositionID,
		&i.DepartmentID,
		&i.RoleID,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.RoleName,
		&i.PositionName,
		&i.DepartmentName,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
select id, username, password, full_name, email, phone, birth_place, birth_date, address, position_id, department_id, role_id, deleted_at, created_at, updated_at from users us where us.id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.BirthPlace,
		&i.BirthDate,
		&i.Address,
		&i.PositionID,
		&i.DepartmentID,
		&i.RoleID,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUserName = `-- name: GetUserByUserName :one
select id, username, password, full_name, email, phone, birth_place, birth_date, address, position_id, department_id, role_id, deleted_at, created_at, updated_at
from users
where username = $1
`

func (q *Queries) GetUserByUserName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUserName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.BirthPlace,
		&i.BirthDate,
		&i.Address,
		&i.PositionID,
		&i.DepartmentID,
		&i.RoleID,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
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
RETURNING id, username, password, full_name, email, phone, birth_place, birth_date, address, position_id, department_id, role_id, deleted_at, created_at, updated_at
`

type UpdateUserParams struct {
	ID           uuid.UUID
	Email        string
	Password     string
	Username     string
	FullName     string
	Phone        string
	RoleID       uuid.UUID
	DepartmentID uuid.UUID
	PositionID   uuid.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.Username,
		arg.FullName,
		arg.Phone,
		arg.RoleID,
		arg.DepartmentID,
		arg.PositionID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.BirthPlace,
		&i.BirthDate,
		&i.Address,
		&i.PositionID,
		&i.DepartmentID,
		&i.RoleID,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
