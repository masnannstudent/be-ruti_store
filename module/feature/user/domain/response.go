package domain

import (
	"ruti-store/module/entities"
	"time"
)

type UserResponse struct {
	ID           uint64    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Phone        string    `json:"phone"`
	Name         string    `json:"name"`
	PhotoProfile string    `json:"photo_profile"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
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

type UserEditProfileResponse struct {
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	PhotoProfile string `json:"photo_profile"`
}

func UserEditProfileFormatter(user *entities.UserModels) *UserEditProfileResponse {
	result := &UserEditProfileResponse{
		Phone:        user.Phone,
		Name:         user.Name,
		PhotoProfile: user.PhotoProfile,
	}
	return result
}

func ResponseArrayUser(data []*entities.UserModels) []*UserResponse {
	res := make([]*UserResponse, 0)

	for _, userItem := range data {
		userRes := &UserResponse{
			ID:           userItem.ID,
			Email:        userItem.Email,
			Phone:        userItem.Phone,
			Name:         userItem.Name,
			PhotoProfile: userItem.PhotoProfile,
			Role:         userItem.Role,
			CreatedAt:    userItem.CreatedAt,
		}
		res = append(res, userRes)
	}

	return res
}
