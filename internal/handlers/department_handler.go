package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"tcfback/internal/dto/department_dto"
	"tcfback/internal/middleware"
	"tcfback/internal/repositories"
	"tcfback/pkg/utils"
)

type DepartmentHandler struct {
	repo *repositories.DepartmentRepository
}

func NewDepartmentHandler(repo *repositories.DepartmentRepository) DepartmentHandler {
	return DepartmentHandler{
		repo: repo,
	}
}
func (d *DepartmentHandler) Router(g *echo.Group) {
	departmentGroup := g.Group("/departments")

	departmentGroup.GET("", d.GetAllDepartment, middleware.AuthMiddleware(middleware.RoleManager, middleware.RoleAdmin, middleware.RoleUser))
	departmentGroup.GET("/:id", d.GetOneDepartment, middleware.AuthMiddleware(middleware.RoleManager, middleware.RoleAdmin, middleware.RoleUser))
	departmentGroup.POST("", d.CreateDepartment, middleware.AuthMiddleware(middleware.RoleAdmin))
	departmentGroup.PATCH("", d.UpdateDepartment, middleware.AuthMiddleware(middleware.RoleAdmin))
}

func (d *DepartmentHandler) GetAllDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var req department.GetAllDepartmentParams

	// Read query parameters manually and convert to int32
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	req.Offset = int32(page)
	req.Limit = int32(limit)
	req.Name = c.QueryParam("name")
	req.IsDeleted = c.QueryParam("is_deleted") == "true"

	// Set default values if not provided
	if req.Offset <= 0 {
		req.Offset = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	result, err := d.repo.GetManyDepartment(ctx, req)

	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Error getting list of department_dto", err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get List department_dto", &result)
}

func (d *DepartmentHandler) GetOneDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user_dto ID", map[string]string{"error": "Invalid UUID format"})
	}
	user, err := d.repo.GetOneDepartment(ctx, id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "User not found", err)
	}

	return utils.SuccessResponse(c, http.StatusOK, "User retrieved successfully", &user)
}

func (d *DepartmentHandler) CreateDepartment(c echo.Context) error {
	ctx := c.Request().Context()

	// Get form values instead of JSON binding
	name := c.FormValue("name")

	if name == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", map[string]interface{}{"error": "Username and password are required"})
	}

	// Create request object manually
	req := department.CreateDepartmentRequest{
		Name: name,
	}

	validationErrors := req.Validate()
	if validationErrors != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", validationErrors)
	}

	// Check if name already exists
	//nameExist, _ := d.repo.GetOneByName(ctx, req.Name)
	//if nameExist.Name != "" {
	//	return utils.ErrorResponse(c, http.StatusBadRequest, "Name already exists", map[string]interface{}{"error": "Email already exists"})
	//}

	// Create Department
	result, err := d.repo.CreateDepartment(ctx, req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to Create Department", map[string]interface{}{"error": "Failed to Create Department"})
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success Create Department", result)
}

func (d *DepartmentHandler) UpdateDepartment(c echo.Context) error {
	ctx := c.Request().Context()

	// Get form values instead of JSON binding
	id := c.FormValue("id")
	name := c.FormValue("name")

	if id == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "User Id is required", map[string]interface{}{"error": "User Id is required"})
	}

	if name == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Validation Failed", map[string]interface{}{"error": "Username and password are required"})
	}

	// Create request object manually
	req := department.UpdateDepartmentRequest{
		ID:   uuid.MustParse(id),
		Name: &name,
	}

	// Check if name already exists
	//nameExist, _ := d.repo.GetOneByName(ctx, req.Name)
	//if nameExist.Name != "" {
	//	return utils.ErrorResponse(c, http.StatusBadRequest, "Name already exists", map[string]interface{}{"error": "Email already exists"})
	//}

	// Create Department
	result, err := d.repo.UpdateDepartment(ctx, req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to Update Department", map[string]interface{}{"error": "Failed to Update Department"})
	}

	return utils.SuccessResponse(c, http.StatusOK, "Success Update Department", result)
}
