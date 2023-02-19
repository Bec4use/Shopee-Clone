package orm

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Address  string
	Birthday string
	Tel      string
}
