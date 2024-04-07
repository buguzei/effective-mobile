package models

type People struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
