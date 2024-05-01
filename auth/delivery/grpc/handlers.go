package grpc

import (
	"context"
	"github.com/buguzei/effective-mobile/auth/internal/usecase"
	"github.com/buguzei/effective-mobile/pkg/logging"
	pb "github.com/buguzei/effective-mobile/protos/gen/auth"
)

type AuthHandler struct {
	l logging.Logger
	pb.UnimplementedAuthServer
	uc usecase.AuthUCI
}

func New(uc usecase.AuthUCI) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h AuthHandler) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	const f = "AuthHandler.SignIn"

	pair, err := h.uc.SignIn(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		h.l.Error("error of sign in", logging.Fields{
			"error": err,
			"func":  f,
		})

		return nil, err
	}

	resp := &pb.SignInResponse{
		RefreshToken: pair.RefreshToken,
		AccessToken:  pair.AccessToken,
	}

	return resp, nil
}

func (h AuthHandler) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	const f = "AuthHandler.SignUp"

	pair, err := h.uc.SignUp(ctx, in.GetUsername(), in.GetEmail(), in.GetPassword())
	if err != nil {
		h.l.Error("error of sign up", logging.Fields{
			"error": err,
			"func":  f,
		})
		return nil, err
	}

	resp := &pb.SignUpResponse{
		RefreshToken: pair.RefreshToken,
		AccessToken:  pair.AccessToken,
	}

	return resp, nil
}
