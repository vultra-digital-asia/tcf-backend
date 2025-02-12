package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tcfback/internal/dto"
	"tcfback/internal/middleware"
	"tcfback/internal/repositories"
	"tcfback/pkg/utils"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) UserHandler {
	return UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) Router(g *echo.Group) {

	user := g.Group("/users")

	//example for use on group
	//user.Use(middleware.AuthMiddleware([]string{"admin"}))

	//example on case by cae
	user.GET("", h.GetAllUser, middleware.AuthMiddleware(middleware.RoleManager, middleware.RoleAdmin))
	user.POST("", h.CreateUser)

	auth := g.Group("/auth")
	auth.POST("/login", h.LoginUser)
}

func (h *UserHandler) GetAllUser(c echo.Context) error {

	ctx := c.Request().Context()

	result, err := h.repo.GetAllUser(ctx)

	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Error get list user", err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get List User", &result)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON request", map[string]interface{}{"error": "Invalid JSON format"})
	}

	validationErrors := req.Validate()
	if validationErrors != nil {

		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", validationErrors)
	}

	result, err := h.repo.CreateUser(ctx, req)

	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to Create User", echo.Map{})
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success Create User", result)

}

func (h *UserHandler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON request", map[string]interface{}{"error": "Invalid JSON format"})
	}

	validationErrors := req.Validate()
	if validationErrors != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", validationErrors)
	}
	result, err := h.repo.Login(ctx, req)

	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to login", err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success Create User", result)
}
