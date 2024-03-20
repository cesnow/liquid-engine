package foundation

import (
	"encoding/json"
	"fmt"
	"github.com/cesnow/liquid-engine/internal/middlewares"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/cesnow/liquid-engine/logger"
	UtilLiquid "github.com/cesnow/liquid-engine/utils/liquid"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteDirect(c *gin.Context) {

	var command *LiquidSDK.CmdCommand
	unmarshalErr := json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	if unmarshalErr != nil {
		logger.SysLog.Warnf("[CMD][Command] Unmarshal Command Data Failed, %s", unmarshalErr)
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("INVALID_REQUEST")),
		)
		return
	}

	c.Set("GIN_MSG", fmt.Sprintf("(sn: %s, id: %s, name: %s)", *command.CmdSn, *command.CmdId, *command.CmdName))

	feature := LiquidSDK.GetServer().GetFeature(*command.CmdId)
	if feature == nil {
		c.String(http.StatusNotFound, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("FEATURE_NOT_FOUND")),
		)
		return
	}

	runCommandData := feature.RunDirectCommand(command)

	result := &LiquidSDK.CmdCommandResponse{
		CmdData: runCommandData,
		CmdSn:   command.CmdSn,
	}
	UtilLiquid.CommandResponse(c, result)
	return
}
