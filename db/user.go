package db

import (
	"github.com/Dooor/atami-go/model"
)

// Users is a function
func Users() interface{} {
	db := Gdb()
	users := []model.User{}
	return db.Find(&users).Value
}

// User is a function
func User(id int64) interface{} {
	db := Gdb()
	return db.First(&model.User{}, id).Value
}

// CreateUser is a function
func CreateUser(param map[string]string) {
	db := Gdb()
	user := model.User{
		Name: param["name"],
	}
	db.Create(&user)
}

// DeleteUser is a function
func DeleteUser(id int64) {
	db := Gdb()
	db.Exec("DELETE FROM users WHERE id = ?;", []int64{id})
}
