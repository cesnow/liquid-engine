package GameFoundation

import (
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Middlewares"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteRegister(c *gin.Context) {

	var command *LiquidSDK.CmdRegister
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	Logger.SysLog.Debugf("[CMD][Register] %+v", command)

	if command.FromType == "" {
		Logger.SysLog.Errorf("[CMD][Register] Create Member Failed, From Type is empty")
		result := gin.H{"registerStatus": 0, "error": "from_type is empty"}
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		c.Abort()
		return
	}

	if command.Account == "" || command.Password == ""{
		Logger.SysLog.Errorf("[CMD][Register] Create Member Failed, Account/Password is empty")
		result := gin.H{"registerStatus": 0, "error": "account/password is empty"}
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		c.Abort()
		return
	}

	member := LiquidSDK.GetServer().GetMemberSystem(command.FromType)
	errorMessage := ""
	resultStatus := 0
	if member == nil {
		errorMessage = "member system is not defined : " + command.FromType
	} else {
		resultStatus, errorMessage = member.Register(
			command.FromType,
			command.Account,
			command.Password,
			"",
			command.ExtraData,
		)
	}

	result := gin.H{"registerStatus": resultStatus, "error": errorMessage}
	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
}
