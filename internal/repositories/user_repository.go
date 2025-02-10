package repositories

import (
	"context"
	"github.com/rs/zerolog/log"
	"tcfback/internal/db"
	"tcfback/internal/dto"
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

func (r *UserRepository) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*db.User, error) {

	result, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		//Email:    utils.HandleString(req.Email),
		//Password: utils.HandleString(req.Password),
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Error().Err(err).Msg("Error Create User")
		return nil, err
	}

	return &result, nil
}
