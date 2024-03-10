package foundation

import (
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootKey(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, LiquidSDK.GetServer().GetKey())
}

func Root(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, "")
}
