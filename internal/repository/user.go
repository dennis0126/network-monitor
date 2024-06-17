package repository

import (
	"context"
	"github.com/dennis0126/network-monitor/internal/db"
	"github.com/dennis0126/network-monitor/internal/model"
)

type UserRepository struct {
	ctx     context.Context
	queries *db.Queries
}

func NewUserRepository(ctx context.Context, queries *db.Queries) UserRepository {
	return UserRepository{ctx: ctx, queries: queries}
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

func unmarshallDbUser(user db.User) model.User {
	return model.User{ID: user.ID,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.Time,
		UpdatedAt:    user.UpdatedAt.Time,
	}
}
