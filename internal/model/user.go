package model

type User struct {
	ContainsDeleteBaseModel
	Username string `json:"username"` 
	Password string `json:"-"`
	Email    string `json:"email"`   
}

func NewUser() *User {
	return &User{}
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) GetUserById(id uint) *User {
	if err := m.DB().First(m, id).Error; err != nil {
		return nil
	}
	return m
}