package usecase

import (
	"context"
	"github.com/buguzei/effective-mobile/cars/internal/models"
)

type IUseCase interface {
	NewCars(ctx context.Context, cars []models.Car) error
	DeleteCar(ctx context.Context, regNum string) error
	UpdateCar(ctx context.Context, updates models.Car, regNums string) error
	GetCars(ctx context.Context, filters models.Car) ([]models.Car, error)
}
