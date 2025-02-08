package utils

import (
	"math/big"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// HandleNullableString converts a pointer to a nullable pgtype.Text.
func HandleNullableString(input *string) pgtype.Text {
	if input == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *input, Valid: true}
}

// HandleString converts a string to a pgtype.Text.
func HandleString(input string) pgtype.Text {
	return pgtype.Text{String: input, Valid: true}
}

// HandleNullableBigInt converts a pointer to a nullable pgtype.Int8.
func HandleNullableBigInt(input *int64) pgtype.Int8 {
	if input == nil {
		return pgtype.Int8{Valid: false}
	}
	return pgtype.Int8{Int64: *input, Valid: true}
}

func HandleBigInt(input int64) pgtype.Int8 {
	return pgtype.Int8{Int64: input, Valid: true}
}

func HandleInt(input int32) pgtype.Int4 {
	return pgtype.Int4{Int32: input, Valid: true}
}

// HandleNullableInt converts a pointer to a nullable pgtype.Int4.
func HandleNullableInt(input *int32) pgtype.Int4 {
	if input == nil {
		return pgtype.Int4{Valid: false}
	}
	return pgtype.Int4{Int32: *input, Valid: true}
}

// HandleNullableNumeric converts a pointer to a nullable pgtype.Numeric.
func HandleNullableNumeric(input *int64) pgtype.Numeric {
	if input == nil {
		return pgtype.Numeric{Valid: false}
	}
	return pgtype.Numeric{Int: big.NewInt(*input), Valid: true}
}

// HandleNumericFloat converts a pointer to a float64 into pgtype.Numeric.
func HandleNumericFloat(input *float64) pgtype.Numeric {
	if input == nil {
		return pgtype.Numeric{Valid: false} // NULL
	}

	bigFloat := new(big.Float).SetFloat64(*input)
	bigInt, _ := bigFloat.Int(nil)
	return pgtype.Numeric{Int: bigInt, Valid: true}
}

// HandleNullableFloat converts a pointer to a nullable pgtype.Float8.
func HandleNullableFloat(input *float64) pgtype.Float8 {
	if input == nil {
		return pgtype.Float8{Valid: false}
	}
	return pgtype.Float8{Float64: *input, Valid: true}
}

// HandleNullableTimestamp converts a pointer to a nullable pgtype.Timestamptz.
func HandleNullableTimestamp(input *time.Time) pgtype.Timestamptz {
	if input == nil {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{Time: *input, Valid: true}
}

func HandleNullableDate(input *time.Time) pgtype.Date {
	if input == nil {
		return pgtype.Date{Valid: false}
	}
	return pgtype.Date{Time: *input, Valid: true}
}

func HandleDate(input time.Time) pgtype.Date {
	return pgtype.Date{Time: input, Valid: true}
}

func HandleBool(input bool) pgtype.Bool {
	return pgtype.Bool{Bool: input, Valid: true}
}

// HandleNullableBool converts a pointer to a nullable pgtype.Bool.
func HandleNullableBool(input *bool) pgtype.Bool {
	if input == nil {
		return pgtype.Bool{Valid: false}
	}
	return pgtype.Bool{Bool: *input, Valid: true}
}
