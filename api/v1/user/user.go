package user

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/Dooor/atami-go/db"
)

// Use is an exported function that
// define routing for user
func Use(r *gin.RouterGroup) {
	r.Use()
	{
		r.GET("/user", getUsers)
		r.GET("/user/:id", getUser)
		r.POST("/user", postUser)
		r.PATCH("/user/:id", updateUser)
		r.DELETE("/user/:id", deleteUser)
	}
}

func getUsers(c *gin.Context) {
	users := db.Users()
	c.JSON(200, users)
}

func getUser(c *gin.Context) {
	id := c.Params.ByName("id")
	userID, _ := strconv.ParseInt(id, 0, 64)
	user := db.User(userID)
	c.JSON(200, user)
}

func postUser(c *gin.Context) {
	// something happen
}

func updateUser(c *gin.Context) {
	// something happen
}

func deleteUser(c *gin.Context) {
	// something happen
}
