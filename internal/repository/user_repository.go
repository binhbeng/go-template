package repository

import (
	"context"

	"github.com/binhbeng/goex/internal/db/sqlc"
)

type UserRepository struct {
	db sqlc.Querier
}

func NewUserRepository(db sqlc.Querier) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (m *UserRepository) GetUserById(id int64) (sqlc.User, error) {
	user, err := m.db.GetUserById(context.Background(), int64(id))
	if err != nil {
		return sqlc.User{}, err
	}
	return user, err
}

func (m *UserRepository) GetUserByUsername(username string) (sqlc.User, error) {
	user, err := m.db.GetUserByUsername(context.Background(), username)
	if err != nil {
		return sqlc.User{}, err
	}
	return user, err
}

func (m *UserRepository) UpdateProfile(ctx context.Context, input sqlc.UpdateUserParams) (sqlc.User, error) {
	user, err := m.db.UpdateUser(ctx, sqlc.UpdateUserParams{})
	if err != nil {
		return sqlc.User{}, err
	}
	return user, nil
}