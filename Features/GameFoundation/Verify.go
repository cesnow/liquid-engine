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

func RouteVerify(c *gin.Context) {

	var command *LiquidSDK.CmdAccount
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	Logger.SysLog.Debugf("[CMD][Login] %+v", command)

	result := &LiquidSDK.CmdAccountResponse{
		AutoId:     nil,
		InviteCode: nil,
	}

	if command.FromId == "" {
		Logger.SysLog.Warnf("[CMD][Verify] FromId is Empty, %+v", command)
		c.Status(http.StatusBadRequest)
	}

	var liquidUser *Models.LiquidUser

	if command.FromType == "guest" {
		autoId := command.FromId
		liquidUser = Models.FindLiquidGuestUser(autoId) // find auto_id
	} else {
		// TODO: Customize Validate User Data (Unsupported)
		liquidUser = Models.FindLiquidUserFromType(command.FromType, command.FromId) // find from_type, from_id
	}

	if liquidUser != nil {
		result.AutoId = &liquidUser.AutoId
		result.InviteCode = &liquidUser.InviteCode
	}

	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))

}
