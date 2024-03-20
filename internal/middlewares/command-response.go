package middlewares

import (
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommandResponse(c *gin.Context, response *LiquidSDK.CmdCommandResponse) {
	if response.CmdData != nil {
		if _, ok := response.CmdData.(LiquidSDK.CmdErrorResponse); ok {
			c.String(http.StatusInternalServerError, GetLiquidResult(response))
			return
		}
		c.String(http.StatusOK, GetLiquidResult(response))
		return
	}
	c.String(http.StatusInternalServerError, GetLiquidResult(
		LiquidSDK.ResponseError("NO_RESPONSE_DATA"),
	))
}
