package model

import (
	"github.com/jinzhu/gorm"
)

// Image is a model
type Image struct {
	gorm.Model
	URL string
}
