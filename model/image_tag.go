package model

import (
	"github.com/jinzhu/gorm"
)

// ImageTag model
type ImageTag struct {
	gorm.Model
	ImageID uint
	TagID   uint
}
