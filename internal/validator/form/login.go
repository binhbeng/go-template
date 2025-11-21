package form

type LoginAuth struct {
	Username string `form:"username" json:"username"  binding:"required,min=5"` 
	Password string `form:"password" json:"password"  binding:"required,min=6"`
}

func NewLoginForm() *LoginAuth {
	return &LoginAuth{}
}
