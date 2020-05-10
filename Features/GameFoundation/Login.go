package GameFoundation

import (
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Middlewares"
	"github.com/cesnow/LiquidEngine/Models"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteLogin(c *gin.Context) {

	var command *LiquidSDK.CmdAccount
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	Logger.SysLog.Debugf("[CMD][Login] %+v", command)

	result := &LiquidSDK.CmdAccountResponse{
		AutoId:     nil,
		InviteCode: nil,
	}

	var liquidUser *Models.LiquidUser

	if command.FromType == "guest" {
		autoId := command.FromId
		inviteCode := command.FromToken
		liquidUser = Models.FindLiquidGuestUser(autoId) // find auto_id
		if liquidUser == nil {
			liquidUser = Models.CreateLiquidUser(command.FromType, "")
		} else {
			if liquidUser.FromType != "guest" || inviteCode != liquidUser.InviteCode {
				c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
				return
			}
		}
	} else {

		if command.FromId == "" || command.FromToken == "" {
			c.String(http.StatusOK, Middlewares.GetLiquidResult(gin.H{"data": "from_id or from_token is empty"}))
			return
		}

		// Run Validation, Third Party Member System
		resultValidate := false
		errorMessage := ""
		if LiquidSDK.GetServer().GetRpcTrafficEnabled() {
			if rpcResult, rpcErr := GRpcLogin(command); rpcErr != nil {
				Logger.SysLog.Warnf("[RpcLogin] Transfer Failed, %s", rpcErr)
				errorMessage = rpcErr.Error()
			} else {
				resultValidate = rpcResult.Valid
				errorMessage = rpcResult.Msg
				if rpcResult.OverrideFromId != "" {
					command.FromId = rpcResult.OverrideFromId
				}
			}
		} else {
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
		}
		if !resultValidate {
			c.String(http.StatusOK, Middlewares.GetLiquidResult(gin.H{
				"data": errorMessage,
			}))
			return
		}

		liquidUser = Models.FindLiquidUserFromType(command.FromType, command.FromId)
		if liquidUser == nil {
			liquidUser = Models.CreateLiquidUser(command.FromType, command.FromId)
		}

	}

	// TODO: BlockSystem (Unsupported)

	if liquidUser != nil {
		result.AutoId = &liquidUser.AutoId
		result.InviteCode = &liquidUser.InviteCode
	}

	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
}
