package Foundation

import (
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootKey(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, LiquidSDK.GetServer().GetKey())
}
