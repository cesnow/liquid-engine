package GameFoundation

import (
	"github.com/cesnow/LiquidEngine/Middlewares"
	"github.com/gin-gonic/gin"
)

func Routers(gin *gin.Engine) {
	GameFoundationRouters := gin.Group("@")
	GameFoundationRouters.Use(Middlewares.GetLiquidData())
	{
		GameFoundationRouters.POST("/register", RouteRegister)
		GameFoundationRouters.POST("/login", RouteLogin)
		GameFoundationRouters.POST("/verify", RouteVerify)
		GameFoundationRouters.POST("/bind", RouteBind)
		GameFoundationRouters.POST("/auth", RouteAuth)
		GameFoundationRouters.POST("/command", RouteCommand)
		GameFoundationRouters.POST("/direct", RouteDirect)
	}
}
