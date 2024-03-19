package foundation

import (
	"encoding/json"
	"github.com/cesnow/liquid-engine/internal/middlewares"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteApiCommand(c *gin.Context) {

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

	feature := LiquidSDK.GetServer().GetFeature(featureId)
	if feature == nil {
		c.JSON(
			http.StatusNotFound,
			LiquidSDK.ResponseError("FEATURE_NOT_FOUND"),
		)
		return
	}

	command := &LiquidSDK.CmdCommand{
		LiquidId: nil,
		Platform: nil,
		CmdId:    &featureId,
		CmdSn:    nil,
		CmdName:  &cmdName,
		CmdData:  cmdData,
	}

	commandExists := feature.IsHttpExists(cmdName)
	if commandExists {
		loginClaims, err := middlewares.GetClaim(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, LiquidSDK.ResponseError(err.Error()))
			return
		}
		command.LiquidId = &loginClaims.AutoId
		feature.RunHttpCommand(c, command)
		return
	}

	directCommandExists := feature.IsHttpDirectExists(cmdName)
	if directCommandExists {
		feature.RunHttpDirectCommand(c, command)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "FEATURE_COMMAND_NOT_FOUND"})
	return

}
