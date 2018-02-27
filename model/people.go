package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Persons []Person

type PersonRepository struct{}

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Person{})
}

func NewPersonRepository() PersonRepository {
	return PersonRepository{}
}

func (p PersonRepository) LoadAllPersons() Persons {
	var peoples = Persons{}
	db.Find(&peoples)
	return peoples
}

func (p PersonRepository) LoadSpecificPerson(id int) Persons {
	var people = Persons{}
	if err = db.Where("id = ?", id).First(&people).Error; err != nil {
		panic(err)
	}
	return people
}

func (p PersonRepository) CreateNewPerson(firstname, lastname string) {
	var people = Person{FirstName: firstname, LastName: lastname}
	db.NewRecord(people)
	db.Create(&people)
	db.Save(&people)
}

func (p PersonRepository) UpdatePerson(id int, firstname, lastname string) Person {
	var people = Person{ID: id}
	db.Model(&people).Where("id = ?", id).Updates(Person{FirstName: firstname, LastName: lastname})
	return people
}

func (p PersonRepository) DeletePerson(id int) {
	var people = Person{ID: id}
	db.Where("id = ?", id).Delete(&people)
}
