package image

import (
	// "log"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/Dooor/atami-go/db"
)

// Use is a function
func Use(r *gin.RouterGroup) {
	r.Use()
	{
		r.GET("/image/all", getAllImages)
		r.GET("/image/search", getSearchImages)
		r.POST("/image", createImage)
	}
}

// getAllImages is a function
func getAllImages(c *gin.Context) {
	images := db.AllImages()
	c.JSON(200, images)
}

// getSearchImages is a function
func getSearchImages(c *gin.Context) {
	keyword := c.Query("q")
	if len(keyword) == 0 {
		getAllImages(c)
		return
	}
	keywords := strings.Split(keyword, ",")
	images := db.SearchImages(keywords)
	c.JSON(200, images)
}

// createImage is a function
func createImage(c *gin.Context) {
	url := c.Query("url")
	image := db.CreateImage(url)
	c.JSON(201, image)
}
