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

func RouteBind(c *gin.Context) {

	var command *LiquidSDK.CmdBind
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	Logger.SysLog.Debugf("[CMD][Bind] %+v", command)

	result := &LiquidSDK.CmdAccountResponse{}

	liquidUser := Models.FindLiquidUserByAutoId(command.AutoId, command.InviteCode)

	if liquidUser == nil {
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		return
	}

	_, bindErr := Models.BindLiquidUser(command.AutoId, command.FromId, command.FromType, command.FromToken)

	if bindErr != nil {
		Logger.SysLog.Warnf("[CMD][Bind] Duplicate Bind Failed, %s", bindErr)
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		return
	}

	result.AutoId = &liquidUser.AutoId
	result.InviteCode = &liquidUser.InviteCode
	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))

}
