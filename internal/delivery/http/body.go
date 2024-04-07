package http

import "github.com/buguzei/effective-mobile/internal/models"

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
