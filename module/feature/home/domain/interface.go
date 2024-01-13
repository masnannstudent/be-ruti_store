package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type HomeRepositoryInterface interface {
	CreateCarousel(carousel *entities.CarouselModels) (*entities.CarouselModels, error)
	GetCarouselById(carouselID uint64) (*entities.CarouselModels, error)
	UpdateCarousel(carouselID uint64, updatedCarousel *entities.CarouselModels) error
	DeleteCarousel(carouselID uint64) error
}

type HomeServiceInterface interface {
	CreateCarousel(req *CreateCarouselRequest) (*entities.CarouselModels, error)
	GetCarouselById(carouselID uint64) (*entities.CarouselModels, error)
	UpdateCarousel(carouselID uint64, updatedCarousel *entities.CarouselModels) error
	DeleteCarousel(carouselID uint64) error
}

type HomeHandlerInterface interface {
	CreateCarousel(c *fiber.Ctx) error
	GetCarouselByID(c *fiber.Ctx) error
}
