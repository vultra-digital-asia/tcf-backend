package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"tcfback/internal/db"
	"tcfback/internal/dto/user_dto"
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

func (r *UserRepository) GetAllUser(ctx context.Context, req user_dto.GetAllUserParams) ([]user_dto.GetAllUserResponse, error) {
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

	mappedUsers := make([]user_dto.GetAllUserResponse, len(users))
	for i, user := range users {
		var birthDateStr string
		if user.BirthDate.Valid {
			birthDateStr = user.BirthDate.Time.Format("2006-01-02 15:04:05") // Adjust format as needed
		}

		var birthPlaceStr string
		if user.BirthPlace.Valid {
			birthPlaceStr = user.BirthPlace.String
		}

		mappedUsers[i] = user_dto.GetAllUserResponse{
			ID:         user.ID.String(),
			Fullname:   user.FullName,
			Username:   user.Username,
			Email:      user.Email,
			Phone:      user.Phone,
			Address:    user.Address.String,
			BirthDate:  birthDateStr,
			BirthPlace: birthPlaceStr,
		}
	}

	return mappedUsers, nil
}

func (r *UserRepository) GetOneUser(ctx context.Context, id uuid.UUID) (db.User, error) {
	result, err := r.queries.GetUserById(ctx, id)
	return result, err
}

func (r *UserRepository) GetOneByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error) {
	result, err := r.queries.GetUserByEmail(ctx, email)
	return result, err
}

func (r *UserRepository) GetOneByUsername(ctx context.Context, username string) (db.User, error) {
	result, err := r.queries.GetUserByUserName(ctx, username)
	return result, err
}
func (r *UserRepository) CreateUser(ctx context.Context, req user_dto.CreateUserRequest) (*user_dto.CreateUserResponse, error) {

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
	response := user_dto.CreateUserResponse{
		ID:       result.ID.String(),
		Username: result.Username,
		Email:    result.Email,
		Role:     result.RoleName.String,
	}

	return &response, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, req user_dto.UpdateUserRequest) (*user_dto.UpdateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return nil, err
	}

	result, err := r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:           req.ID,
		Email:        *req.Email,
		Password:     string(hashedPassword),
		Username:     *req.Username,
		FullName:     *req.Fullname,
		Phone:        *req.Phone,
		RoleID:       *req.RoleId,
		DepartmentID: *req.DepartmentId,
		PositionID:   *req.PositionId,
	})

	// Log error with detailed message if the query fails
	if err != nil {
		log.Error().
			Err(err).
			Str("Query", "UpdateUser").
			Str("Email", *req.Email).
			Msg("Error executing UpdateUser query")
		return nil, err
	}

	response := user_dto.UpdateUserResponse{
		ID:       result.ID.String(),
		Username: result.Username,
		Email:    result.Email,
	}

	return &response, nil
}

func (r *UserRepository) Login(ctx context.Context, request user_dto.LoginRequest) (*user_dto.LoginResponse, map[string]custom_errors.FieldError) {

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

	return &user_dto.LoginResponse{
		ID:          getUser.ID.String(),
		Username:    getUser.Username,
		Email:       getUser.Email,
		AccessToken: token,
	}, nil

}
