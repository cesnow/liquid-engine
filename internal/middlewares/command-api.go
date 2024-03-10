package middlewares

import (
	"encoding/base64"
	"encoding/json"
	LiquidModels "github.com/cesnow/liquid-engine/liquid-models"
	"github.com/gin-gonic/gin"
	"github.com/xxtea/xxtea-go/xxtea"
	"net/http"
	"time"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("liquid-token")
		if header == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "header liquid-token is empty"})
			c.Abort()
			return
		}
		decoded, err := base64.URLEncoding.DecodeString(header)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "header liquid-token is invalid"})
			c.Abort()
			return
		}
		decrypted := xxtea.Decrypt(decoded, []byte(apiUserEncryptedXxTeaKey))
		var claims *LoginClaims
		err = json.Unmarshal(decrypted, &claims)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "header liquid-token is invalid"})
			c.Abort()
			return
		}
		if claims.Audience != apiUserDefaultAudience {
			c.JSON(http.StatusBadRequest, gin.H{"error": "header liquid-token is invalid, aud failed."})
			c.Abort()
			return
		}

		if claims.ExpiresAt < time.Now().UnixNano()/int64(time.Millisecond) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired."})
			c.Abort()
			return
		}

		c.Set("LiquidUser", claims)
		c.Next()
	}
}

var apiUserEncryptedXxTeaKey = "-LiquidEngine|Api|User-"
var apiUserDefaultAudience = "LiquidEngine"

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
		Audience:   apiUserDefaultAudience,
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
