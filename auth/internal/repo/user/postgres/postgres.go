package postgres

import (
	"context"
	"database/sql"
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

	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres dbname=auth sslmode=disable")
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

func (p Postgres) NewUser(ctx context.Context, user models.User) (int, error) {
	res, err := p.DB.ExecContext(ctx, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (p Postgres) FindUserByEmail(ctx context.Context, email, password string) (models.User, int, error) {
	var id int
	var user models.User

	row := p.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1 AND password = $2", email, password)

	err := row.Scan(&id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, 0, err
	}

	return user, id, nil
}
