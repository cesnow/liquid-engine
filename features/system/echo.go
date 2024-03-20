package system

import (
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/cesnow/liquid-engine/options"
	"github.com/gin-gonic/gin"
)

type EchoSystem struct {
	LiquidSDK.CommandSDK
}

func NewEchoSystem() LiquidSDK.CommandSystem {
	echo := new(EchoSystem)
	echo.Register("echo", echo.Echo, options.Command().WithHttpSupport())
	echo.Register("ping", echo.Ping, options.Command().WithHttpSupport())
	echo.RegisterDirect("direct", echo.Direct, options.Command().WithHttpSupport())
	echo.Register("error-test", echo.ErrorTest, options.Command().WithHttpSupport())

	echo.RegisterRouter(echo.UseRouter)

	return echo
}

func (echo *EchoSystem) UseRouter(router *gin.RouterGroup) {
	// if same command name, the UseRouter will be overwritten
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "pong2"})
	})
}

func (echo *EchoSystem) Echo(LiquidID string, CmdData LiquidSDK.CommandRequest) interface{} {
	return CmdData
}

func (echo *EchoSystem) Ping(LiquidID string, CmdData LiquidSDK.CommandRequest) interface{} {
	return gin.H{"result": "pong"}
}

func (echo *EchoSystem) Direct(LiquidID string, CmdData LiquidSDK.CommandRequest) interface{} {
	return gin.H{"result": "direct"}
}

func (echo *EchoSystem) ErrorTest(LiquidID string, CmdData LiquidSDK.CommandRequest) interface{} {
	return LiquidSDK.CmdErrorResponse{Error: "TEST_ERROR"}
}

//
//func init() {
//	feature_register.AddFeature("echo", NewEchoSystem())
//}
