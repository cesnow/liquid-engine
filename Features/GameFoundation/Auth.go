package GameFoundation

import (
	"encoding/json"
	"fmt"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Middlewares"
	"github.com/cesnow/LiquidEngine/Models"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/cesnow/LiquidEngine/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func RouteAuth(c *gin.Context) {

	var command *LiquidSDK.CmdAuth
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	Logger.SysLog.Debugf("[CMD][Auth] %+v", command)

	result := &LiquidSDK.CmdAuthResponse{}

	if command.AutoId == nil || command.InviteCode == nil {
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		return
	}

	if command.Platform == nil {
		platformMain := "main"
		command.Platform = &platformMain
	}

	liquidUser := Models.FindLiquidUserByAutoId(*command.AutoId, *command.InviteCode)
	if liquidUser == nil {
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		return
	}

	Models.CheckDefaultPlayerData(liquidUser.AutoId)

	liquidToken := generateNewToken()

	// TODO: BlockSystem (not unsupported)

	// TODO: maybe multi login
	tokenKey := fmt.Sprintf("token_%s_%s", liquidUser.AutoId, *command.Platform)
	setUserTokenErr := LiquidSDK.GetServer().GetCacheDb().SetString(tokenKey, liquidToken, 86400)
	if setUserTokenErr != nil {
		Logger.SysLog.Warnf("[CMD][Auth] Create User Token Failed, %s", setUserTokenErr)
		c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
		return
	}

	result.LiquidId = &liquidUser.AutoId
	result.LiquidToken = &liquidToken

	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))

}

func generateNewToken() string {
	authTime := strconv.Itoa(time.Now().Nanosecond())
	return Utils.EncodeMD5(authTime)
}
