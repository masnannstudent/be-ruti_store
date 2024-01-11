package entities

import "time"

type UserModels struct {
	ID           uint64          `gorm:"column:id;primaryKey" json:"id"`
	Email        string          `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Password     string          `gorm:"column:password;type:VARCHAR(255)" json:"password"`
	Phone        string          `gorm:"column:phone;type:VARCHAR(255)" json:"phone"`
	Role         string          `gorm:"column:role;type:VARCHAR(255)" json:"role"`
	Name         string          `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	PhotoProfile string          `gorm:"column:photo_profile;type:VARCHAR(255)" json:"photo_profile"`
	DeviceToken  string          `gorm:"column:device_token;type:VARCHAR(255)" json:"device_token"`
	CreatedAt    time.Time       `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time       `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt    *time.Time      `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Address      []AddressModels `gorm:"foreignKey:UserID" json:"addresses"`
	//Reviews      []ReviewModels  `gorm:"foreignKey:UserID" json:"reviews"`
}

type AddressModels struct {
	ID           uint64     `gorm:"column:id;primaryKey" json:"id"`
	UserID       uint64     `gorm:"column:user_id" json:"user_id"`
	AcceptedName string     `gorm:"column:accepted_name;type:VARCHAR(255)" json:"accepted_name"`
	Phone        string     `gorm:"column:phone;type:VARCHAR(255)" json:"phone"`
	ProvinceID   string     `gorm:"column:province_id;type:VARCHAR(255)" json:"province_id"`
	ProvinceName string     `gorm:"column:province_name;type:VARCHAR(255)" json:"province_name"`
	CityID       string     `gorm:"column:city_id;type:VARCHAR(255)" json:"city_id"`
	CityName     string     `gorm:"column:city_name;type:VARCHAR(255)" json:"city_name"`
	Address      string     `gorm:"column:address;type:VARCHAR(255)" json:"address"`
	IsPrimary    bool       `gorm:"column:is_primary;type:BOOLEAN" json:"is_primary"`
	CreatedAt    time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (UserModels) TableName() string {
	return "users"
}

func (AddressModels) TableName() string {
	return "address"
}
