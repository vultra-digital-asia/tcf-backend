package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"tcfback/internal/db"
	"tcfback/internal/dto/department_dto"
)

type DepartmentRepository struct {
	queries *db.Queries
}

func NewDepartmentRepository(queries *db.Queries) DepartmentRepository {
	return DepartmentRepository{
		queries: queries,
	}
}

func (d *DepartmentRepository) GetManyDepartment(ctx context.Context, req department.GetAllDepartmentParams) ([]department.GetDepartmentResponse, error) {
	offset := (req.Offset - 1) * req.Limit

	departments, err := d.queries.GetManyDepartment(ctx, db.GetManyDepartmentParams{
		Name:   pgtype.Text{String: req.Name, Valid: req.Name != ""},
		Limit:  int32(req.Limit),
		Offset: int32(offset),
	})

	if err != nil {
		log.Info().Err(err).Msg("Error fetching users")
		return nil, err
	}

	mappedDepartments := make([]department.GetDepartmentResponse, len(departments))
	for i, dept := range departments {

		var deletedAt string
		if dept.DeletedAt.Valid {
			deletedAt = dept.DeletedAt.Time.Format("2006-01-02 15:04:05") // Adjust format as needed
		}

		mappedDepartments[i] = department.GetDepartmentResponse{
			ID:        dept.ID.String(),
			Name:      dept.Name.String,
			DeletedAt: deletedAt,
		}
	}

	return mappedDepartments, nil
}

func (d *DepartmentRepository) GetOneDepartment(ctx context.Context, id uuid.UUID) (department.GetDepartmentResponse, error) {
	result, err := d.queries.GetDepartmentById(ctx, id)
	var deletedAt string
	if result.DeletedAt.Valid {
		deletedAt = result.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	dept := department.GetDepartmentResponse{
		ID:        result.ID.String(),
		Name:      result.Name.String,
		DeletedAt: deletedAt,
	}
	return dept, err
}

func (d *DepartmentRepository) CreateDepartment(ctx context.Context, req department.CreateDepartmentRequest) (*department.CreateDepartmentResponse, error) {

	result, err := d.queries.CreateDepartment(ctx, db.CreateDepartmentParams{
		ID:   uuid.New(),
		Name: pgtype.Text{String: req.Name, Valid: true},
	})

	if err != nil {
		log.Error().Err(err).Msg("Error Create Department")
		return nil, err
	}
	response := department.CreateDepartmentResponse{
		ID:   result.ID.String(),
		Name: result.Name.String,
	}

	return &response, nil
}

func (d *DepartmentRepository) UpdateDepartment(ctx context.Context, req department.UpdateDepartmentRequest) (*department.UpdateDepartmentResponse, error) {

	result, err := d.queries.UpdateDepartment(ctx, db.UpdateDepartmentParams{
		ID:   req.ID,
		Name: pgtype.Text{String: *req.Name, Valid: true},
	})

	if err != nil {
		log.Error().Err(err).Msg("Error Create Department")
		return nil, err
	}

	response := department.UpdateDepartmentResponse{
		ID:   result.ID.String(),
		Name: result.Name.String,
	}

	return &response, nil
}
