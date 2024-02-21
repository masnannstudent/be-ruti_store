package domain

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Price       uint64   `json:"price" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Discount    uint64   `json:"discount"`
	CategoryID  []uint64 `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	ID          uint64   `json:"id"`
	Name        string   `json:"name"`
	Price       uint64   `json:"price"`
	Description string   `json:"description"`
	Discount    uint64   `json:"discount"`
	CategoryID  []uint64 `json:"category_id"`
}

type AddPhotoProductRequest struct {
	ProductID uint64 `form:"product_id" json:"product_id" validate:"required"`
	Photo     string `form:"photo" json:"photo"`
}

type UpdatePhotoProductRequest struct {
	ProductID uint64 `form:"product_id" json:"product_id" validate:"required"`
	Photo     string `form:"photo" json:"photo"`
}

type CreateVariantRequest struct {
	ProductID uint64 `json:"product_id"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	Weight    uint64 `json:"weight"`
	Stock     uint64 `json:"stock"`
}

type UpdateStatusRequest struct {
	ProductID uint64 `json:"product_id" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
