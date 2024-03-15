package LiquidSDK

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommandToHttpAdapter(f func(string, CommandRequest) interface{}) func(*gin.Context, CommandRequest) {
	return func(c *gin.Context, cmd CommandRequest) {
		exec := f(cmd.GetLiquidId(), cmd)
		if exec != nil {
			if execMap, ok := exec.(*CmdErrorResponse); ok {
				c.JSON(http.StatusInternalServerError, execMap)
				return
			}
			c.JSON(http.StatusOK, exec)
		} else {
			c.Status(http.StatusInternalServerError)
		}
	}
}
