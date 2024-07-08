package repository

import (
	"context"
	"errors"
	"github.com/dennis0126/network-monitor/internal/db"
	"github.com/dennis0126/network-monitor/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	ctx     context.Context
	queries *db.Queries
}

func NewUserRepository(ctx context.Context, queries *db.Queries) UserRepository {
	return UserRepository{ctx: ctx, queries: queries}
}

func (r UserRepository) CreateUser(user model.User) (model.User, error) {
	dbUser, err := r.queries.CreateUser(r.ctx, db.CreateUserParams{
		ID: user.ID, Name: user.Name, PasswordHash: user.PasswordHash,
		CreatedAt: pgtype.Timestamptz{Time: user.CreatedAt, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: user.UpdatedAt, Valid: true},
	})
	if err != nil {
		return model.User{}, err
	}
	return unmarshallDbUser(dbUser), nil
}

func (r UserRepository) ListUsers() ([]model.User, error) {
	dbUsers, err := r.queries.ListUsers(r.ctx)
	if err != nil {
		return []model.User{}, err
	}

	users := []model.User{}
	for _, dbUser := range dbUsers {
		users = append(users, unmarshallDbUser(dbUser))
	}
	return users, nil
}

func (r UserRepository) GetUserById(id string) (*model.User, error) {
	dbUser, err := r.queries.GetUserById(r.ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user := unmarshallDbUser(dbUser)
	return &user, nil
}

func (r UserRepository) GetUserByName(name string) (*model.User, error) {

	dbUser, err := r.queries.GetUserByName(r.ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user := unmarshallDbUser(dbUser)
	return &user, nil
}

func unmarshallDbUser(user db.User) model.User {
	return model.User{ID: user.ID,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.Time,
		UpdatedAt:    user.UpdatedAt.Time,
	}
}
