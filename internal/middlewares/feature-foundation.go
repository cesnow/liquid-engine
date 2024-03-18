package middlewares

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/cesnow/liquid-engine/liquid-sdk"

	"github.com/cesnow/liquid-engine/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLiquidData() gin.HandlerFunc {
	return func(c *gin.Context) {

		CodenameLiquidKey := LiquidSDK.GetServer().GetKeyStatic()

		RawDataBody, GetRawBodyErr := c.GetRawData()
		if GetRawBodyErr != nil {
			logger.SysLog.Errorf("[Engine][Middleware] Get Liquid Data Failed, %s", GetRawBodyErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST"})
			c.Abort()
			return
		}

		DataBody, decodeBodyErr := base64.StdEncoding.DecodeString(string(RawDataBody))
		if decodeBodyErr != nil {
			logger.SysLog.Errorf("[Engine][Middleware] Decode Liquid Data Failed, %s", decodeBodyErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST"})
			c.Abort()
			return
		}

		var StructureLiquidData *LiquidSDK.CmdSignedBody
		DataUnmarshalError := json.Unmarshal(DataBody, &StructureLiquidData)
		if DataUnmarshalError != nil {
			logger.SysLog.Errorf("[Engine][Middleware] Unmarshal Liquid Data Failed, %s", DataUnmarshalError)
			c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST"})
			c.Abort()
			return
		}

		DataVerify := hmac.New(sha1.New, []byte(CodenameLiquidKey))
		DataVerify.Write([]byte(StructureLiquidData.LiData))
		DataVerifyHexDigest := hex.EncodeToString(DataVerify.Sum(nil))

		if StructureLiquidData.LiSign != DataVerifyHexDigest {
			logger.SysLog.Error("[Engine][Middleware] Verify Liquid Data Failed")

			c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST"})
			c.Abort()
			return
		}

		DecodedCommandData, DecodedCommandDataError := base64.StdEncoding.DecodeString(StructureLiquidData.LiData)
		if DecodedCommandDataError != nil {
			logger.SysLog.Errorf("[Engine][Middleware] Decode Command Liquid Data Failed, %s", DecodedCommandDataError)
			c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST"})
			c.Abort()
			return
		}

		//var CommandData map[string]interface{}
		//UnmarshalCommandDataErr := json.Unmarshal(DecodedCommandData, &CommandData)
		//if UnmarshalCommandDataErr != nil {
		//	Logger.SysLog.Errorf("[Engine][Middleware] Unmarshal Command Liquid Data Failed, %s", UnmarshalCommandDataErr)
		//	c.Status(http.StatusBadRequest)
		//	c.Abort()
		//}
		c.Set("CommandData", DecodedCommandData)
		c.Next()
	}
}

func GetLiquidResult(liquidData interface{}) string {
	MarshalData, MarshalDataErr := json.Marshal(liquidData)
	if MarshalDataErr != nil {
		logger.SysLog.Errorf("[CMD][ResultData] Can't Marshal Data, %s", MarshalDataErr)
		return ""
	}
	ResultData := base64.StdEncoding.EncodeToString(MarshalData)
	return ResultData
}
