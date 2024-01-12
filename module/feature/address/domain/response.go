package domain

import (
	"ruti-store/module/entities"
	"time"
)

type AddressResponse struct {
	ID           uint64     `json:"id"`
	UserID       uint64     `json:"user_id"`
	AcceptedName string     `json:"accepted_name"`
	Phone        string     `json:"phone"`
	ProvinceID   string     `json:"province_id"`
	ProvinceName string     `json:"province_name"`
	CityID       string     `json:"city_id"`
	CityName     string     `json:"city_name"`
	Address      string     `json:"address"`
	IsPrimary    bool       `json:"is_primary"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func ResponseArrayAddresses(data []*entities.AddressModels) []*AddressResponse {
	res := make([]*AddressResponse, 0)

	for _, address := range data {
		addressRes := &AddressResponse{
			ID:           address.ID,
			UserID:       address.UserID,
			AcceptedName: address.AcceptedName,
			Phone:        address.Phone,
			ProvinceID:   address.ProvinceID,
			ProvinceName: address.ProvinceName,
			CityID:       address.CityID,
			CityName:     address.CityName,
			Address:      address.Address,
			IsPrimary:    address.IsPrimary,
			CreatedAt:    address.CreatedAt,
			UpdatedAt:    address.UpdatedAt,
			DeletedAt:    address.DeletedAt,
		}
		res = append(res, addressRes)
	}

	return res
}
