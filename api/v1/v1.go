package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Dooor/atami-go/api/v1/user"
	"github.com/Dooor/atami-go/api/v1/image"
)

// Init is an exported function that
// define routing for v1
func Init(r *gin.RouterGroup) {
	user.Use(r)
	image.Use(r)
}
