package domain

import "time"

type EditProfileRequest struct {
	Name         string    `form:"name" json:"name"`
	Phone        string    `form:"phone" json:"phone"`
	PhotoProfile string    `form:"photo" json:"photo_profile"`
	Gender       string    `form:"gender" json:"gender"`
	DateOfBirth  time.Time `form:"date_of_birth" json:"date_of_birth"`
}

type CreateChatBotRequest struct {
	Message string `json:"message" validate:"required"`
}
