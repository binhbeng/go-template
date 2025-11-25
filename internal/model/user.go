package model

type User struct {
	BaseModelWithSoftDelete
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
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

func (m *UserRepository) GetUserById(id uint) (*User, error) {
	var user User
	if err := m.DB().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
