package domain

type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required,min=6,noSpace"`
}

type RegisterRequest struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required,min=6,noSpace"`
	Name     string `form:"name" json:"name" validate:"required"`
	Phone    string `form:"phone" json:"phone" validate:"required,noSpace"`
	Role     string `json:"role"`
}
