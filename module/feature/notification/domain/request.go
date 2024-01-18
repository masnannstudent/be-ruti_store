package domain

type CreateNotificationRequest struct {
	Title   string `json:"title"`
	UserID  uint64 `json:"user_id"`
	OrderID string `json:"order_id"`
	Message string `json:"message"`
}
