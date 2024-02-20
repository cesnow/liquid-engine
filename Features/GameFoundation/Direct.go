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

	gameFeature := LiquidSDK.GetServer().GetGameFeature(*command.CmdId)
	if gameFeature == nil {
		c.String(http.StatusForbidden, Middlewares.GetLiquidResult(gin.H{
			"status": 5001,
			"error":  fmt.Sprintf("feature(cmd_id) not found !"),
		}))
		c.Abort()
		return
	}
	runCommandData := gameFeature.RunDirectCommand(command)
	result.CmdData = runCommandData

	result.CmdSn = command.CmdSn
	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
}
