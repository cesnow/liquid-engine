package foundation

import (
	"encoding/json"
	"github.com/cesnow/liquid-engine/internal/middlewares"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteApiCommand(c *gin.Context) {

	cmdLiquidUser, _ := c.Get("LiquidUser")
	loginClaims := cmdLiquidUser.(*middlewares.LoginClaims)

	featureId := c.Param("FeatureId")
	cmdName := c.Param("CmdName")
	rawBody, _ := c.GetRawData()

	var cmdData interface{}
	if len(rawBody) > 0 {
		err := json.Unmarshal(rawBody, &cmdData)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				LiquidSDK.ResponseError("INVALID_REQUEST"),
			)
			return
		}
	}

	command := &LiquidSDK.CmdCommand{
		LiquidId: &loginClaims.AutoId,
		Platform: nil,
		CmdId:    &featureId,
		CmdSn:    nil,
		CmdName:  &cmdName,
		CmdData:  cmdData,
	}

	feature := LiquidSDK.GetServer().GetFeature(*command.CmdId)
	if feature == nil {
		c.JSON(
			http.StatusNotFound,
			LiquidSDK.ResponseError("FEATURE_NOT_FOUND"),
		)
		return
	}
	feature.RunHttpCommand(c, command)
	return
}
