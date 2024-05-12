package usecase

import (
	"context"
	"errors"
	errors2 "github.com/buguzei/effective-mobile/auth/internal/errors"
	"github.com/buguzei/effective-mobile/auth/internal/models"
	"github.com/buguzei/effective-mobile/auth/internal/repo"
	"github.com/buguzei/effective-mobile/pkg/logging"
	"github.com/buguzei/effective-mobile/pkg/token"
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

func (uc AuthUC) Refresh(ctx context.Context, email string, refreshToken string) (token.Pair, error) {
	const f = "AuthUC.Refresh"

	user, err := uc.ur.FindUserByEmail(ctx, email)
	if err != nil {
		uc.l.Error("error of UserRepo.FindUserByEmail", logging.Fields{
			"error": err,
			"func":  f,
		})

		return token.Pair{}, err
	}

	actualRefreshToken, err := uc.rr.GetRefresh(ctx, user.ID)
	if err != nil {
		uc.l.Error("error of UserRepo.GetRefresh", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	if refreshToken != actualRefreshToken {
		return token.Pair{}, errors.New("refresh token mismatch")
	}

	newRefresh, err := token.NewRefreshToken()
	if err != nil {
		uc.l.Error("error of Token.NewRefreshToken", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	err = uc.rr.SetRefresh(ctx, newRefresh, user.ID)
	if err != nil {
		uc.l.Error("error of SetRefresh", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	newAccess, err := token.NewAccessToken(user.ID)
	if err != nil {
		uc.l.Error("error of NewAccessToken", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	newPair := token.Pair{
		Access:  newAccess,
		Refresh: newRefresh,
	}

	return newPair, nil
}

func (uc AuthUC) SignIn(ctx context.Context, user models.User) (token.Pair, error) {
	const f = "AuthUC.SignIn"

	user, err := uc.ur.VerifyEmailAndPass(ctx, user.Email, user.Password)
	if err != nil {
		uc.l.Error("error of UserRepo.VerifyEmailAndPass", logging.Fields{
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

	err = uc.rr.SetRefresh(ctx, newRefresh, user.ID)
	if err != nil {
		uc.l.Error("error of SetRefresh", logging.Fields{
			"error": err,
			"func":  f,
		})
	}

	newAccess, err := token.NewAccessToken(user.ID)
	if err != nil {
		uc.l.Error("error of NewAccessToken", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	newPair := token.Pair{
		Access:  newAccess,
		Refresh: newRefresh,
	}

	return newPair, nil
}

func (uc AuthUC) SignUp(ctx context.Context, user models.User) (token.Pair, error) {
	const f = "AuthUC.SignUp"

	exists, err := uc.ur.IsUserExist(ctx, user.Email)
	if err != nil {
		uc.l.Error("error of UserRepo.IsUserExist", logging.Fields{
			"error": err,
			"func":  f,
		})
	}

	if exists {
		err = errors.New(errors2.ErrUserAlreadyExists)

		uc.l.Error("error of UserRepo.IsUserExist", logging.Fields{
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

	err = uc.rr.SetRefresh(ctx, newRefresh, id)
	if err != nil {
		uc.l.Error("error of Redis.SetRefresh", logging.Fields{})
		return token.Pair{}, err
	}

	newAccess, err := token.NewAccessToken(id)
	if err != nil {
		uc.l.Error("error of NewAccessToken", logging.Fields{
			"error": err,
			"func":  f,
		})
		return token.Pair{}, err
	}

	newPair := token.Pair{
		Access:  newAccess,
		Refresh: newRefresh,
	}

	return newPair, nil
}
