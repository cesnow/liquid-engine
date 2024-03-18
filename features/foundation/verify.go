package foundation

import (
	"encoding/json"
	LiquidModels "github.com/cesnow/liquid-engine/liquid-models"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"

	"github.com/cesnow/liquid-engine/internal/middlewares"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteVerify(c *gin.Context) {

	var command *LiquidSDK.CmdAccount
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	logger.SysLog.Debugf("[CMD][Login] %+v", command)

	if command.FromId == "" {
		logger.SysLog.Warnf("[CMD][Verify] FromId is Empty, %+v", command)
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("INVALID_REQUEST_FROM_ID")),
		)
	}

	var liquidUser *LiquidModels.LiquidUser

	if command.FromType == "guest" {
		autoId := command.FromId
		liquidUser = LiquidModels.FindLiquidGuestUser(autoId) // find auto_id
	} else {
		// TODO: Customize Validate User Data (Unsupported)
		liquidUser = LiquidModels.FindLiquidUserFromType(command.FromType, command.FromId) // find from_type, from_id
	}

	if liquidUser != nil {
		result := &LiquidSDK.CmdAccountResponse{
			AutoId:     &liquidUser.AutoId,
			InviteCode: &liquidUser.InviteCode,
		}
		c.String(http.StatusOK, middlewares.GetLiquidResult(result))
		return
	}

	c.String(http.StatusInternalServerError, middlewares.GetLiquidResult(
		LiquidSDK.ResponseError("USER_NOT_FOUND"),
	))
}
