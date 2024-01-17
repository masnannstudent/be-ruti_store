package entities

import "time"

type NotificationModels struct {
	ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
	UserID    uint64     `gorm:"column:user_id" json:"user_id"`
	OrderID   string     `gorm:"column:order_id" json:"order_id"`
	Title     string     `gorm:"column:title" json:"title"`
	Message   string     `gorm:"column:message" json:"message"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	User      UserModels `gorm:"foreignKey:UserID" json:"user"`
}

func (NotificationModels) TableName() string {
	return "notification"
}
