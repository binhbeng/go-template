package form

type UpdateUser struct {
	Email string `form:"email" json:"email" binding:"required"`
}

func NewUpdateUserForm() *UpdateUser {
	return &UpdateUser{}
}