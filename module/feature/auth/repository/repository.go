package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/auth/domain"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepositoryInterface {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) GetUsersByEmail(email string) (*entities.UserModels, error) {
	var user entities.UserModels
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) CreateUser(req *entities.UserModels) (*entities.UserModels, error) {
	if err := r.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}
