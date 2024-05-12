package postgres

import (
	"context"
	"database/sql"
	"errors"
	errors2 "github.com/buguzei/effective-mobile/auth/internal/errors"
	"github.com/buguzei/effective-mobile/auth/internal/models"
	"github.com/buguzei/effective-mobile/pkg/logging"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
	l  logging.Logger
}

func New() *Postgres {
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("auth.postgres")

	db, err := sql.Open("postgres", "host=postgres-auth port=5432 user=buguzei password=password dbname=auth sslmode=disable")
	if err != nil {
		logger.Fatal("error of opening db", logging.Fields{
			"err": err,
		})
		return nil
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal("error of pinging db", logging.Fields{
			"err": err,
		})
	}

	return &Postgres{DB: db, l: logger}
}

func (p Postgres) IsUserExist(ctx context.Context, email string) (bool, error) {
	var exists bool

	row := p.DB.QueryRowContext(ctx, "SELECT EXISTS(SELECT FROM users WHERE email = $1);", email)
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (p Postgres) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	row := p.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1;", email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.New(errors2.ErrUserNotFound)
		}

		return models.User{}, err
	}

	return user, nil
}

func (p Postgres) NewUser(ctx context.Context, user models.User) (int, error) {
	res := p.DB.QueryRowContext(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id;", user.Name, user.Email, user.Password)

	var id int
	err := res.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p Postgres) VerifyEmailAndPass(ctx context.Context, email, password string) (models.User, error) {
	var user models.User

	row := p.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1 AND password = $2;", email, password)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.New(errors2.ErrUserNotFound)
		}

		return models.User{}, err
	}

	return user, nil
}
