package LiquidSDK

import (
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
)

type IGameSystem interface {
	RunCommand(*CmdCommand) interface{}
	RunDirectCommand(*CmdCommand) interface{}
	RunRpcCommand(*LiquidRpc.RpcCmdCommand) interface{}
}

type GameSystem struct {
	ILiquidSystem,
	functionDict map[string]func(string, interface{}) interface{}
	drtFunctionDict map[string]func(string, interface{}) interface{}
}

func (gameSystem *GameSystem) RunCommand(data *CmdCommand) interface{} {
	if opFunc, opFuncExist := gameSystem.functionDict[*data.CmdName]; opFuncExist {
		return opFunc(*data.LiquidId, data.CmdData)
	}
	return nil
}

func (gameSystem *GameSystem) RunDirectCommand(data *CmdCommand) interface{} {
	if opFunc, opFuncExist := gameSystem.drtFunctionDict[*data.CmdName]; opFuncExist {
		return opFunc(*data.LiquidId, data.CmdData)
	}
	return nil
}

func (gameSystem *GameSystem) RunRpcCommand(data *LiquidRpc.RpcCmdCommand) interface{} {
	searchDic := gameSystem.functionDict
	if data.Direct {
		searchDic = gameSystem.drtFunctionDict
	}
	if opFunc, opFuncExist := searchDic[data.CmdName]; opFuncExist {
		var CmdData interface{}
		unmarshalErr := json.Unmarshal(data.CmdData, &CmdData)
		if unmarshalErr != nil {
			return nil
		}
		return opFunc(data.UserID, CmdData)
	}
	return nil
}

func (gameSystem *GameSystem) Register(operator string, f func(string, interface{}) interface{}) {
	if gameSystem.functionDict == nil {
		gameSystem.functionDict = make(map[string]func(string, interface{}) interface{})
	}
	gameSystem.functionDict[operator] = f
	Logger.SysLog.Debugf("[Engine][OperatorRegister] `%s` Registered", operator)
}

func (gameSystem *GameSystem) RegisterDirect(operator string, f func(interface{}) interface{}) {
	if gameSystem.drtFunctionDict == nil {
		gameSystem.drtFunctionDict = make(map[string]func(interface{}) interface{})
	}
	gameSystem.drtFunctionDict[operator] = f
	Logger.SysLog.Debugf("[Engine][OperatorRegisterDirect] `%s` Registered", operator)
}
