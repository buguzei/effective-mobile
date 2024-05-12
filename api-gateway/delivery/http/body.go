package http

import (
	"github.com/buguzei/effective-mobile/api-gateway/internal/models"
	"github.com/buguzei/effective-mobile/pkg/token"
)

// auth microservice bodies

type signUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type signUpResponse struct {
	Tokens token.Pair `json:"tokens"`
}

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	Tokens token.Pair `json:"tokens"`
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
	Email        string `json:"email"`
}

type refreshTokenResponse struct {
	Tokens token.Pair `json:"tokens"`
}

// cars microservice bodies

type newCarRequest struct {
	RegNums []string `json:"regNums"`
}

type getCarsRequest struct {
	Filters models.Car `json:"filters"`
}

type getCarsResponse struct {
	Cars []models.Car `json:"cars"`
}

type updateCarRequest struct {
	RegNum  string     `json:"regNum"`
	Updates models.Car `json:"updates"`
}

type deleteCarRequest struct {
	RegNum string `json:"regNum"`
}
