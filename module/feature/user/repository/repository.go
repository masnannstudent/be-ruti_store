package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/user/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(addressID uint64) (*entities.UserModels, error) {
	var users *entities.UserModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", addressID).First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) EditProfile(userID uint64, req *entities.UserModels) error {
	var user *entities.UserModels
	if err := r.db.Model(&user).Where("id = ?", userID).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetTotalUserItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.UserModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *UserRepository) GetPaginatedUsers(page, pageSize int) ([]*entities.UserModels, error) {
	var users []*entities.UserModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Offset(offset).
		Limit(pageSize).
		Find(&users).
		Order("created_at DESC").
		Error; err != nil {
		return nil, err
	}

	return users, nil
}
