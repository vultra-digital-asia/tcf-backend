package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
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
	user.GET("", h.GetAllUser, middleware.AuthMiddleware(middleware.RoleManager, middleware.RoleAdmin, middleware.RoleUser))
	user.GET("/:id", h.GetOneUser, middleware.AuthMiddleware(middleware.RoleUser, middleware.RoleAdmin))
	user.POST("", h.CreateUser, middleware.AuthMiddleware(middleware.RoleAdmin))
	user.PATCH("", h.UpdateUser, middleware.AuthMiddleware(middleware.RoleUser, middleware.RoleAdmin))

	auth := g.Group("/auth")

	//use Json
	//auth.POST("/login", h.LoginUser)

	//use formData
	auth.POST("/login", h.LoginUserFormData)
}
func (h *UserHandler) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.GetAllUserParams

	// Read query parameters manually and convert to int32
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	req.Offset = int32(page)
	req.Limit = int32(limit)
	req.FullName = c.QueryParam("full_name")
	req.UserName = c.QueryParam("user_name")
	req.Email = c.QueryParam("email")

	// Set default values if not provided
	if req.Offset <= 0 {
		req.Offset = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	//log.Info().Msgf("Fetching users with Page: %d, Limit: %d, FullName: %s, UserName: %s, Email: %s",
	//	req.Offset, req.Limit, req.FullName, req.UserName, req.Email)

	result, err := h.repo.GetAllUser(ctx, req)

	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Error getting list of users", err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get List User", &result)
}

func (h *UserHandler) GetOneUser(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID", map[string]string{"error": "Invalid UUID format"})
	}
	log.Info().Msgf("Received ID param: %s", id)
	user, err := h.repo.GetOneUser(ctx, id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "User not found", err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "User retrieved successfully", &user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	// Get form values instead of JSON binding
	email := c.FormValue("email")
	password := c.FormValue("password")
	fullName := c.FormValue("full_name")
	phone := c.FormValue("phone")
	username := c.FormValue("username")
	roleId := c.FormValue("role_id")
	positionId := c.FormValue("position_id")
	departmentId := c.FormValue("department_id")

	if email == "" || password == "" || fullName == "" || phone == "" || username == "" || roleId == "" || positionId == "" || departmentId == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", map[string]interface{}{"error": "Username and password are required"})
	}

	// Create request object manually
	req := dto.CreateUserRequest{
		Email:        email,
		Password:     password,
		Fullname:     fullName,
		Phone:        phone,
		Username:     username,
		RoleId:       uuid.MustParse(roleId),
		PositionId:   uuid.MustParse(positionId),
		DepartmentId: uuid.MustParse(departmentId),
	}

	//if err := c.Bind(&req); err != nil {
	//	return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON request", map[string]interface{}{"error": "Invalid JSON format"})
	//}

	validationErrors := req.Validate()
	if validationErrors != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", validationErrors)
	}

	// Check if email already exists
	userEmail, _ := h.repo.GetOneByEmail(ctx, req.Email)
	if userEmail.Email != "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Email already exists", map[string]interface{}{"error": "Email already exists"})
	}

	// wrong logic must get use GetOneByUserName
	userName, _ := h.repo.GetOneByUsername(ctx, req.Username)
	if userName.Username != "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Username already exists", map[string]interface{}{"error": "Username already exists"})
	}

	// Create user
	result, err := h.repo.CreateUser(ctx, req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to Create User", map[string]interface{}{"error": "Failed to Create User"})
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success Create User", result)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	// Get form values instead of JSON binding
	id := c.FormValue("id")
	email := c.FormValue("email")
	password := c.FormValue("password")
	fullName := c.FormValue("full_name")
	phone := c.FormValue("phone")
	username := c.FormValue("username")
	roleId := c.FormValue("role_id")
	positionId := c.FormValue("position_id")
	departmentId := c.FormValue("department_id")

	if id == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "User Id is required", map[string]interface{}{"error": "User Id is required"})
	}

	// Check if email already exists
	userEmail, _ := h.repo.GetOneByEmail(ctx, email)
	if userEmail.Email != "" && userEmail.ID != uuid.MustParse(id) {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Email already exists", map[string]interface{}{"error": "Email already exists"})
	}

	// wrong logic must get use GetOneByUserName
	userName, _ := h.repo.GetOneByUsername(ctx, username)
	if userName.Username != "" && userName.ID != uuid.MustParse(id) {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Username already exists", map[string]interface{}{"error": "Username already exists"})
	}

	//if email == "" || password == "" || fullName == "" || phone == "" || username == "" || roleId == "" || positionId == "" || departmentId == "" {
	//	return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", map[string]interface{}{"error": "Username and password are required"})
	//}

	req := dto.UpdateUserRequest{
		ID:           uuid.MustParse(id),
		Email:        utils.ToPtr(email),
		Password:     utils.ToPtr(password),
		Fullname:     utils.ToPtr(fullName),
		Phone:        utils.ToPtr(phone),
		Username:     utils.ToPtr(username),
		RoleId:       utils.ToUUIDPtr(roleId),
		PositionId:   utils.ToUUIDPtr(positionId),
		DepartmentId: utils.ToUUIDPtr(departmentId),
	}

	//log.Info().Msgf("Received ID param: %s", req.ID)
	//
	//return utils.ErrorResponse(c, http.StatusInternalServerError, "Debug", map[string]interface{}{"error": "Debug"})

	// update user
	result, err := h.repo.UpdateUser(ctx, req)

	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to Update User", map[string]interface{}{"error": "Failed to Update User"})
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success Update User", result)

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
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Wrong username or password", err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success Create User", result)
}

func (h *UserHandler) LoginUserFormData(c echo.Context) error {
	ctx := c.Request().Context()

	// Get form values instead of JSON binding
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate input
	if email == "" || password == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", map[string]string{"error": "Username and password are required"})
	}

	// Create request object manually
	req := dto.LoginRequest{
		Email:    email,
		Password: password,
	}

	// Validate using existing validation function
	validationErrors := req.Validate()
	if validationErrors != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", validationErrors)
	}

	// Call the repository to process login
	result, err := h.repo.Login(ctx, req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to login", err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Login Successful", result)
}
