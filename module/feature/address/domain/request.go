package domain

type CreateAddressRequest struct {
	UserID       uint64 `form:"user_id" json:"user_id"`
	AcceptedName string `form:"accepted_name" json:"accepted_name" validate:"required"`
	Phone        string `form:"phone" json:"phone" validate:"required"`
	ProvinceID   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
	CityID       string `json:"city_id"`
	CityName     string `json:"city_name"`
	Address      string `form:"address" json:"address" validate:"required"`
	IsPrimary    bool   `form:"is_primary" json:"is_primary"`
}

type UpdateAddressRequest struct {
	UserID       uint64 `form:"user_id" json:"user_id"`
	AcceptedName string `form:"accepted_name" json:"accepted_name"`
	Phone        string `form:"phone" json:"phone"`
	ProvinceID   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
	CityID       string `json:"city_id"`
	CityName     string `json:"city_name"`
	Address      string `form:"address" json:"address"`
	IsPrimary    bool   `form:"is_primary" json:"is_primary"`
}
