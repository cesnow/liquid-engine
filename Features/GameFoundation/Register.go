package GameFoundation

import (
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Middlewares"
	"github.com/cesnow/LiquidEngine/Models"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func RouteRegister(c *gin.Context) {

	var command *LiquidSDK.CmdRegister
	_ = json.Unmarshal(c.MustGet("CommandData").([]byte), &command)
	Logger.SysLog.Debugf("[CMD][Register] %+v", command)

	dateNow := time.Now()

	liquidMember := bson.M{
		"account":    command.Account,
		"password":   command.Password,
		"createTime": dateNow,
	}

	registerStatus := 1
	_, err := Models.CreateLiquidMember(liquidMember)
	if err != nil {
		Logger.SysLog.Errorf("[CMD][Register] Create Member Failed, %s", err)
		registerStatus = 0
	}

	result := gin.H{"registerStatus": registerStatus}
	c.String(http.StatusOK, Middlewares.GetLiquidResult(result))
}
