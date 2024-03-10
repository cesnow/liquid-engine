package system

import (
	"github.com/cesnow/liquid-engine/internal/feature-register"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/gin-gonic/gin"
)

type EchoSystem struct {
	LiquidSDK.CommandSDK
}

func NewEchoSystem() LiquidSDK.CommandSystem {
	echo := new(EchoSystem)
	echo.Register("echo", echo.Echo)
	echo.Register("ping", echo.Ping)
	echo.RegisterDirect("direct", echo.Direct)
	return echo
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

func init() {
	feature_register.AddFeature("echo", NewEchoSystem())
}
