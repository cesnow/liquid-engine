package foundation

import (
	"github.com/cesnow/liquid-engine/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Routers(gin *gin.Engine) {

	gin.GET("/", Root)

	ApiFoundationRouters := gin.Group("/api")
	ApiFoundationRouters.GET("key", RootKey)
	ApiFoundationRouters.GET(":FeatureId/:CmdName", RouteApiDirect)
	ApiFoundationRouters.POST("login", RouteApiLogin)
	ApiFoundationRouters.POST(":FeatureId/:CmdName", RouteApiCommand)

	CommandFoundationRouters := gin.Group("@")
	CommandFoundationRouters.GET("", RootKey)
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
