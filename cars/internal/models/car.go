package models

type Car struct {
	RegNum  string `json:"regNum"`
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	OwnerID int    `json:"-"`
}
