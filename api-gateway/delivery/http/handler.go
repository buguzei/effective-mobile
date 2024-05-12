package http

import (
	"encoding/json"
	"fmt"
	_ "github.com/buguzei/effective-mobile/api-gateway/docs"
	"github.com/buguzei/effective-mobile/pkg/logging"
	pb "github.com/buguzei/effective-mobile/pkg/protos/gen/auth"
	"google.golang.org/grpc"
	"net/http"
)

const (
	requestURL = "/" // write your API URL here
)

type Handler struct {
	l logging.Logger
}

func NewHandler() *Handler {
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("handler")

	return &Handler{l: logger}
}

func (h Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var reqBody signInRequest
	var respBody signInResponse

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cli pb.AuthClient

	grpcRequest := &pb.SignInRequest{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	}

	grpcResp, err := cli.SignIn(r.Context(), grpcRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("GRPCRESP:", grpcResp.GetRefreshToken(), grpcResp.GetAccessToken())

	respBody.Tokens.Refresh = grpcResp.GetRefreshToken()
	respBody.Tokens.Access = grpcResp.GetAccessToken()

	bRespBody, err := json.Marshal(respBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bRespBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var reqBody signUpRequest
	var respBody signUpResponse

	fmt.Println("i am hereeee")
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	grpcRequest := &pb.SignUpRequest{
		Email:    reqBody.Email,
		Username: reqBody.Username,
		Password: reqBody.Password,
	}

	var opts []grpc.DialOption

	conn, err := grpc.Dial("localhost:8071", opts...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer conn.Close()

	cli := pb.NewAuthClient(conn)

	grpcResp, err := cli.SignUp(r.Context(), grpcRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Tokens.Refresh = grpcResp.GetRefreshToken()
	respBody.Tokens.Access = grpcResp.GetAccessToken()

	fmt.Println("GRPCRESP:", grpcResp.GetRefreshToken(), grpcResp.GetAccessToken())

	bRespBody, err := json.Marshal(respBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bRespBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var reqBody refreshTokenRequest
	var respBody refreshTokenResponse

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cli pb.AuthClient

	grpcRequest := &pb.RefreshRequest{
		Email:        reqBody.Email,
		RefreshToken: reqBody.RefreshToken,
	}

	grpcResp, err := cli.Refresh(r.Context(), grpcRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Tokens.Refresh = grpcResp.GetRefreshToken()
	respBody.Tokens.Access = grpcResp.GetAccessToken()

	bRespBody, err := json.Marshal(respBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bRespBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary GetCars
// @Description get cars
// @Accept  json
// @Produce  json
// @Param regNum query string false "regNum"
// @Param model query string false "model"
// @Param mark query string false "mark"
// @Param name query string false "name"
// @Param surname query string false "surname"
// @Param patronymic query string false "patronymic"
// @Success 200 {object} getCarsResponse
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /cars/get [get]

func (h Handler) GetCars(w http.ResponseWriter, r *http.Request) {
	//const funcName = "GetCars"
	//
	//var respBody getCarsResponse
	//
	//r.URL.Query().Get("regNum")
	//
	//var car models.Car
	//
	//car.RegNum = r.URL.Query().Get("regNum")
	//car.Model = r.URL.Query().Get("model")
	//car.Mark = r.URL.Query().Get("mark")
	//car.Owner.Name = r.URL.Query().Get("name")
	//car.Owner.Surname = r.URL.Query().Get("surname")
	//car.Owner.Patronymic = r.URL.Query().Get("patronymic")
	//
	//h.l.Debug(fmt.Sprintf("got car's filters: %v", car), logging.Fields{
	//	"func": funcName,
	//})
	//
	//cars, err := h.uc.GetCars(r.Context(), car)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//h.l.Debug(fmt.Sprintf("cars to response:: %v", cars), logging.Fields{
	//	"func": funcName,
	//})
	//
	//respBody.Cars = cars
	//
	//bRespBody, err := json.Marshal(respBody)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//_, err = w.Write(bRespBody)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
}

// @Summary DeleteCar
// @Description delete car
// @Accept  json
// @Produce  json
// @Param input body deleteCarRequest true "deleting car"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /cars/delete [delete]

func (h Handler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	//const funcName = "DeleteCar"
	//
	//var reqBody deleteCarRequest
	//
	//err := json.NewDecoder(r.Body).Decode(&reqBody)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//h.l.Info(fmt.Sprintf("regNums od deleting car equals to %v", reqBody.RegNum), logging.Fields{
	//	"func": funcName,
	//})
	//
	//err = h.uc.DeleteCar(r.Context(), reqBody.RegNum)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
}

// @Summary NewCars
// @Description new cars
// @Accept  json
// @Produce  json
// @Param input body newCarRequest true "h"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /cars/new [post]

func (h Handler) NewCars(w http.ResponseWriter, r *http.Request) {
	//const funcName = "NewCars"
	//
	//var reqBody newCarRequest
	//var cars []models.Car
	//
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//err = json.Unmarshal(body, &reqBody)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//h.l.Debug(fmt.Sprintf("regNums to add: %v", reqBody.RegNums), logging.Fields{
	//	"func": funcName,
	//})
	//
	//for _, regNum := range reqBody.RegNums {
	//	bRefName := []byte(regNum)
	//
	//	req, err := http.NewRequest(http.MethodGet, requestURL, bytes.NewBuffer(bRefName))
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	res, err := http.DefaultClient.Do(req)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	if res.Status == "200" {
	//		resBody, err := io.ReadAll(res.Body)
	//		if err != nil {
	//			http.Error(w, err.Error(), http.StatusInternalServerError)
	//			return
	//		}
	//
	//		var car models.Car
	//		err = json.Unmarshal(resBody, &car)
	//		if err != nil {
	//			http.Error(w, err.Error(), http.StatusInternalServerError)
	//			return
	//		}
	//
	//		cars = append(cars, car)
	//	}
	//
	//	if res.Status == "400" {
	//		h.l.Info(fmt.Sprintf("400 status code got from API"), logging.Fields{
	//			"func": funcName,
	//		})
	//
	//		w.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//
	//	if res.Status == "500" {
	//		h.l.Info(fmt.Sprintf("500 status code got from API"), logging.Fields{
	//			"func": funcName,
	//		})
	//
	//		w.WriteHeader(http.StatusServiceUnavailable)
	//		return
	//	}
	//}
	//
	//err = h.uc.NewCars(r.Context(), cars)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
}

// @Summary UpdateCar
// @Description update car
// @Accept  json
// @Param input body updateCarRequest true "h"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /cars/update [put]

func (h Handler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	//const funcName = "UpdateCar"
	//
	//var bodyReq updateCarRequest
	//
	//err := json.NewDecoder(r.Body).Decode(&bodyReq)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//h.l.Debug(fmt.Sprintf("gor regNum: %s; got updates: %v", bodyReq.RegNum, bodyReq.Updates), logging.Fields{
	//	"func": funcName,
	//})
	//
	//err = h.uc.UpdateCar(r.Context(), bodyReq.Updates, bodyReq.RegNum)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
}
