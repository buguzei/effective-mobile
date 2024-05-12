package usecase

import (
	"context"
	"github.com/buguzei/effective-mobile/auth/internal/models"
	"github.com/buguzei/effective-mobile/pkg/token"
)

type AuthUCI interface {
	SignIn(ctx context.Context, user models.User) (token.Pair, error)
	SignUp(ctx context.Context, user models.User) (token.Pair, error)
	Refresh(ctx context.Context, email string, refreshToken string) (token.Pair, error)
}
