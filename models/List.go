package models

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Title  string
	UserID uint
	Tasks  []Task
}
