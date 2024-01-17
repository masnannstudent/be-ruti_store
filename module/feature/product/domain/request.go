package domain

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Price       uint64   `json:"price" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Stock       uint64   `json:"stock" validate:"required"`
	Discount    uint64   `json:"discount"`
	Weight      uint64   `json:"weight" validate:"required"`
	CategoryID  []uint64 `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	ID          uint64   `json:"id"`
	Name        string   `json:"name" `
	Price       uint64   `json:"price" `
	Description string   `json:"description" `
	Stock       uint64   `json:"stock"`
	Discount    uint64   `json:"discount"`
	Weight      uint64   `json:"weight"`
	CategoryID  []uint64 `json:"category_id"`
}

type AddPhotoProductRequest struct {
	ProductID uint64 `form:"product_id" json:"product_id" validate:"required"`
	Photo     string `form:"photo" json:"photo"`
}
