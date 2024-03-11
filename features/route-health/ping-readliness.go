package route_health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Readiness(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, "")
}

func Ping(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, "PONG")
}
