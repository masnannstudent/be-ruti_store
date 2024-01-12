package domain

type Status struct {
	PaymentStatus string
	OrderStatus   string
}

type CreatePaymentRequest struct {
	OrderID         string `json:"order_id"`
	TotalAmountPaid uint64 `json:"total_amount_paid"`
	Name            string `json:"name"`
	Email           string `json:"email"`
}

type CreateOrderRequest struct {
	AddressID uint64 `form:"address_id" json:"address_id" validate:"required"`
	Note      string `form:"note" json:"note"`
	ProductID uint64 `json:"product_id" validate:"required"`
	Quantity  uint64 `json:"quantity" validate:"required"`
}
