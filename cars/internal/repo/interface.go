package repo

import (
	"context"
	models2 "github.com/buguzei/effective-mobile/cars/internal/models"
)

type IRepo interface {
	ICarRepo
	IPeopleRepo
}

type ICarRepo interface {
	NewCar(ctx context.Context, car models2.Car) error
	GetCarsByFilters(ctx context.Context, filters models2.Car) ([]models2.Car, error)
	GetCarByRegNum(ctx context.Context, regNum string) (*models2.Car, error)
	UpdateCar(ctx context.Context, updates models2.Car, regNum string) error
	DeleteCar(ctx context.Context, regNum string) error
}

type IPeopleRepo interface {
	NewPeople(ctx context.Context, people models2.People) (int, error)
	GetPeopleByFullName(ctx context.Context, name, surname, patronymic string) (*int, error)
	GetPeopleByID(ctx context.Context, id int) (*models2.People, error)
}
