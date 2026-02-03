package repository

import (
	"context"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/model/entity"
	"github.com/binhbeng/goex/internal/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (m *UserRepository) TableName() string {
	return "users"
}

func (m *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: tx,
	}
}

func (m *UserRepository) Update(ctx context.Context, id int, user entity.User) (entity.User, error) {
	if err := m.DB.WithContext(ctx).Save(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (m *UserRepository) GetUserById(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	if err := m.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (m *UserRepository) GetListUser(ctx context.Context, filter dto.QueryUsersInput) ([]entity.User, error) {
	var users []entity.User

	query := m.DB.WithContext(ctx).Model(&entity.User{})

	err := query.Scopes(utils.Paginate(filter.PageOptionsDto)).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
