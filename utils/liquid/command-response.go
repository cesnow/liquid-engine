package UtilLiquid

import (
	"github.com/cesnow/liquid-engine/internal/middlewares"
	"github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommandResponse(c *gin.Context, response *LiquidSDK.CmdCommandResponse) {
	if response.CmdData != nil {
		if _, ok := response.CmdData.(LiquidSDK.CmdErrorResponse); ok {
			c.String(http.StatusInternalServerError, middlewares.GetLiquidResult(response))
			return
		}
		c.String(http.StatusOK, middlewares.GetLiquidResult(response))
		return
	}
	c.String(http.StatusInternalServerError, middlewares.GetLiquidResult(
		LiquidSDK.ResponseError("NO_RESPONSE_DATA"),
	))
}

func CommandRouteResponse(c *gin.Context, response any) {
	if response != nil {
		if _, ok := response.(LiquidSDK.CmdErrorResponse); ok {
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusInternalServerError, LiquidSDK.ResponseError("NO_RESPONSE_DATA"))
}
