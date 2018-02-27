package controller

import (
	"../model"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func NewPerson() Person {
	return Person{}
}

func (c Person) LoadAllPersons() interface{} {
	repo := model.NewPersonRepository()
	return repo.LoadAllPersons()
}

func (c Person) LoadSpecificPerson(id int) interface{} {
	repo := model.NewPersonRepository()
	return repo.LoadSpecificPerson(id)
}

func (c Person) CreateNewPerson(firstname, lastname string) {
	repo := model.NewPersonRepository()
	repo.CreateNewPerson(firstname, lastname)
}

func (c Person) UpdatePerson(id int, firstname, lastname string) interface{} {
	repo := model.NewPersonRepository()
	return repo.UpdatePerson(id, firstname, lastname)
}

func (c Person) DeletePerson(id int) {
	repo := model.NewPersonRepository()
	repo.DeletePerson(id)
}
