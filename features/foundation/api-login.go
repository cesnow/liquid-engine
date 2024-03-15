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
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1211,
			"error": "Invalid Request",
		})
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
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1212,
				"error": "from_id or from_token is empty",
			})
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
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  1213,
				"error": errorMessage,
			})
			return
		}

		liquidUser = LiquidModels.FindLiquidUserFromType(command.FromType, command.FromId)
		if liquidUser == nil {
			liquidUser = LiquidModels.CreateLiquidUser(command.FromType, command.FromId)
		}

	}

	if liquidUser.IsDeactivate == true {
		c.JSON(http.StatusForbidden, gin.H{
			"code":  1214,
			"error": "user is deactivated",
		})
		return
	}

	// TODO: BlockSystem (Unsupported)

	if liquidUser != nil {
		token := middlewares.GenerateToken(liquidUser)
		c.SetCookie("liquid_token", token, 24*3600, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"code":  1215,
		"error": "something went wrong !",
	})

}
