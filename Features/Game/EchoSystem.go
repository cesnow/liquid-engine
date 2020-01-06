package Game

import (
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
)

type EchoSystem struct {
	LiquidSDK.GameSystem
}

func NewEchoSystem() *EchoSystem {
	echo := new(EchoSystem)
	echo.LoadDefaultOperator()
	return echo
}

func (echo *EchoSystem) LoadDefaultOperator() {
	echo.Register("echo", echo.Echo)
	echo.Register("ping", echo.Ping)
}

func (echo *EchoSystem) Echo(LiquidID string, CmdData interface{}) interface{} {
	return CmdData
}

func (echo *EchoSystem) Ping(LiquidID string, CmdData interface{}) interface{} {
	return "Pong"
}
