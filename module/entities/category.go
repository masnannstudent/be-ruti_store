package entities

import "time"

type CategoryModels struct {
	ID          uint64           `gorm:"column:id;primaryKey" json:"id"`
	Name        string           `gorm:"column:name" json:"name" `
	Description string           `gorm:"column:description" json:"description"`
	Photo       string           `gorm:"column:photo;type:VARCHAR(255)" json:"photo"`
	CreatedAt   time.Time        `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt   time.Time        `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt   *time.Time       `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Products    []*ProductModels `gorm:"many2many:product_categories;" json:"products,omitempty"`
}

func (CategoryModels) TableName() string {
	return "category"
}
