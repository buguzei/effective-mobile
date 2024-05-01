package http

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func (h Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/sign-in", h.NewCars).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", h.DeleteCar).Methods(http.MethodPost)

	r.Use(h.VerifyToken)

	r.HandleFunc("/cars/new", h.NewCars).Methods(http.MethodPost)
	r.HandleFunc("/cars/delete", h.DeleteCar).Methods(http.MethodDelete)
	r.HandleFunc("/cars/update", h.UpdateCar).Methods(http.MethodPut)
	r.HandleFunc("/cars/get", h.GetCars).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8087/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	return r
}
