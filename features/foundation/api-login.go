package foundation

import (
	"encoding/json"
	"github.com/cesnow/liquid-engine/internal/middlewares"
	LiquidModels "github.com/cesnow/liquid-engine/liquid-models"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteApiLogin(c *gin.Context) {

	var command *LiquidSDK.CmdAccount
	rawBody, _ := c.GetRawData()
	err := json.Unmarshal(rawBody, &command)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			LiquidSDK.ResponseError("INVALID_REQUEST"),
		)
		return
	}

	var liquidUser *LiquidModels.LiquidUser

	if command.FromType == "guest" {
		autoId := command.FromId
		inviteCode := command.FromToken
		liquidUser = LiquidModels.FindLiquidGuestUser(autoId) // find auto_id
		if liquidUser == nil {
			liquidUser = LiquidModels.CreateLiquidUser(command.FromType, "")
		} else {
			if liquidUser.FromType != "guest" || inviteCode != liquidUser.InviteCode {
				token := middlewares.GenerateToken(liquidUser)
				c.SetCookie("liquid_token", token, 24*3600, "/", "", false, true)
				c.JSON(http.StatusOK, gin.H{"token": token})
				return
			}
		}
	} else {

		if command.FromId == "" || command.FromToken == "" {
			c.JSON(
				http.StatusBadRequest,
				LiquidSDK.ResponseError("FROM_ID_OR_FROM_TOKEN_IS_REQUIRED"),
			)
			return
		}

		// Run Validation, Third Party Member System
		resultValidate := false
		errorMessage := ""

		member := LiquidSDK.GetServer().GetMemberSystem(command.FromType)
		if member == nil {
			errorMessage = "member system is not defined : " + command.FromToken
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
			c.JSON(http.StatusUnauthorized, LiquidSDK.ResponseError(errorMessage))
			return
		}

		liquidUser = LiquidModels.FindLiquidUserFromType(command.FromType, command.FromId)
		if liquidUser == nil {
			liquidUser = LiquidModels.CreateLiquidUser(command.FromType, command.FromId)
		}

	}

	if liquidUser.IsDeactivate == true {
		c.JSON(http.StatusForbidden, LiquidSDK.ResponseError("USER_DEACTIVATED"))
		return
	}

	// TODO: BlockSystem (Unsupported)

	if liquidUser != nil {
		token := middlewares.GenerateToken(liquidUser)
		c.SetCookie("liquid_token", token, 24*3600, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusInternalServerError, LiquidSDK.ResponseError("INTERNAL_SERVER_ERROR"))

}
