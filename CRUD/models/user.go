package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName string    `json:firstName`
	LastName  string    `json:lastname`
	Email     string    `json:email`
	BirthDate time.Time `json:birthDate`
	Country   string    `json:country`
}
