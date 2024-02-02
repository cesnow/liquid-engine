package GameFoundation

import (
	"encoding/json"
	"fmt"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FakeCmdData struct {
	R string `json:"r,omitempty"`
}

func RouteApis(c *gin.Context) {

	cmdId := c.Param("CmdId")
	cmdName := c.Param("CmdName")
	rawBody, _ := c.GetRawData()

	fakeCommand := &LiquidSDK.CmdCommand{
		LiquidId:    nil,
		LiquidToken: nil,
		Platform:    nil,
		CmdId:       &cmdId,
		CmdSn:       nil,
		CmdName:     &cmdName,
		CmdData:     &FakeCmdData{R: string(rawBody)},
	}

	gameFeature := LiquidSDK.GetServer().GetGameFeature(*fakeCommand.CmdId)
	if gameFeature == nil {
		c.String(http.StatusOK, "[GameFeature] Not Found")
		c.Abort()
		return
	}
	featureResp := gameFeature.RunHttpCommand(fakeCommand)
	marshalResp, marshalErr := json.Marshal(featureResp)
	if marshalErr != nil {
		c.String(http.StatusOK, fmt.Sprintf("%+v", featureResp))
	} else {
		c.String(http.StatusOK, string(marshalResp))
	}
	return

}
