package foundation

import (
	"encoding/json"
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
		result := gin.H{
			"registerStatus": 0,
			"code":           1101,
			"error":          "from_type is empty",
		}
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(result))
		c.Abort()
		return
	}

	if command.Account == "" || command.Password == "" {
		logger.SysLog.Errorf("[CMD][Register] Create Member Failed, Account/Password is empty")
		result := gin.H{
			"registerStatus": 0,
			"code":           1102,
			"error":          "account/password is empty",
		}
		c.String(http.StatusBadRequest, middlewares.GetLiquidResult(result))
		c.Abort()
		return
	}

	member := LiquidSDK.GetServer().GetMemberSystem(command.FromType)
	if member == nil {
		c.String(http.StatusForbidden, middlewares.GetLiquidResult(gin.H{
			"registerStatus": 0,
			"code":           1103,
			"error":          "member system is not defined : " + command.FromType,
		}))
		c.Abort()
	}

	resultStatus, errorMessage := member.Register(
		command.FromType,
		command.Account,
		command.Password,
		"",
		command.ExtraData,
	)
	result := gin.H{"registerStatus": resultStatus, "error": errorMessage}
	c.String(http.StatusOK, middlewares.GetLiquidResult(result))
}
