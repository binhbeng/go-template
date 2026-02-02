package repository

import (
	"context"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/model/entity"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) *UserRepository {
	return &UserRepository{
		Repository: r,
	}
}

func (m *UserRepository) TableName() string {
	return "users"
}

func (m *UserRepository) GetUserById(ctx context.Context, id int64) (entity.User, error) {
	var user entity.User
	if err := m.DB().WithContext(ctx).First(&user, id).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (m *UserRepository) GetListUser(ctx context.Context, filter dto.QueryUsersInput) ([]entity.User, error) {
	var users []entity.User

	query := m.DB().WithContext(ctx).Model(&entity.User{})

	err := query.Scopes(m.Paginate(filter.PageOptionsDto)).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
