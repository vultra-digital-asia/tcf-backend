package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"tcfback/internal/db"
	"tcfback/internal/dto"
	services "tcfback/internal/service"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return UserRepository{
		queries: queries,
	}
}

func (r *UserRepository) GetAllUser(ctx context.Context) ([]db.User, error) {

	result, err := r.queries.GetAllUser(ctx)

	return result, err
}

func (r *UserRepository) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.CreateUserResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return nil, err
	}

	result, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid.New(),
		Email:    req.Email,
		Password: string(hashedPassword),
		Username: req.Username,
		FullName: req.Fullname,
		Phone:    req.Phone,
	})

	if err != nil {
		log.Error().Err(err).Msg("Error Create User")
		return nil, err
	}
	response := dto.CreateUserResponse{
		ID:       result.ID.String(),
		Username: result.Username,
		Email:    result.Email,
	}

	return &response, nil
}

func (r *UserRepository) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {

	getUser, err := r.queries.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, errors.New("User not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(request.Password)); err != nil {
		log.Error().Err(err).Msg("invalid password")
		return nil, errors.New("invalid password")
	}

	token, err := services.GenerateJWT(getUser.ID.String(), getUser.Username, getUser.Email, getUser.RoleName.String, getUser.PositionName.String, getUser.DepartmentName.String)

	if err != nil {
		log.Error().Err(err).Msg("failed to generate token")
		return nil, errors.New("failed to generate token")
	}

	return &dto.LoginResponse{
		ID:          getUser.ID.String(),
		Username:    getUser.Username,
		Email:       getUser.Email,
		AccessToken: token,
	}, nil

}
