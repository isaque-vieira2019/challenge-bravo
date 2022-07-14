package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	Name  string
	Value float64
}
