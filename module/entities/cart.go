package entities

type CartModels struct {
	ID        uint64        `gorm:"column:id;primaryKey" json:"id"`
	UserID    uint64        `gorm:"column:user_id" json:"user_id"`
	ProductID uint64        `gorm:"column:product_id" json:"product_id"`
	Size      string        `gorm:"column:size;type:VARCHAR(255)" json:"size"`
	Quantity  uint64        `gorm:"column:quantity" json:"quantity"`
	User      UserModels    `gorm:"foreignKey:UserID" json:"user" `
	Product   ProductModels `gorm:"foreignKey:ProductID" json:"product" `
}

func (CartModels) TableName() string {
	return "carts"
}
