package repo

import (
	"context"
	"github.com/buguzei/effective-mobile/auth/internal/models"
)

type UserRepo interface {
	NewUser(ctx context.Context, user models.User) (int, error)
	FindUserByEmail(ctx context.Context, email, password string) (models.User, int, error)
}

type RefreshRepo interface {
	SetRefresh(ctx context.Context, refreshToken string, id int)
}
