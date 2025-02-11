// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApprovalType string

const (
	ApprovalTypeIZIN     ApprovalType = "IZIN"
	ApprovalTypeLEMBUR   ApprovalType = "LEMBUR"
	ApprovalTypeCUTI     ApprovalType = "CUTI"
	ApprovalTypeREIMBURS ApprovalType = "REIMBURS"
)

func (e *ApprovalType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ApprovalType(s)
	case string:
		*e = ApprovalType(s)
	default:
		return fmt.Errorf("unsupported scan type for ApprovalType: %T", src)
	}
	return nil
}

type NullApprovalType struct {
	ApprovalType ApprovalType
	Valid        bool // Valid is true if ApprovalType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullApprovalType) Scan(value interface{}) error {
	if value == nil {
		ns.ApprovalType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ApprovalType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullApprovalType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ApprovalType), nil
}

type ApprovalFlow struct {
	ID           uuid.UUID
	OrderNumber  int32
	IsLastOrder  bool
	ApprovalID   uuid.UUID
	DepartmentID uuid.UUID
	FlowsNameID  uuid.UUID
	ApprovalType ApprovalType
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

type ApprovalFlowName struct {
	ID   uuid.UUID
	Name pgtype.Text
}

type CommonRequest struct {
	ID                uuid.UUID
	Status            string
	CurrentOrder      int32
	Reply             pgtype.Text
	Details           string
	ApprovalFlowsID   uuid.UUID
	UserRequestID     uuid.UUID
	CurrentApprovalID uuid.UUID
	SettleBy          uuid.UUID
	DepartmentID      uuid.UUID
	StartDate         pgtype.Timestamp
	EndDate           pgtype.Timestamp
	StartTime         pgtype.Text
	EndTime           pgtype.Text
	Url               pgtype.Text
	Amount            pgtype.Text
	RequestNumber     pgtype.Int4
}

type Department struct {
	ID   uuid.UUID
	Name pgtype.Text
}

type Notification struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	NotificationType pgtype.Text
	Title            pgtype.Text
	Message          pgtype.Text
	Status           pgtype.Text
	CreatedAt        pgtype.Timestamptz
	ReadAt           pgtype.Timestamp
}

type Permission struct {
	ID     uuid.UUID
	Name   pgtype.Text
	RoleID uuid.UUID
}

type Position struct {
	ID             uuid.UUID
	Name           pgtype.Text
	HierarchyLevel pgtype.Int4
}

type Role struct {
	ID   uuid.UUID
	Name pgtype.Text
}

type User struct {
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
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}
