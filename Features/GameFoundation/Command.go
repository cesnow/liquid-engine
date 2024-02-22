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

	c.Set("GIN_MSG", fmt.Sprintf("(sn: %s, id: %s, name: %s)", *command.CmdSn, *command.CmdId, *command.CmdName))

	result := &LiquidSDK.CmdCommandResponse{
		CmdData: nil,
		CmdSn:   nil,
	}

	if command.LiquidId == nil || command.LiquidToken == nil {
		Logger.SysLog.Warnf("[CMD][Command] ID & Token is empty !")
		c.String(http.StatusBadRequest, Middlewares.GetLiquidResult(gin.H{
			"status": 4001,
			"error":  fmt.Sprintf("ID & Token is empty !"),
		}))
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
		c.String(http.StatusUnauthorized, Middlewares.GetLiquidResult(gin.H{
			"status": 4002,
			"error":  fmt.Sprintf("data verify failed !"),
		}))
		c.Abort()
		return
	}

	// TODO: Server Maintain States (Unsupported) 4003

	setUserTokenErr := LiquidSDK.GetServer().GetCacheDb().SetString(tokenKey, liquidToken, 1800)
	if setUserTokenErr != nil {
		Logger.SysLog.Warnf("[CMD][Command] Refresh User Token Failed, %s", setUserTokenErr)
	}

	gameFeature := LiquidSDK.GetServer().GetGameFeature(*command.CmdId)
	if gameFeature == nil {
		c.String(http.StatusForbidden, Middlewares.GetLiquidResult(gin.H{
			"status": 4004,
			"error":  fmt.Sprintf("feature(cmd_id) not found !"),
		}))
		c.Abort()
		return
	}
	runCommandData := gameFeature.RunCommand(command)
	result.CmdData = runCommandData

	result.CmdSn = command.CmdSn
	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
}
