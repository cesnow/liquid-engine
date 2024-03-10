package foundation

import (
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"

	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteApiDirect(c *gin.Context) {

	cmdId := c.Param("CmdId")
	cmdName := c.Param("CmdName")
	rawBody, _ := c.GetRawData()

	command := &LiquidSDK.CmdCommand{
		LiquidId:    nil,
		LiquidToken: nil,
		Platform:    nil,
		CmdId:       &cmdId,
		CmdSn:       nil,
		CmdName:     &cmdName,
		CmdData:     string(rawBody),
	}

	feature := LiquidSDK.GetServer().GetFeature(*command.CmdId)
	if feature == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feature Not Found"})
		c.Abort()
		return
	}
	feature.RunHttpDirectCommand(c, command)
	return
}
