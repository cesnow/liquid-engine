package foundation

import (
	"encoding/json"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"

	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteApiDirect(c *gin.Context) {

	featureId := c.Param("FeatureId")
	cmdName := c.Param("CmdName")
	rawBody, _ := c.GetRawData()

	var cmdData interface{}
	if len(rawBody) > 0 {
		err := json.Unmarshal(rawBody, &cmdData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 1510,
				"error":  "Invalid Request",
			})
			c.Abort()
			return
		}
	}

	command := &LiquidSDK.CmdCommand{
		LiquidId:    nil,
		LiquidToken: nil,
		Platform:    nil,
		CmdId:       &featureId,
		CmdSn:       nil,
		CmdName:     &cmdName,
		CmdData:     cmdData,
	}

	feature := LiquidSDK.GetServer().GetFeature(*command.CmdId)
	if feature == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1511,
			"error": "Feature Not Found",
		})
		c.Abort()
		return
	}
	feature.RunHttpDirectCommand(c, command)
	return
}
