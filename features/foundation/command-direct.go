package foundation

import (
	"encoding/json"
	"fmt"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"

	"github.com/cesnow/liquid-engine/internal/middlewares"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteDirect(c *gin.Context) {

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

	feature := LiquidSDK.GetServer().GetFeature(*command.CmdId)
	if feature == nil {
		c.String(http.StatusForbidden, middlewares.GetLiquidResult(gin.H{
			"status": 5001,
			"error":  fmt.Sprintf("feature(cmd_id) not found !"),
		}))
		c.Abort()
		return
	}
	runCommandData := feature.RunDirectCommand(command)
	result.CmdData = runCommandData

	result.CmdSn = command.CmdSn
	c.String(http.StatusOK, middlewares.GetLiquidResult(result))
}
