package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"tcfback/internal/db"
	"tcfback/internal/dto"
	services "tcfback/internal/service"
	"tcfback/pkg/custom_errors"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return UserRepository{
		queries: queries,
	}
}

func (r *UserRepository) GetAllUser(ctx context.Context, req dto.GetAllUserParams) ([]db.User, error) {
	offset := (req.Offset - 1) * req.Limit

	users, err := r.queries.GetAllUser(ctx, db.GetAllUserParams{
		Fullname: pgtype.Text{String: req.FullName, Valid: req.FullName != ""},
		Username: pgtype.Text{String: req.UserName, Valid: req.UserName != ""},
		Email:    pgtype.Text{String: req.Email, Valid: req.Email != ""},
		Limit:    int32(req.Limit),
		Offset:   int32(offset),
	})

	if err != nil {
		log.Error().Err(err).Msg("Error fetching users")
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetOneUser(ctx context.Context, id uuid.UUID) (db.User, error) {
	result, err := r.queries.GetUserById(ctx, id)
	return result, err
}

func (r *UserRepository) GetOneByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error) {
	result, err := r.queries.GetUserByEmail(ctx, email)
	return result, err
}
func (r *UserRepository) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.CreateUserResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return nil, err
	}

	result, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		ID:           uuid.New(),
		Email:        req.Email,
		Password:     string(hashedPassword),
		Username:     req.Username,
		FullName:     req.Fullname,
		Phone:        req.Phone,
		RoleID:       req.RoleId,
		DepartmentID: req.DepartmentId,
		PositionID:   req.PositionId,
	})

	if err != nil {
		log.Error().Err(err).Msg("Error Create User")
		return nil, err
	}
	response := dto.CreateUserResponse{
		ID:       result.ID.String(),
		Username: result.Username,
		Email:    result.Email,
		Role:     result.RoleName.String,
	}

	return &response, nil
}

func (r *UserRepository) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, map[string]custom_errors.FieldError) {

	getUser, err := r.queries.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, custom_errors.MapValidationErrors(err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(request.Password)); err != nil {
		log.Error().Err(err).Msg("invalid password")
		//return nil, errors.New("invalid password")
		return nil, custom_errors.MapValidationErrors(err)
	}

	token, err := services.GenerateJWT(getUser.ID.String(), getUser.Username, getUser.Email, getUser.RoleName.String, getUser.PositionName.String, getUser.DepartmentName.String)

	if err != nil {
		log.Error().Err(err).Msg("failed to generate token")
		//return nil, errors.New("failed to generate token")
		return nil, custom_errors.MapValidationErrors(err)
	}

	return &dto.LoginResponse{
		ID:          getUser.ID.String(),
		Username:    getUser.Username,
		Email:       getUser.Email,
		AccessToken: token,
	}, nil

}
