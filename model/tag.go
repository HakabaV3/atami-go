package model

import (
	"github.com/jinzhu/gorm"
)

// Tag is a model
type Tag struct {
	gorm.Model
	Name string
}
