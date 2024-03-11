package route_health

import (
	"github.com/gin-gonic/gin"
)

func Routers(gin *gin.Engine) {
	HealthRouters := gin.Group("/health")
	HealthRouters.Use()
	{
		HealthRouters.GET("/readiness", Readiness)
		HealthRouters.GET("/ping", Ping)
	}
}
