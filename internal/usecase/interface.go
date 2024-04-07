package usecase

import "github.com/buguzei/effective-mobile/internal/models"

type IUseCase interface {
	NewCars(cars []models.Car) error
	DeleteCar(regNum string) error
	UpdateCar(updates models.Car, regNums string) error
	GetCars(filters models.Car) ([]models.Car, error)
}
