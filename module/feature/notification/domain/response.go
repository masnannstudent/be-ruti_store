package domain

import (
	"ruti-store/module/entities"
	"time"
)

type NotificationResponse struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	OrderID   string    `json:"order_id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func ResponseArrayNotificationUser(data []*entities.NotificationModels) []*NotificationResponse {
	res := make([]*NotificationResponse, 0)

	for _, notify := range data {
		notifyRes := &NotificationResponse{
			ID:        notify.ID,
			UserID:    notify.UserID,
			OrderID:   notify.OrderID,
			Title:     notify.Title,
			Message:   notify.Message,
			CreatedAt: notify.CreatedAt,
		}
		res = append(res, notifyRes)
	}

	return res
}
