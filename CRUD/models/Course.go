package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Title 		string `json:title`
	Level 		string `json:level`
	Description string `json:title`
	UserRefer   uint
}