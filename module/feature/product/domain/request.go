package domain

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Price       uint64   `json:"price" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Stock       uint64   `json:"stock" validate:"required"`
	Discount    uint64   `json:"discount"`
	CategoryID  []uint64 `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	ID          uint64   `json:"id"`
	Name        string   `json:"name" validate:"required"`
	Price       uint64   `json:"price" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Stock       uint64   `json:"stock" validate:"required"`
	Discount    uint64   `json:"discount"`
	CategoryID  []uint64 `json:"category_id" validate:"required"`
}
