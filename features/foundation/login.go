package foundation

import (
	"encoding/json"
	"github.com/cesnow/liquid-engine/internal/middlewares"
	LiquidModels "github.com/cesnow/liquid-engine/liquid-models"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteLogin(c *gin.Context) {

	var command *LiquidSDK.CmdAccount
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	logger.SysLog.Debugf("[CMD][Login] %+v", command)

	var liquidUser *LiquidModels.LiquidUser

	if command.FromType == "guest" {
		autoId := command.FromId
		inviteCode := command.FromToken
		liquidUser = LiquidModels.FindLiquidGuestUser(autoId) // find auto_id
		if liquidUser == nil {
			liquidUser = LiquidModels.CreateLiquidUser(command.FromType, "")
		} else {
			if liquidUser.FromType != "guest" || inviteCode != liquidUser.InviteCode {
				c.String(http.StatusUnauthorized, middlewares.GetLiquidResult(
					LiquidSDK.ResponseError("GUEST_FROM_TOKEN_INVALID"),
				))
				return
			}
		}
	} else {

		if command.FromId == "" || command.FromToken == "" {
			c.String(http.StatusBadRequest, middlewares.GetLiquidResult(
				LiquidSDK.ResponseError("INVALID_REQUEST_FROM_ID_OR_FROM_TOKEN")),
			)
			return
		}

		// Run Validation, Third Party Member System
		resultValidate := false
		errorMessage := ""

		member := LiquidSDK.GetServer().GetMemberSystem(command.FromType)
		if member == nil {
			errorMessage = "MEMBER_SYSTEM_NOT_DEFINED:" + command.FromToken
		} else {
			overrideFromId := ""
			resultValidate, errorMessage, overrideFromId = member.Validate(
				command.FromId,
				command.FromToken,
				command.Platform,
				command.ExtraData,
			)
			if overrideFromId != "" {
				command.FromId = overrideFromId
			}
		}

		if !resultValidate {
			c.String(http.StatusUnauthorized, middlewares.GetLiquidResult(
				LiquidSDK.ResponseError(errorMessage),
			))
			return
		}

		liquidUser = LiquidModels.FindLiquidUserFromType(command.FromType, command.FromId)
		if liquidUser == nil {
			liquidUser = LiquidModels.CreateLiquidUser(command.FromType, command.FromId)
		}

	}

	if liquidUser.IsDeactivate == true {
		c.String(http.StatusForbidden, middlewares.GetLiquidResult(
			LiquidSDK.ResponseError("USER_DEACTIVATED"),
		))
		return
	}

	// TODO: BlockSystem (Unsupported)

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
