package usecase

import (
	"context"
	"github.com/buguzei/effective-mobile/cars/internal/models"
)

type CarsUCI interface {
	NewCar(ctx context.Context, cars models.Car) error
	DeleteCar(ctx context.Context, regNum string) error
	UpdateCar(ctx context.Context, updates models.Car) error
	GetCars(ctx context.Context, filters models.Car) ([]models.Car, error)
}
