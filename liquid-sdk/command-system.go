package LiquidSDK

import (
	"github.com/cesnow/liquid-engine/logger"
	"github.com/cesnow/liquid-engine/options"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommandSystem interface {
	RunCommand(*CmdCommand) interface{}
	RunDirectCommand(*CmdCommand) interface{}
	RunHttpDirectCommand(*gin.Context, *CmdCommand)
	RunHttpCommand(*gin.Context, *CmdCommand)
}

type CommandSDK struct {
	functionDict        map[string]func(string, CommandRequest) interface{}
	drtFunctionDict     map[string]func(string, CommandRequest) interface{}
	httpFunctionDict    map[string]func(*gin.Context, CommandRequest)
	drtHttpFunctionDict map[string]func(*gin.Context, CommandRequest)
}

func (system *CommandSDK) RunCommand(data *CmdCommand) interface{} {
	if opFunc, opFuncExist := system.functionDict[*data.CmdName]; opFuncExist {
		return opFunc(*data.LiquidId, &LiquidRequest{
			LiquidId: data.LiquidId,
			Platform: data.Platform,
			CmdId:    data.CmdId,
			CmdSn:    data.CmdSn,
			CmdName:  data.CmdName,
			CmdData:  data.CmdData,
		})
	}
	return nil
}

func (system *CommandSDK) RunDirectCommand(data *CmdCommand) interface{} {
	RequestData := &LiquidRequest{
		LiquidId: data.LiquidId,
		Platform: data.Platform,
		CmdId:    data.CmdId,
		CmdSn:    data.CmdSn,
		CmdName:  data.CmdName,
		CmdData:  data.CmdData}
	if data.LiquidId == nil {
		emptyLiquidId := ""
		data.LiquidId = &emptyLiquidId
	}
	if opFunc, opFuncExist := system.drtFunctionDict[*data.CmdName]; opFuncExist {
		return opFunc(*data.LiquidId, RequestData)
	}
	return nil
}

func (system *CommandSDK) RunHttpDirectCommand(c *gin.Context, data *CmdCommand) {
	RequestData := &LiquidRequest{
		LiquidId: data.LiquidId,
		Platform: data.Platform,
		CmdId:    data.CmdId,
		CmdSn:    data.CmdSn,
		CmdName:  data.CmdName,
		CmdData:  data.CmdData}
	if httpFunc, httpFuncExist := system.drtHttpFunctionDict[*data.CmdName]; httpFuncExist {
		httpFunc(c, RequestData)
		c.Abort()
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "feature command not found"})
	c.Abort()
	return
}

func (system *CommandSDK) RunHttpCommand(c *gin.Context, data *CmdCommand) {
	RequestData := &LiquidRequest{
		LiquidId: data.LiquidId,
		Platform: data.Platform,
		CmdId:    data.CmdId,
		CmdSn:    data.CmdSn,
		CmdName:  data.CmdName,
		CmdData:  data.CmdData}
	if httpFunc, httpFuncExist := system.httpFunctionDict[*data.CmdName]; httpFuncExist {
		httpFunc(c, RequestData)
		c.Abort()
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "feature command not found"})
	c.Abort()
	return
}

func (system *CommandSDK) Register(name string, f func(string, CommandRequest) interface{}, opts ...*options.CommandOptions) {
	mergeOpts := options.MergeCommandOptions(opts...)
	if system.functionDict == nil {
		system.functionDict = make(map[string]func(string, CommandRequest) interface{})
	}
	system.functionDict[name] = f
	logger.SysLog.Debugf("[Engine][OperatorRegister] `%s` Registered", name)
	if *mergeOpts.HttpSupport {
		system.RegisterHttp(name, CommandToHttpAdapter(f))
	}
}

func (system *CommandSDK) RegisterDirect(name string, f func(string, CommandRequest) interface{}, opts ...*options.CommandOptions) {
	mergeOpts := options.MergeCommandOptions(opts...)
	if system.drtFunctionDict == nil {
		system.drtFunctionDict = make(map[string]func(string, CommandRequest) interface{})
	}
	system.drtFunctionDict[name] = f
	logger.SysLog.Debugf("[Engine][OperatorRegisterDirect] `%s` Registered", name)
	if *mergeOpts.HttpSupport {
		system.RegisterHttpDirect(name, CommandToHttpAdapter(f))
	}
}

func (system *CommandSDK) RegisterHttpDirect(name string, f func(*gin.Context, CommandRequest)) {
	if system.drtHttpFunctionDict == nil {
		system.drtHttpFunctionDict = make(map[string]func(*gin.Context, CommandRequest))
	}
	system.drtHttpFunctionDict[name] = f
	logger.SysLog.Debugf("[Engine][OperatorHttpDirectRegister] `%s` Registered", name)
}

func (system *CommandSDK) RegisterHttp(name string, f func(*gin.Context, CommandRequest)) {
	if system.httpFunctionDict == nil {
		system.httpFunctionDict = make(map[string]func(*gin.Context, CommandRequest))
	}
	system.httpFunctionDict[name] = f
	logger.SysLog.Debugf("[Engine][OperatorHttpRegister] `%s` Registered", name)
}
