// middleware/auth_middleware.go
package middleware

import (
	"net/http"
	"os"
	"strings"
	"tcfback/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleManager Role = "manager"
	RoleUser    Role = "user_dto"
)

// AuthMiddleware for Echo with Role-Based Access Control
func AuthMiddleware(allowedRoles ...Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization token required"})
			}

			// Extract token (format: "Bearer <token>")
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader { // No "Bearer " prefix
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
			}

			token, err := jwt.ParseWithClaims(tokenString, &services.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
			}

			// Store user_dto info in context
			claims, ok := token.Claims.(*services.JWTClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token claims"})
			}

			c.Set("username", claims.Username)
			c.Set("department_dto", claims.Department)
			c.Set("position", claims.Position)
			c.Set("role", claims.Role)

			// Role-based access control
			if len(allowedRoles) > 0 {
				userRole := claims.Role
				isRoleAllowed := false
				for _, allowedRole := range allowedRoles {
					if Role(userRole) == allowedRole {
						isRoleAllowed = true
						break
					}
				}

				if !isRoleAllowed {
					return c.JSON(http.StatusForbidden, map[string]string{"error": "Forbidden: You don't have permission to access this resource"})
				}
			}

			return next(c) // Continue to the next handler
		}
	}
}
