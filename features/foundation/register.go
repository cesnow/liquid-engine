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

func RouteRegister(c *gin.Context) {

	var command *LiquidSDK.CmdRegister
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	logger.SysLog.Debugf("[CMD][Register] %+v", command)

	if command.FromType == "" {
		logger.SysLog.Errorf("[CMD][Register] Create Member Failed, From Type is empty")
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("INVALID_REQUEST_FROM_TYPE"),
		))
		return
	}

	if command.Account == "" || command.Password == "" {
		logger.SysLog.Errorf("[CMD][Register] Create Member Failed, Account/Password is empty")
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("INVALID_REQUEST_ACCOUNT_OR_PASSWORD"),
		))
		return
	}

	member := LiquidSDK.GetServer().GetMemberSystem(command.FromType)
	if member == nil {
		c.String(http.StatusForbidden, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError(fmt.Sprintf("MEMBER_SYSTEM_NOT_DEFINED:%s", command.FromType)),
		))
		return
	}

	resultStatus, errorMessage := member.Register(
		command.FromType,
		command.Account,
		command.Password,
		"",
		command.ExtraData,
	)

	if errorMessage != "" {
		c.String(http.StatusInternalServerError, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError(errorMessage),
		))
		return
	}

	c.String(http.StatusOK, middlewares.GetLiquidResult(gin.H{"registerStatus": resultStatus}))
}
