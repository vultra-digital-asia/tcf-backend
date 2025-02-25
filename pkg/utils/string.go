package utils

import "github.com/google/uuid"

// Helper functions to convert values to pointers
func ToPtr[T any](v T) *T {
	return &v
}

func ToUUIDPtr(s string) *uuid.UUID {
	if s == "" {
		return nil
	}
	u := uuid.MustParse(s)
	return &u
}
