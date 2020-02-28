package Foundation

import (
	"github.com/gin-gonic/gin"
)

func Routers(gin *gin.Engine) {
	gin.GET("/", Root)
	gin.GET("/@", RootKey)
}
