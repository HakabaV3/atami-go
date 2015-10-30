package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Dooor/atami-go/model"
	"github.com/Dooor/atami-go/config"
)

func main() {
	db, _ := gorm.Open(config.ADAPTER, config.USERNAME + config.PASSWORD + ":@/" + config.DATABASE + "?charset=" + config.CHARSET + "&parseTime=True")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Image{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.ImageTag{})
}
