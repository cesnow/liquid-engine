package foundation

import (
	"github.com/cesnow/liquid-engine/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Routers(gin *gin.Engine) {

	gin.GET("/", Root)
	gin.GET("/@", RootKey)

	FoundationRouters := gin.Group("@")
	FoundationRouters.GET("/api/:CmdId/:CmdName", RouteApiDirect)
	FoundationRouters.POST("/api/login", RouteApiLogin)
	FoundationRouters.Use(middlewares.VerifyToken())
	{
		FoundationRouters.POST("/api/:CmdId/:CmdName", RouteApiCommand)
	}
	FoundationRouters.Use(middlewares.GetLiquidData())
	{
		FoundationRouters.POST("/register", RouteRegister)
		FoundationRouters.POST("/login", RouteLogin)
		FoundationRouters.POST("/verify", RouteVerify)
		FoundationRouters.POST("/bind", RouteBind)
		FoundationRouters.POST("/auth", RouteAuth)
		FoundationRouters.POST("/command", RouteCommand)
		FoundationRouters.POST("/direct", RouteDirect)
	}

}
