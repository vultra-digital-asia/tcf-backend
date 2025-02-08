package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 20)),
	)
}
