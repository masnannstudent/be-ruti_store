package domain

import "debtomate/module/entities"

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token"`
}

func LoginFormatter(user *entities.UserModels, accessToken string) LoginResponse {
	return LoginResponse{
		User: UserResponse{
			Name:  user.Name,
			Email: user.Email,
		},
		AccessToken: accessToken,
	}
}
