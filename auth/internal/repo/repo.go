package repo

import (
	"context"
	"github.com/buguzei/effective-mobile/auth/internal/models"
)

type UserRepo interface {
	IsUserExist(ctx context.Context, email string) (bool, error)
	NewUser(ctx context.Context, user models.User) (int, error)
	VerifyEmailAndPass(ctx context.Context, email, password string) (models.User, error)
	FindUserByEmail(ctx context.Context, email string) (models.User, error)
}

type RefreshRepo interface {
	SetRefresh(ctx context.Context, refreshToken string, id int) error
	GetRefresh(ctx context.Context, id int) (string, error)
}
