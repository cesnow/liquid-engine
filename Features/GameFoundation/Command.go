package GameFoundation

import (
	"encoding/json"
	"fmt"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Middlewares"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteCommand(c *gin.Context) {

	var command *LiquidSDK.CmdCommand
	unmarshalErr := json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	if unmarshalErr != nil {
		Logger.SysLog.Warnf("[CMD][Command] Unmarshal Command Data Failed, %s", unmarshalErr)
	}

	result := &LiquidSDK.CmdCommandResponse{
		CmdData: nil,
		CmdSn:   nil,
	}

	if command.LiquidId == nil || command.LiquidToken == nil {
		Logger.SysLog.Warnf("[CMD][Command] ID & Token is empty !")
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		c.Abort()
		return
	}

	if command.Platform == nil {
		platformMain := "main"
		command.Platform = &platformMain
	}

	tokenKey := fmt.Sprintf("token_%s_%s", *command.LiquidId, *command.Platform)
	authToken, authTokenErr := LiquidSDK.GetServer().GetCacheDb().Get(tokenKey)
	liquidToken := string(authToken)

	if authTokenErr != nil || liquidToken != *command.LiquidToken {
		Logger.SysLog.Warnf("[CMD][Command] Data Verify Failed")
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		c.Abort()
		return
	}

	// TODO: Server Maintain States (Unsupported)

	setUserTokenErr := LiquidSDK.GetServer().GetCacheDb().SetString(
		tokenKey,
		liquidToken,
		LiquidSDK.GetServer().TokenExpireTime,
	)
	if setUserTokenErr != nil {
		Logger.SysLog.Warnf("[CMD][Command] Refresh User Token Failed, %s", setUserTokenErr)
	}

	// gRpc Routing Mode Checking
	if LiquidSDK.GetServer().GetRpcTrafficEnabled() {
		if rpcResult, rpcErr := GRpcCommand(command, false); rpcErr != nil {
			Logger.SysLog.Warnf("[CMD][Command] RPC Command Transfer Failed, %s", rpcErr)
		} else {
			var CmdResult interface{}
			_ = json.Unmarshal(rpcResult, &CmdResult)
			result.CmdData = CmdResult
		}
	} else {
		gameFeature := LiquidSDK.GetServer().GetGameFeature(*command.CmdId)
		if gameFeature == nil {
			c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
			c.Abort()
			return
		}
		runCommandData := gameFeature.RunCommand(command)
		result.CmdData = runCommandData
	}

	result.CmdSn = command.CmdSn
	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
}
