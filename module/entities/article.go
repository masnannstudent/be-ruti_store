package entities

import "time"

type ArticleModels struct {
	ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
	Title     string     `gorm:"column:title;type:VARCHAR(255)" json:"title"`
	Content   string     `gorm:"column:content;type:TEXT" json:"content"`
	Author    string     `gorm:"column:author;type:VARCHAR(255)" json:"author"`
	Photo     string     `gorm:"column:photo;type:VARCHAR(255)" json:"photo"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

func (ArticleModels) TableName() string {
	return "article"
}
