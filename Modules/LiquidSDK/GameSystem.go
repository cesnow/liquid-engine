package LiquidSDK

import (
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
)

type IGameSystem interface {
	RunCommand(*CmdCommand) interface{}
	RunDirectCommand(*CmdCommand) interface{}
	RunHttpCommand(*CmdCommand) interface{}
	RunRpcCommand(cmd *LiquidRpc.ReqCmd) interface{}
}

type GameSystem struct {
	functionDict map[string]func(string, IGameRequest) interface{}
	drtFunctionDict  map[string]func(string, IGameRequest) interface{}
	httpFunctionDict map[string]func(IGameRequest) interface{}
}

func (gameSystem *GameSystem) RunCommand(data *CmdCommand) interface{} {
	if opFunc, opFuncExist := gameSystem.functionDict[*data.CmdName]; opFuncExist {
		return opFunc(*data.LiquidId, &GameRequest{CmdData: data.CmdData})
	}
	return nil
}

func (gameSystem *GameSystem) RunDirectCommand(data *CmdCommand) interface{} {
	RequestData := &GameRequest{CmdData: data.CmdData}
	if data.LiquidId == nil {
		emptyLiquidId := ""
		data.LiquidId = &emptyLiquidId
	}
	if opFunc, opFuncExist := gameSystem.drtFunctionDict[*data.CmdName]; opFuncExist {
		return opFunc(*data.LiquidId, RequestData)
	}
	return nil
}

func (gameSystem *GameSystem) RunHttpCommand(data *CmdCommand) interface{} {
	RequestData := &GameRequest{CmdData: data.CmdData}
	if httpFunc, httpFuncExist := gameSystem.httpFunctionDict[*data.CmdName]; httpFuncExist {
		return httpFunc(RequestData)
	}
	return nil
}

func (gameSystem *GameSystem) RunRpcCommand(data *LiquidRpc.ReqCmd) interface{} {
	searchDic := gameSystem.functionDict
	if data.Direct {
		searchDic = gameSystem.drtFunctionDict
	}
	if opFunc, opFuncExist := searchDic[data.CmdName]; opFuncExist {
		Req := &GameRequest{CmdData: nil}
		_ = json.Unmarshal(data.CmdData, &Req.CmdData)
		return opFunc(data.UserID, Req)
	} else {
		if httpFunc, httpFuncExist := gameSystem.httpFunctionDict[data.CmdName]; httpFuncExist {
			return httpFunc(&GameRequest{CmdData: data.CmdData})
		}
	}
	return nil
}

func (gameSystem *GameSystem) Register(operator string, f func(string, IGameRequest) interface{}) {
	if gameSystem.functionDict == nil {
		gameSystem.functionDict = make(map[string]func(string, IGameRequest) interface{})
	}
	gameSystem.functionDict[operator] = f
	Logger.SysLog.Debugf("[Engine][OperatorRegister] `%s` Registered", operator)
}

func (gameSystem *GameSystem) RegisterDirect(operator string, f func(string, IGameRequest) interface{}) {
	if gameSystem.drtFunctionDict == nil {
		gameSystem.drtFunctionDict = make(map[string]func(string, IGameRequest) interface{})
	}
	gameSystem.drtFunctionDict[operator] = f
	Logger.SysLog.Debugf("[Engine][OperatorRegisterDirect] `%s` Registered", operator)
}

func (gameSystem *GameSystem) RegisterHttp(operator string, f func(IGameRequest) interface{}) {
	if gameSystem.httpFunctionDict == nil {
		gameSystem.httpFunctionDict = make(map[string]func(IGameRequest) interface{})
	}
	gameSystem.httpFunctionDict[operator] = f
	Logger.SysLog.Debugf("[Engine][OperatorHttpRegister] `%s` Registered", operator)
}
