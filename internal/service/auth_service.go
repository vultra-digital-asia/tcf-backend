package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"sync"
	"time"
)

var (
	jwtSecret     []byte
	jwtSecretOnce sync.Once
)

func getJWTSecret() []byte {
	jwtSecretOnce.Do(func() {
		jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	})
	return jwtSecret
}

type JWTClaims struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Position   string `json:"position"`
	Department string `json:"department"`
	jwt.RegisteredClaims
}

// Generate JWT Token with Role
func GenerateJWT(id, username, email, role, position, department string) (string, error) {
	expirationTime := time.Now().Add(100 * 365 * 24 * time.Hour) // 100 years validity

	claims := JWTClaims{
		Id:         id,
		Username:   username,
		Email:      email,
		Role:       role,
		Position:   position,
		Department: department,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return getJWTSecret(), nil
	})
}
