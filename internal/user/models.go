package userModels

import "github.com/lib/pq"

type User struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type Landlord struct {
	Name         string        `json:"name"`
	Surname      string        `json:"surname"`
	Patronymic   string        `json:"patronymic"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Post         string        `json:"post"`
	Places       pq.Int64Array `json:"places"`
	Legal_entity string        `json:"legal_entity"`
	INN          string        `json:"INN"`
}
