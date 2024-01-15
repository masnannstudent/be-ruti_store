package domain

type EditProfileRequest struct {
	Name         string `form:"name" json:"name"`
	Phone        string `form:"phone" json:"phone"`
	PhotoProfile string `form:"photo" json:"photo_profile"`
}
