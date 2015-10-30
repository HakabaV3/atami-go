package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/Dooor/atami-go/config"
)

// OnError is a struct
type OnError struct {
	code int
	message string
}

var gdb *gorm.DB

// Gdb is an exported function that
// define database
func Gdb() *gorm.DB {
	if gdb == nil {
		db, _ := gorm.Open(config.ADAPTER, config.USERNAME + config.PASSWORD + ":@/" + config.DATABASE + "?charset=" + config.CHARSET + "&parseTime=True")
		gdb = &db
		gdb.DB()
	}
	return gdb
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
