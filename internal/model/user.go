package model

import "gorm.io/plugin/soft_delete"

type User struct {
	BaseModel
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index;" json:"-"`
	Username  string                `json:"username"`
	Password  string                `json:"password"`
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

func (m *UserRepository) GetUserById(id int64) (User, error) {
	var user User
	if err := m.DB().First(&user, id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
