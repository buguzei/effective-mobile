package usecase

import (
	"context"
	"github.com/buguzei/effective-mobile/auth/internal/token"
)

type AuthUCI interface {
	SignIn(ctx context.Context, email, password string) (token.Pair, error)
	SignUp(ctx context.Context, username, email, password string) (token.Pair, error)
}
