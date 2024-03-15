package foundation

import (
	"encoding/json"
	"fmt"
	LiquidDB "github.com/cesnow/liquid-engine/liquid-db"
	LiquidModels "github.com/cesnow/liquid-engine/liquid-models"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"

	"github.com/cesnow/liquid-engine/internal/middlewares"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/cesnow/liquid-engine/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func RouteAuth(c *gin.Context) {

	var command *LiquidSDK.CmdAuth
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	logger.SysLog.Debugf("[CMD][Auth] %+v", command)

	result := &LiquidSDK.CmdAuthResponse{}

	if command.AutoId == nil || command.InviteCode == nil {
		c.String(http.StatusUnauthorized, middlewares.GetLiquidResult(gin.H{
			"code":  1301,
			"error": "auto_id or invite_code is empty",
		}))
		return
	}

	if command.Platform == nil {
		platformMain := "main"
		command.Platform = &platformMain
	}

	liquidUser := LiquidModels.FindLiquidUserByAutoId(*command.AutoId, *command.InviteCode)
	if liquidUser == nil {
		c.String(http.StatusUnauthorized, middlewares.GetLiquidResult(gin.H{
			"code":  1302,
			"error": "user is not found",
		}))
		return
	}

	if liquidUser.IsDeactivate == true {
		c.String(http.StatusForbidden, middlewares.GetLiquidResult(gin.H{
			"code":  1303,
			"error": "user is deactivated",
		}))
		return
	}

	LiquidModels.CheckDefaultPlayerData(liquidUser.AutoId)

	liquidToken := generateNewToken()

	// TODO: BlockSystem (not unsupported)

	// TODO: maybe multi login
	tokenKey := fmt.Sprintf("token_%s_%s", liquidUser.AutoId, *command.Platform)
	setUserTokenErr := LiquidDB.GetCacheDb().SetString(
		tokenKey,
		liquidToken,
		LiquidSDK.GetServer().TokenExpireTime,
	)
	if setUserTokenErr != nil {
		logger.SysLog.Warnf("[CMD][Auth] Set User Token Failed, %s", setUserTokenErr)
		c.String(http.StatusInternalServerError, middlewares.GetLiquidResult(gin.H{
			"code":  1304,
			"error": "set user token failed",
		}))
		return
	}

	result.LiquidId = &liquidUser.AutoId
	result.LiquidToken = &liquidToken

	c.String(http.StatusOK, middlewares.GetLiquidResult(result))

}

func generateNewToken() string {
	authTime := strconv.Itoa(time.Now().Nanosecond())
	return utils.EncodeMD5(authTime)
}
