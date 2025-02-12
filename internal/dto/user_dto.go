package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"tcfback/pkg/custom_errors"
)

type CreateUserRequest struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Fullname     string    `json:"full_name"`
	Username     string    `json:"username"`
	Phone        string    `json:"phone"`
	RoleId       uuid.UUID `json:"role_id"`
	DepartmentId uuid.UUID `json:"department_id"`
	PositionId   uuid.UUID `json:"position_id"`
}

func (r CreateUserRequest) Validate() map[string]custom_errors.FieldError {
	err := validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Phone, validation.Required),
		validation.Field(&r.Fullname, validation.Required),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 20)),
	)

	return custom_errors.MapValidationErrors(err)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate() map[string]custom_errors.FieldError {
	err := validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 20)),
	)

	return custom_errors.MapValidationErrors(err)
}

type LoginResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
