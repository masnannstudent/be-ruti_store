package domain

import "ruti-store/module/entities"

type UserRepositoryInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
}

type UserServiceInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
}

type UserHandlerInterface interface {
}
