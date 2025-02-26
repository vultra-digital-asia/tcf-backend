package department

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"tcfback/pkg/custom_errors"
)

type GetManyResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	DeletedAt string    `json:"deleted_at"`
}

type GetAllDepartmentParams struct {
	Name      string
	IsDeleted bool
	Limit     int32
	Offset    int32
}

type GetDepartmentResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	DeletedAt string `json:"deleted_at"`
}

type GetOneQuery struct {
	ID string `json:"id"`
}

type CreateDepartmentRequest struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name string    `json:"name"`
}

type CreateDepartmentResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateDepartmentRequest struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      *string   `json:"name"`
	DeletedAt *string   `json:"deleted_at,omitempty"`
}

type UpdateDepartmentResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

func (d CreateDepartmentRequest) Validate() map[string]custom_errors.FieldError {
	err := validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required),
	)

	return custom_errors.MapValidationErrors(err)
}
