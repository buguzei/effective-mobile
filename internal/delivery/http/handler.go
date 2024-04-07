package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/buguzei/effective-mobile/docs"
	"github.com/buguzei/effective-mobile/internal/models"
	"github.com/buguzei/effective-mobile/internal/usecase"
	"io"
	"net/http"
)

const (
	requestURL = "http://localhost:8089" // write your API URL here
)

type Handler struct {
	uc usecase.IUseCase
}

func NewHandler(uc usecase.IUseCase) *Handler {
	return &Handler{uc: uc}
}

// @Summary GetCars
// @Description get cars
// @Accept  json
// @Produce  json
// @Param input body getCarsRequest true "h"
// @Success 200 {object} getCarsResponse
// @Failure 400 {integer} integer 1
// @Failure 500 {integer} integer 1
// @Router /cars/get [get]
func (h Handler) GetCars(w http.ResponseWriter, r *http.Request) {
	var reqBody getCarsRequest
	var respBody getCarsResponse

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respBody.Cars, err = h.uc.GetCars(reqBody.Filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bRespBody, err := json.Marshal(respBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bRespBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary DeleteCar
// @Description delete car
// @Accept  json
// @Produce  json
// @Param input body deleteCarRequest true "h"
// @Success 200 {integer} integer 1
// @Failure 400 {integer} integer 1
// @Failure 500 {integer} integer 1
// @Router /cars/delete [delete]
func (h Handler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	var reqBody deleteCarRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteCar(reqBody.RegNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary NewCars
// @Description new cars
// @Accept  json
// @Produce  json
// @Param input body newCarRequest true "h"
// @Success 200 {integer} integer 1
// @Failure 400 {integer} integer 1
// @Failure 500 {integer} integer 1
// @Router /cars/new [post]
func (h Handler) NewCars(w http.ResponseWriter, r *http.Request) {
	var reqBody newCarRequest
	var cars []models.Car

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, regNum := range reqBody.RegNums {
		bRefName := []byte(regNum)

		fmt.Println(regNum)

		req, err := http.NewRequest(http.MethodGet, requestURL, bytes.NewBuffer(bRefName))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("client: got response!\n")
		fmt.Printf("client: status code: %d\n", res.StatusCode)

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(resBody))

		var car models.Car
		err = json.Unmarshal(resBody, &car)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cars = append(cars, car)
	}

	err = h.uc.NewCars(cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary UpdateCar
// @Description update car
// @Accept  json
// @Produce  json
// @Param input body updateCarRequest true "h"
// @Success 200 {integer} integer 1
// @Failure 400 {integer} integer 1
// @Failure 500 {integer} integer 1
// @Router /cars/change [put]
func (h Handler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	var bodyReq updateCarRequest

	err := json.NewDecoder(r.Body).Decode(&bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(bodyReq.Updates)

	err = h.uc.UpdateCar(bodyReq.Updates, bodyReq.RegNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
