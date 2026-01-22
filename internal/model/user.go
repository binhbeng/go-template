package model

import (
	"context"

	"github.com/binhbeng/goex/internal/dto"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	BaseModel
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
	Username  string                `json:"username"`
	Password  string                `json:"-"`
	Email     string                `json:"email"`
}

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

func (m *UserRepository) GetUserById(ctx context.Context, id int64) (User, error) {
	var user User
	if err := m.DB().WithContext(ctx).First(&user, id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (m *UserRepository) GetListUser(ctx context.Context, filter dto.QueryUsersInput) ([]User, error) {
	var users []User

	query := m.DB().WithContext(ctx).Model(&User{})

	err := query.Scopes(m.Paginate(filter.PageOptionsDto)).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
