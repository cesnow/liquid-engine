package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	LiquidModels "github.com/cesnow/liquid-engine/liquid-models"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/gin-gonic/gin"
	"github.com/xxtea/xxtea-go/xxtea"
	"net/http"
	"time"
)

func GetClaim(c *gin.Context) (*LoginClaims, error) {
	header := c.GetHeader("liquid-token")
	if header == "" {
		return nil, errors.New("LIQUID_TOKEN_HEADER_REQUIRED")
	}
	decoded, err := base64.URLEncoding.DecodeString(header)
	if err != nil {
		return nil, errors.New("LIQUID_TOKEN_HEADER_INVALID")
	}
	decrypted := xxtea.Decrypt(decoded, []byte(apiUserEncryptedXxTeaKey))
	var claims *LoginClaims
	err = json.Unmarshal(decrypted, &claims)
	if err != nil {
		return nil, errors.New("LIQUID_TOKEN_HEADER_INVALID")
	}
	if claims.Audience != LiquidSDK.GetServer().CodeName {
		return nil, errors.New("LIQUID_TOKEN_AUDIENCE_INVALID")
	}
	if claims.ExpiresAt < time.Now().UnixNano()/int64(time.Millisecond) {
		return nil, errors.New("LIQUID_TOKEN_EXPIRED")
	}
	return claims, nil
}

var apiUserEncryptedXxTeaKey = "-LiquidEngine|Api|User-"

type LoginClaims struct {
	Audience   string `json:"aud,omitempty"`
	ExpiresAt  int64  `json:"exp,omitempty"`
	IssuedAt   int64  `json:"iat,omitempty"`
	AutoId     string `json:"aid,omitempty"`
	InviteCode string `json:"ivc,omitempty"`
	FromType   string `json:"ftp,omitempty"`
	FromId     string `json:"fid,omitempty"`
}

func GenerateToken(user *LiquidModels.LiquidUser) string {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	twoWeeks := 14 * 24 * time.Hour
	expired := now + int64(twoWeeks/time.Millisecond)
	claims := LoginClaims{
		Audience:   LiquidSDK.GetServer().CodeName,
		ExpiresAt:  expired,
		IssuedAt:   now,
		AutoId:     user.AutoId,
		InviteCode: user.InviteCode,
		FromType:   user.FromType,
		FromId:     user.FromId,
	}
	data, _ := json.Marshal(claims)
	encryptedData := xxtea.Encrypt(data, []byte(apiUserEncryptedXxTeaKey))
	ResultData := base64.URLEncoding.EncodeToString(encryptedData)
	return ResultData
}

func CommandRouteResponse(c *gin.Context, response any) {
	if response != nil {
		if _, ok := response.(LiquidSDK.CmdErrorResponse); ok {
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusInternalServerError, LiquidSDK.ResponseError("NO_RESPONSE_DATA"))
}
