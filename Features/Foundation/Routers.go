package Foundation

import (
	"github.com/gin-gonic/gin"
)

func Routers(gin *gin.Engine) {
	gin.GET("/@", RootKey)
}
