package repository

import (
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/order/domain"
	"ruti-store/utils/payment"
)

type OrderRepository struct {
	db   *gorm.DB
	snap snap.Client
	core coreapi.Client
}

func NewOrderRepository(db *gorm.DB, snap snap.Client, core coreapi.Client) domain.OrderRepositoryInterface {
	return &OrderRepository{
		db:   db,
		snap: snap,
		core: core,
	}
}

func (r *OrderRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Model(&entities.OrderModels{}).Count(&totalItems).Where("deleted_at IS NULL").Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *OrderRepository) GetPaginatedOrders(page, pageSize int) ([]*entities.OrderModels, error) {
	var orders []*entities.OrderModels

	offset := (page - 1) * pageSize

	err := r.db.Offset(offset).Limit(pageSize).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Preload("User").
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) CreateSnap(orderID, name, email string, totalAmountPaid uint64) (*snap.Response, error) {
	snapClient := r.snap
	result, err := payment.CreatePaymentRequest(snapClient, orderID, totalAmountPaid, name, email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *OrderRepository) CheckTransaction(orderID string) (domain.Status, error) {
	var status domain.Status
	transactionStatusResp, err := r.core.CheckTransaction(orderID)
	if err != nil {
		return domain.Status{}, err
	} else {
		if transactionStatusResp != nil {
			status = payment.TransactionStatus(transactionStatusResp)
			return status, nil
		}
	}
	return domain.Status{}, err
}

func (r *OrderRepository) CreateOrder(newOrder *entities.OrderModels) (*entities.OrderModels, error) {
	err := r.db.Create(newOrder).Error
	if err != nil {
		return nil, err
	}
	return newOrder, nil
}

func (r *OrderRepository) GetOrderByID(orderID string) (*entities.OrderModels, error) {
	var order *entities.OrderModels
	if err := r.db.Where("id = ? AND deleted_at IS NULL", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) UpdatePayment(orderID, orderStatus, paymentStatus string) error {
	var orders entities.OrderModels
	if err := r.db.Model(&orders).Where("id = ?", orderID).Updates(map[string]interface{}{
		"order_status":   orderStatus,
		"payment_status": paymentStatus,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) CreateCart(newCart *entities.CartModels) (*entities.CartModels, error) {
	err := r.db.Create(newCart).Error
	if err != nil {
		return nil, err
	}
	return newCart, nil
}

func (r *OrderRepository) GetCartItem(userID, productID uint64) (*entities.CartModels, error) {
	var cartItem *entities.CartModels
	if err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error; err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (r *OrderRepository) UpdateCartItem(cartItem *entities.CartModels) error {
	if err := r.db.Save(&cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) DeleteCartItem(cartItemID uint64) error {
	if err := r.db.Where("id = ?", cartItemID).Delete(&entities.CartModels{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetCartByID(cartID uint64) (*entities.CartModels, error) {
	var carts *entities.CartModels
	if err := r.db.Where("id = ?", cartID).First(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *OrderRepository) GetCartByUserID(userID uint64) ([]*entities.CartModels, error) {
	var carts []*entities.CartModels
	if err := r.db.Preload("Product.Photos").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *OrderRepository) AcceptOrder(orderID, orderStatus string) error {
	if err := r.db.Model(&entities.OrderModels{}).
		Where("id = ?", orderID).
		Update("order_status", orderStatus).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) UpdateOrderStatus(orderID, orderStatus string) error {
	var orders *entities.OrderModels
	if err := r.db.Where("id = ?", orderID).First(&orders).Error; err != nil {
		return err
	}

	if err := r.db.Model(&orders).Updates(map[string]interface{}{
		"order_status": orderStatus,
	}).Error; err != nil {
		return err
	}

	return nil
}
