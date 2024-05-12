package http

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func (h Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// auth routes
	r.HandleFunc("/sign-in", h.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", h.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/refresh", h.Refresh).Methods(http.MethodPost)

	// cars routes
	carRouter := r.PathPrefix("/cars").Subrouter()
	carRouter.Use(h.VerifyUser)

	carRouter.HandleFunc("/cars/new", h.NewCars).Methods(http.MethodPost)
	carRouter.HandleFunc("/cars/delete", h.DeleteCar).Methods(http.MethodDelete)
	carRouter.HandleFunc("/cars/update", h.UpdateCar).Methods(http.MethodPut)
	carRouter.HandleFunc("/cars/get", h.GetCars).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8072/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	return r
}
