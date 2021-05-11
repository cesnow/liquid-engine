package Game

import (
	"github.com/cesnow/LiquidEngine/Modules/LiquidModule"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
)

type EchoSystem struct {
	LiquidSDK.GameSystem
}

func NewEchoSystem() LiquidSDK.IGameSystem {
	echo := new(EchoSystem)
	echo.Register("echo", echo.Echo)
	echo.Register("ping", echo.Ping)
	echo.RegisterDirect("direct", echo.Direct)
	return echo
}

func (echo *EchoSystem) Echo(LiquidID string, CmdData LiquidSDK.IGameRequest) interface{} {
	return CmdData
}

func (echo *EchoSystem) Ping(LiquidID string, CmdData LiquidSDK.IGameRequest) interface{} {
	return gin.H{"result": "pong"}
}

func (echo *EchoSystem) Direct(LiquidID string, CmdData LiquidSDK.IGameRequest) interface{} {
	return gin.H{"result": "direct"}
}

func init() {
	LiquidModule.AddFeature("echo", NewEchoSystem())
}
