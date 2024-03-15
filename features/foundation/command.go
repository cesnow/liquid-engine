package foundation

import (
	"encoding/json"
	"fmt"
	"github.com/cesnow/liquid-engine/internal/middlewares"
	LiquidDB "github.com/cesnow/liquid-engine/liquid-db"
	"github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteCommand(c *gin.Context) {

	var command *LiquidSDK.CmdCommand
	unmarshalErr := json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	if unmarshalErr != nil {
		logger.SysLog.Warnf("[CMD][Command] Unmarshal Command Data Failed, %s", unmarshalErr)
	}

	c.Set("GIN_MSG", fmt.Sprintf("(sn: %s, id: %s, name: %s)", *command.CmdSn, *command.CmdId, *command.CmdName))

	result := &LiquidSDK.CmdCommandResponse{
		CmdData: nil,
		CmdSn:   nil,
	}

	if command.LiquidId == nil || command.LiquidToken == nil {
		logger.SysLog.Warnf("[CMD][Command] ID & Token is empty !")
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(gin.H{
			"code":  1401,
			"error": fmt.Sprintf("ID & Token is empty !"),
		}))
		c.Abort()
		return
	}

	if command.Platform == nil {
		platformMain := "main"
		command.Platform = &platformMain
	}

	tokenKey := fmt.Sprintf("token_%s_%s", *command.LiquidId, *command.Platform)
	authToken, authTokenErr := LiquidDB.GetCacheDb().Get(tokenKey)
	liquidToken := string(authToken)

	if authTokenErr != nil || liquidToken != *command.LiquidToken {
		logger.SysLog.Warnf("[CMD][Command] Data Verify Failed")
		c.String(http.StatusUnauthorized, middlewares.GetLiquidResult(gin.H{
			"code":  1402,
			"error": fmt.Sprintf("data verify failed !"),
		}))
		c.Abort()
		return
	}

	// TODO: Server Maintain States (Unsupported) 4003

	setUserTokenErr := LiquidDB.GetCacheDb().SetString(tokenKey, liquidToken, 1800)
	if setUserTokenErr != nil {
		logger.SysLog.Warnf("[CMD][Command] Refresh User Token Failed, %s", setUserTokenErr)
	}

	feature := LiquidSDK.GetServer().GetFeature(*command.CmdId)
	if feature == nil {
		c.String(http.StatusForbidden, middlewares.GetLiquidResult(gin.H{
			"code":  1404,
			"error": fmt.Sprintf("feature(cmd_id) not found !"),
		}))
		c.Abort()
		return
	}
	runCommandData := feature.RunCommand(command)
	result.CmdData = runCommandData

	result.CmdSn = command.CmdSn
	c.String(http.StatusOK, middlewares.GetLiquidResult(result))
}
