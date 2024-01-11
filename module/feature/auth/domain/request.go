package domain

type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required,min=6,noSpace"`
}
