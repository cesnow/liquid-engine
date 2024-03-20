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

	if command.LiquidId == nil || command.LiquidToken == nil {
		logger.SysLog.Warnf("[CMD][Command] ID & Token is empty !")
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("INVALID_REQUEST_LIQUID_ID_OR_LIQUID_TOKEN"),
		))
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
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("COMMAND_DATA_VERIFY_FAILED"),
		))
		return
	}

	// TODO: Server Maintain States (Unsupported) 4003

	setUserTokenErr := LiquidDB.GetCacheDb().SetString(tokenKey, liquidToken, 1800)
	if setUserTokenErr != nil {
		logger.SysLog.Warnf("[CMD][Command] Refresh User Token Failed, %s", setUserTokenErr)
	}

	feature := LiquidSDK.GetServer().GetFeature(*command.CmdId)
	if feature == nil {
		c.String(http.StatusNotFound, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("FEATURE_NOT_FOUND"),
		))
		return
	}
	runCommandData := feature.RunCommand(command)

	result := &LiquidSDK.CmdCommandResponse{
		CmdData: runCommandData,
		CmdSn:   command.CmdSn,
	}
	middlewares.CommandResponse(c, result)
	return
}
