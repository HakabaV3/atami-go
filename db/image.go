package db

import (
	// "log"
	// "github.com/jinzhu/gorm"
	"github.com/Dooor/atami-go/model"
)

// AllImages is a function
func AllImages() interface{} {
	db := Gdb()
	images := []model.Image{}
	return db.Find(&images).Value
}

// SearchImages is a function
func SearchImages(keywords []string) interface{} {
	db := Gdb()
	images := []model.Image{}
	rows, _ := db.Raw("SELECT images.* FROM images INNER JOIN image_tags ON image_tags.image_id = images.id WHERE image_tags.tag_id = ANY (SELECT tags.id FROM tags WHERE name LIKE (?));", keywords).Rows()
	for rows.Next() {
		image := model.Image{}
		rows.Scan(&image.ID, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt, &image.URL)
		images = append(images, image)
	}
	return images
}

// CreateImage is a function
func CreateImage(url string) interface{} {
	db := Gdb()
	image := model.Image{URL: url}
	if result := db.Where("url = ?", url).Find(&model.Image{}); result.Error == nil {
		return result.Value
	}
	return db.Create(&image).Value
}
