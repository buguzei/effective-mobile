package main

import (
	"encoding/json"
	"fmt"
	"github.com/buguzei/effective-mobile/internal/models"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		b, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Println("FROM 222220000: ", err.Error())
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		var bCar []byte

		switch string(b) {
		case "E123EE777":
			bCar, err = json.Marshal(models.Car{
				RegNum: "E123EE777",
				Mark:   "Ferrari",
				Model:  "Sport",
				Owner: models.People{
					Name:       "Ivan",
					Surname:    "Ivanov",
					Patronymic: "Ivanovich",
				},
			})
			if err != nil {
				fmt.Println("FROM 22222222222: ", err.Error())
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		case "E123EE888":
			bCar, err = json.Marshal(models.Car{
				RegNum: "E123EE888",
				Mark:   "Lamba",
				Model:  "Tractor",
				Owner: models.People{
					Name:       "Petr",
					Surname:    "Ivanov",
					Patronymic: "Ivanovich",
				},
			})
			if err != nil {
				fmt.Println("FROM 22222222222: ", err.Error())
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		case "E123EE999":
			bCar, err = json.Marshal(models.Car{
				RegNum: "E123EE999",
				Mark:   "BMW",
				Model:  "M8",
				Owner: models.People{
					Name:       "Semen",
					Surname:    "Ivanov",
					Patronymic: "Ivanovich",
				},
			})
			if err != nil {
				fmt.Println("FROM 22222222222: ", err.Error())
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(bCar)
		if err != nil {
			fmt.Println("FROM 222222222221212121: ", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Println("start listening")
	if err := http.ListenAndServe(":8089", mux); err != nil {
		log.Fatal(err)
	}
}
