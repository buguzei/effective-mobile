package repo

import "github.com/buguzei/effective-mobile/internal/models"

type IRepo interface {
	ICarRepo
	IPeopleRepo
}

type ICarRepo interface {
	NewCar(car models.Car) error
	GetCarsByFilters(filters models.Car) ([]models.Car, error)
	GetCarByRegNum(regNum string) (*models.Car, error)
	UpdateCar(updates models.Car, regNum string) error
	DeleteCar(regNum string) error
}

type IPeopleRepo interface {
	NewPeople(people models.People) (int, error)
	GetPeopleByFullName(name, surname, patronymic string) (*int, error)
	GetPeopleByID(id int) (*models.People, error)
}
