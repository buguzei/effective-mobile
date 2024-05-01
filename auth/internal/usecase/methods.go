package usecase

import (
	"context"
	"errors"
	errors2 "github.com/buguzei/effective-mobile/auth/internal/errors"
	"github.com/buguzei/effective-mobile/auth/internal/models"
	"github.com/buguzei/effective-mobile/auth/internal/repo"
	"github.com/buguzei/effective-mobile/auth/internal/token"
	"github.com/buguzei/effective-mobile/pkg/logging"
)

type AuthUC struct {
	l  logging.Logger
	rr repo.RefreshRepo
	ur repo.UserRepo
}

func New(rr repo.RefreshRepo, ur repo.UserRepo) *AuthUC {
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("auth.uc")

	return &AuthUC{l: logger, rr: rr, ur: ur}
}

func (uc AuthUC) SignIn(ctx context.Context, email, password string) (token.Pair, error) {
	const f = "AuthUC.SignIn"

	_, id, err := uc.ur.FindUserByEmail(ctx, email, password)
	if err != nil {
		uc.l.Error("error of UserRepo.FindUserByEmail", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	newRefresh, err := token.NewRefreshToken()
	if err != nil {
		uc.l.Error("error of Token.NewRefreshToken", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	uc.rr.SetRefresh(ctx, newRefresh, id)

	newAccess, err := token.NewAccessToken(id)
	if err != nil {
		uc.l.Error("error of NewAccessToken", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	newPair := token.Pair{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}

	return newPair, nil
}

func (uc AuthUC) SignUp(ctx context.Context, username, email, password string) (token.Pair, error) {
	const f = "AuthUC.SignUp"

	user := models.User{
		Email:    email,
		Password: password,
		Name:     username,
	}

	_, _, err := uc.ur.FindUserByEmail(ctx, email, password)
	if err != nil && !errors.Is(err, errors.New(errors2.ErrUserNotFound)) {
		uc.l.Error("error of UserRepo.FindUserByEmail", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	id, err := uc.ur.NewUser(ctx, user)
	if err != nil {
		uc.l.Error("error of UserRepo.NewUser", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	newRefresh, err := token.NewRefreshToken()
	if err != nil {
		uc.l.Error("error of Token.NewRefreshToken", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	uc.rr.SetRefresh(ctx, newRefresh, id)

	newAccess, err := token.NewAccessToken(id)
	if err != nil {
		uc.l.Error("error of NewAccessToken", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	newPair := token.Pair{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}

	return newPair, nil
}
