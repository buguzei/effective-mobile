package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/buguzei/effective-mobile/internal/models"
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

func (p Postgres) NewPeople(people models.People) (int, error) {
	const funcName = "NewPeople"

	res := p.DB.QueryRow("INSERT INTO people(name, surname, patronymic) VALUES (($1), ($2), ($3)) RETURNING id;", people.Name, people.Surname, people.Patronymic)

	var id int
	err := res.Scan(&id)
	if err != nil {
		p.L.Error("error of scanning", logging.Fields{
			"error": err,
			"func":  funcName,
		})
	}

	p.L.Debug(fmt.Sprintf("id of new people: %d", id), logging.Fields{
		"func": funcName,
	})

	return id, nil
}

func (p Postgres) GetPeopleByID(id int) (*models.People, error) {
	const funcName = "GetPeopleByID"

	var people models.People
	people.ID = id

	row := p.DB.QueryRow("SELECT name, surname, patronymic FROM people WHERE id=($1)", id)

	err := row.Scan(&people.Name, &people.Surname, &people.Patronymic)
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

	return &people, nil
}

func (p Postgres) GetPeopleByFullName(name, surname, patronymic string) (*int, error) {
	const funcName = "GetPeopleByFullName"

	var peopleID int

	row := p.DB.QueryRow("SELECT id FROM people WHERE name=($1) and surname=($2) and patronymic=($3)", name, surname, patronymic)

	err := row.Scan(&peopleID)
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

	p.L.Debug(fmt.Sprintf("id of people: %d", peopleID), logging.Fields{
		"func": funcName,
	})

	return &peopleID, nil
}

func (p Postgres) NewCar(car models.Car) error {
	const funcName = "NewCar"

	_, err := p.DB.Exec("INSERT INTO cars(regnum, mark, model, owner_id) VALUES (($1), ($2), ($3), ($4));",
		car.RegNum,
		car.Mark,
		car.Model,
		car.Owner.ID,
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

func (p Postgres) GetCarByRegNum(regNum string) (*models.Car, error) {
	const funcName = "GetCarByRegNum"

	var car models.Car

	row := p.DB.QueryRow("SELECT * FROM cars WHERE regnum=($1);", regNum)

	err := row.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Owner.ID)
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

func (p Postgres) GetCarsByFilters(filters models.Car) ([]models.Car, error) {
	const funcName = "GetCarsByFilters"

	var cars []models.Car

	queryFilter := makeQueryFilter(filters)

	rows, err := p.DB.Query("SELECT regnum, mark, model, name, surname, patronymic FROM cars JOIN people ON cars.owner_id = people.id" + queryFilter)
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
		var car models.Car

		err = rows.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
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

func makeQueryFilter(filters models.Car) string {
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

	if filters.Owner.Name != "" {
		queryFilter += fmt.Sprintf(" name='%s' AND", filters.Owner.Name)
	}

	if filters.Owner.Surname != "" {
		queryFilter += fmt.Sprintf(" surname='%s' AND", filters.Owner.Surname)
	}

	if filters.Owner.Patronymic != "" {
		queryFilter += fmt.Sprintf(" patronymic='%s' AND", filters.Owner.Patronymic)
	}

	return queryFilter[:len(queryFilter)-4]
}

func (p Postgres) UpdateCar(car models.Car, regNum string) error {
	const funcName = "UpdateCar"

	_, err := p.DB.Exec("UPDATE cars SET regnum=($1), mark = ($2), model = ($3), owner_id = ($4) WHERE regnum=($5);",
		car.RegNum,
		car.Mark,
		car.Model,
		car.Owner.ID,
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

func (p Postgres) DeleteCar(regNum string) error {
	const funcName = "DeleteCar"

	_, err := p.DB.Exec("DELETE FROM cars WHERE regnum=($1);", regNum)
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
