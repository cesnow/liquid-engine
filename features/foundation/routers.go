package foundation

import (
	"github.com/cesnow/liquid-engine/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Routers(gin *gin.Engine) {

	gin.GET("/", Root)
	gin.GET("/@", RootKey)

	ApiFoundationRouters := gin.Group("@")
	ApiFoundationRouters.GET("/api/:FeatureId/:CmdName", RouteApiDirect)
	ApiFoundationRouters.POST("/api/login", RouteApiLogin)
	ApiFoundationRouters.Use(middlewares.VerifyToken())
	{
		ApiFoundationRouters.POST("/api/:FeatureId/:CmdName", RouteApiCommand)
	}

	CommandFoundationRouters := gin.Group("@")
	CommandFoundationRouters.Use(middlewares.GetLiquidData())
	{
		CommandFoundationRouters.POST("/register", RouteRegister)
		CommandFoundationRouters.POST("/login", RouteLogin)
		CommandFoundationRouters.POST("/verify", RouteVerify)
		CommandFoundationRouters.POST("/bind", RouteBind)
		CommandFoundationRouters.POST("/auth", RouteAuth)
		CommandFoundationRouters.POST("/command", RouteCommand)
		CommandFoundationRouters.POST("/direct", RouteDirect)
	}

}
