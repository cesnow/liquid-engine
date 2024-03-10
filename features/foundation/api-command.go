package foundation

import (
	"github.com/cesnow/liquid-engine/internal/middlewares"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/cesnow/liquid-engine/logger"

	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteApiCommand(c *gin.Context) {

	cmdLiquidUser, _ := c.Get("LiquidUser")
	loginClaims := cmdLiquidUser.(*middlewares.LoginClaims)

	logger.SysLog.Infof("[CMD][xxxxx] %+v", loginClaims)

	cmdId := c.Param("CmdId")
	cmdName := c.Param("CmdName")
	rawBody, _ := c.GetRawData()

	command := &LiquidSDK.CmdCommand{
		LiquidId: &loginClaims.AutoId,
		Platform: nil,
		CmdId:    &cmdId,
		CmdSn:    nil,
		CmdName:  &cmdName,
		CmdData:  string(rawBody),
	}

	feature := LiquidSDK.GetServer().GetFeature(*command.CmdId)
	if feature == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Feature Not Found"})
		c.Abort()
		return
	}
	feature.RunHttpCommand(c, command)
	return
}