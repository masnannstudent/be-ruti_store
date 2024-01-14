package domain

import (
	"ruti-store/module/entities"
)

type UserResponse struct {
	ID           uint64 `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	PhotoProfile string `json:"photo_profile"`
}

func UserFormatter(user *entities.UserModels) *UserResponse {
	result := &UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		Password:     "",
		Phone:        user.Phone,
		Name:         user.Name,
		PhotoProfile: user.PhotoProfile,
	}
	return result
}
