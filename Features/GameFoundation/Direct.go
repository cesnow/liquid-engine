package GameFoundation

import (
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Middlewares"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteDirect(c *gin.Context) {

	var command *LiquidSDK.CmdCommand
	unmarshalErr := json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	if unmarshalErr != nil {
		Logger.SysLog.Warnf("[CMD][Command] Unmarshal Command Data Failed, %s", unmarshalErr)
	}

	result := &LiquidSDK.CmdCommandResponse{
		CmdData: nil,
		CmdSn:   nil,
	}

	// gRpc Routing Mode Checking
	if LiquidSDK.GetServer().GetRpcTrafficEnabled(){
		if rpcResult, rpcErr := GRpcCommand(command, true); rpcErr != nil {
			Logger.SysLog.Warnf("[CMD][Command] RPC Command Transfer Failed, %s", rpcErr)
		} else {
			var CmdResult interface{}
			_ = json.Unmarshal(rpcResult, &CmdResult)
			result.CmdData = CmdResult
		}
	}else{
		gameFeature := LiquidSDK.GetServer().GetGameFeature(*command.CmdId)
		if gameFeature == nil {
			c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
			c.Abort()
			return
		}
		runCommandData := gameFeature.RunDirectCommand(command)
		result.CmdData = runCommandData
	}

	result.CmdSn = command.CmdSn
	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
}
