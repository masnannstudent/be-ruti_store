package entities

import "time"

type ReviewModels struct {
	ID          uint64              `gorm:"column:id;primaryKey" json:"id"`
	UserID      uint64              `gorm:"column:user_id" json:"user_id"`
	User        UserModels          `gorm:"foreignKey:UserID" json:"user"`
	ProductID   uint64              `gorm:"column:product_id" json:"product_id"`
	Rating      uint64              `gorm:"column:rating" json:"rating"`
	Description string              `gorm:"column:description;type:text" json:"description"`
	CreatedAt   time.Time           `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt   time.Time           `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt   *time.Time          `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Photos      []ReviewPhotoModels `gorm:"foreignKey:ReviewID" json:"photos"`
}

type ReviewPhotoModels struct {
	ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
	ReviewID  uint64     `gorm:"column:review_id" json:"review_id"`
	ImageURL  string     `gorm:"column:url;type:varchar(255)" json:"url"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (ReviewModels) TableName() string {
	return "reviews"
}

func (ReviewPhotoModels) TableName() string {
	return "review_photos"
}
