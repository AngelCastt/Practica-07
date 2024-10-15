package database

import "gorm.io/gorm"

type Usuarios struct {
	gorm.Model
	Name  string
	Email string
}
