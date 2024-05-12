package usecase

import (
	"context"
	"github.com/buguzei/effective-mobile/cars/internal/models"
	"github.com/buguzei/effective-mobile/cars/internal/repo"
	"github.com/buguzei/effective-mobile/pkg/logging"
)

type CarsUC struct {
	repo repo.ICarRepo
	l    logging.Logger
}

func NewUseCase(repo repo.ICarRepo) *CarsUC {
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("uc")

	return &CarsUC{repo: repo, l: logger}
}

func (uc CarsUC) NewCar(ctx context.Context, car models.Car) error {
	const funcName = "CarUC.NewCar"

	err := uc.repo.NewCar(ctx, car)
	if err != nil {
		uc.l.Error("adding new car failed", logging.Fields{
			"error": err,
			"func":  funcName,
		})
		return err
	}

	return nil
}

func (uc CarsUC) DeleteCar(ctx context.Context, regNum string) error {
	const funcName = "DeleteCar"

	err := uc.repo.DeleteCar(ctx, regNum)
	if err != nil {
		uc.l.Error("error of deleting car", logging.Fields{
			"error": err,
			"func":  funcName,
		})
		return err
	}

	return nil
}

func (uc CarsUC) UpdateCar(ctx context.Context, updates models.Car) error {
	const funcName = "UpdateCar"

	oldCar, err := uc.repo.GetCarByRegNum(ctx, updates.RegNum)
	if err != nil {
		uc.l.Error("error of getting car by regNum", logging.Fields{
			"error": err,
			"func":  funcName,
		})
		return err
	}

	updatedCar := validateUpdate(*oldCar, updates)

	err = uc.repo.UpdateCar(ctx, updatedCar, updates.RegNum)
	if err != nil {
		uc.l.Error("error of updating car", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return err
	}

	return nil
}

func validateUpdate(oldCar, newCar models.Car) models.Car {
	if newCar.RegNum != "" {
		oldCar.RegNum = newCar.RegNum
	}

	if newCar.Mark != "" {
		oldCar.Mark = newCar.Mark
	}

	if newCar.Model != "" {
		oldCar.Model = newCar.Model
	}

	return oldCar
}

func (uc CarsUC) GetCars(ctx context.Context, filters models.Car) ([]models.Car, error) {
	const funcName = "GetCars"

	cars, err := uc.repo.GetCarsByFilters(ctx, filters)
	if err != nil {
		uc.l.Error("error of getting car using filters", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return nil, err
	}

	return cars, nil
}
