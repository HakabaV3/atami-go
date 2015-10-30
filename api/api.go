package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Dooor/atami-go/api/v1"
)

// Init is an exported function that
// define routing for api
func Init() {
	r := gin.Default()
	v1.Init(r.Group("api/v1"))
	r.Run(":8080")
}
