package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	models2 "github.com/buguzei/effective-mobile/cars/internal/models"
	"github.com/buguzei/effective-mobile/pkg/logging"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
	L  logging.Logger
}

func NewPostgres(dbConn string) (*Postgres, error) {
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("postgres")

	return &Postgres{DB: db, L: logger}, nil
}

func (p Postgres) NewCar(ctx context.Context, car models2.Car) error {
	const funcName = "NewCar"

	_, err := p.DB.ExecContext(ctx, "INSERT INTO cars(regnum, mark, model, owner_id) VALUES (($1), ($2), ($3), ($4));",
		car.RegNum,
		car.Mark,
		car.Model,
		car.OwnerID,
	)
	if err != nil {
		p.L.Error("error of executing a query", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return err
	}

	return nil
}

func (p Postgres) GetCarByRegNum(ctx context.Context, regNum string) (*models2.Car, error) {
	const funcName = "GetCarByRegNum"

	var car models2.Car

	row := p.DB.QueryRowContext(ctx, "SELECT * FROM cars WHERE regnum=($1);", regNum)

	err := row.Scan(&car.RegNum, &car.Mark, &car.Model, &car.OwnerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		p.L.Error("error of scanning", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return nil, err
	}

	p.L.Debug(fmt.Sprintf("car with regNum=%s is %v", regNum, car), logging.Fields{
		"func": funcName,
	})

	return &car, nil
}

func (p Postgres) GetCarsByFilters(ctx context.Context, filters models2.Car) ([]models2.Car, error) {
	const funcName = "GetCarsByFilters"

	var cars []models2.Car

	queryFilter := makeQueryFilter(filters)

	rows, err := p.DB.QueryContext(ctx, "SELECT regnum, mark, model FROM cars"+queryFilter)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		p.L.Error("error of executing a query", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return nil, err
	}

	for rows.Next() {
		var car models2.Car

		err = rows.Scan(&car.RegNum, &car.Mark, &car.Model)
		if err != nil {
			p.L.Error("error of scanning", logging.Fields{
				"error": err,
				"func":  funcName,
			})

			return nil, err
		}

		cars = append(cars, car)
	}

	p.L.Debug(fmt.Sprintf("filters: %v --> res: %v", filters, cars), logging.Fields{
		"func": funcName,
	})

	return cars, nil
}

func makeQueryFilter(filters models2.Car) string {
	var queryFilter = " WHERE"

	if filters.RegNum != "" {
		queryFilter += fmt.Sprintf(" regnum='%s' AND", filters.RegNum)
	}

	if filters.Model != "" {
		queryFilter += fmt.Sprintf(" model='%s' AND", filters.Model)
	}

	if filters.Mark != "" {
		queryFilter += fmt.Sprintf(" mark='%s' AND", filters.Mark)
	}

	return queryFilter[:len(queryFilter)-4]
}

func (p Postgres) UpdateCar(ctx context.Context, car models2.Car, regNum string) error {
	const funcName = "UpdateCar"

	_, err := p.DB.ExecContext(ctx, "UPDATE cars SET regnum=($1), mark = ($2), model = ($3), owner_id = ($4) WHERE regnum=($5);",
		car.RegNum,
		car.Mark,
		car.Model,
		car.OwnerID,
		regNum,
	)
	if err != nil {
		p.L.Error("error of executing a query", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return err
	}

	p.L.Info(fmt.Sprintf("car with regNum=%d was updated successfully", regNum), logging.Fields{
		"func": funcName,
	})

	return nil
}

func (p Postgres) DeleteCar(ctx context.Context, regNum string) error {
	const funcName = "DeleteCar"

	_, err := p.DB.ExecContext(ctx, "DELETE FROM cars WHERE regnum=($1);", regNum)
	if err != nil {
		p.L.Error("error of executing a query", logging.Fields{
			"error": err,
			"func":  funcName,
		})

		return err
	}

	p.L.Info(fmt.Sprintf("car with regNum=%d was deleted successfully", regNum), logging.Fields{
		"func": funcName,
	})

	return nil
}
