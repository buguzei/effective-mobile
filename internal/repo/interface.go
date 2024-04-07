package repo

import (
	"context"
	"github.com/buguzei/effective-mobile/internal/models"
)

type IRepo interface {
	ICarRepo
	IPeopleRepo
}

type ICarRepo interface {
	NewCar(ctx context.Context, car models.Car) error
	GetCarsByFilters(ctx context.Context, filters models.Car) ([]models.Car, error)
	GetCarByRegNum(ctx context.Context, regNum string) (*models.Car, error)
	UpdateCar(ctx context.Context, updates models.Car, regNum string) error
	DeleteCar(ctx context.Context, regNum string) error
}

type IPeopleRepo interface {
	NewPeople(ctx context.Context, people models.People) (int, error)
	GetPeopleByFullName(ctx context.Context, name, surname, patronymic string) (*int, error)
	GetPeopleByID(ctx context.Context, id int) (*models.People, error)
}
