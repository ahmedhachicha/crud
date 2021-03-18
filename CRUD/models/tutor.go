package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tutor struct {
	gorm.Model
	FirstName string    `json:firstName`
	LastName  string    `json:lastname`
	Email     string    `json:email`
	BirthDate time.Time `json:birthDate`
	Country   string    `json:country`
	Biography string `json:Biography`
	Courses []Course  `gorm:"foreignKey:UserRefer"`
}
