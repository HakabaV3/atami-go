package model

import (
	"github.com/jinzhu/gorm"
)

// User is a model
type User struct {
	gorm.Model
	Name string
}
