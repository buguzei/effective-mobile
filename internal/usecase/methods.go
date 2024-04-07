package usecase

import (
	"context"
	"github.com/buguzei/effective-mobile/internal/models"
	"github.com/buguzei/effective-mobile/internal/repo"
	"github.com/buguzei/effective-mobile/pkg/logging"
)

type UseCase struct {
	repo repo.IRepo
	l    logging.Logger
}

func NewUseCase(repo repo.IRepo) *UseCase {
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("uc")

	return &UseCase{repo: repo, l: logger}
}

func (uc UseCase) NewCars(ctx context.Context, cars []models.Car) error {
	const funcName = "NewCars"

	for _, car := range cars {
		peopleID, err := uc.repo.GetPeopleByFullName(ctx, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic)
		if err != nil {
			uc.l.Error("error of scanning", logging.Fields{
				"error": err,
				"func":  funcName,
			})

			return err
		}

		if peopleID == nil {
			id, err := uc.repo.NewPeople(ctx, car.Owner)
			if err != nil {
				uc.l.Error("error of adding people", logging.Fields{
					"error": err,
					"func":  funcName,
				})

				return err
			}

			peopleID = &id
		}

		car.Owner.ID = *peopleID

		err = uc.repo.NewCar(ctx, car)
		if err != nil {
			uc.l.Error("adding new car failed", logging.Fields{
				"error": err,
				"func":  funcName,
			})

			return err
		}
	}

	return nil
}

func (uc UseCase) DeleteCar(ctx context.Context, regNum string) error {
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

func (uc UseCase) UpdateCar(ctx context.Context, updates models.Car, regNum string) error {
	const funcName = "UpdateCar"

	oldCar, err := uc.repo.GetCarByRegNum(ctx, regNum)
	if err != nil {
		uc.l.Error("error of getting car by regNum", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return err
	}

	owner, err := uc.repo.GetPeopleByID(ctx, oldCar.Owner.ID)
	if err != nil {
		uc.l.Error("error of getting people by ID", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return err
	}

	oldCar.Owner = *owner

	updatedCar := validateUpdate(*oldCar, updates)

	updatedID, err := uc.repo.GetPeopleByFullName(ctx, updatedCar.Owner.Name, updatedCar.Owner.Surname, updatedCar.Owner.Patronymic)
	if err != nil {
		uc.l.Error("error of getting people by full name", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return err
	}

	if updatedID == nil {
		id, err := uc.repo.NewPeople(ctx, updates.Owner)
		if err != nil {
			uc.l.Error("error of adding new people", logging.Fields{
				"error": err,
				"func":  funcName,
			})

			return err
		}

		updatedID = &id
	}

	updatedCar.Owner.ID = *updatedID

	err = uc.repo.UpdateCar(ctx, updatedCar, regNum)
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

	if newCar.Owner.Name != "" {
		oldCar.Owner.Name = newCar.Owner.Name
	}

	if newCar.Owner.Surname != "" {
		oldCar.Owner.Surname = newCar.Owner.Surname
	}

	if newCar.Owner.Patronymic != "" {
		oldCar.Owner.Patronymic = newCar.Owner.Patronymic
	}

	return oldCar
}

func (uc UseCase) GetCars(ctx context.Context, filters models.Car) ([]models.Car, error) {
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
