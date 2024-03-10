package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthReadiness(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, "")
}

func HealthPing(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, "System Served")
}
