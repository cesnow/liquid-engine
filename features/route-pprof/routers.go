package route_pprof

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/pprof"
)

func Routers(gin *gin.Engine) {
	DebugRouters := gin.Group("/debug")
	DebugRouters.GET("/pprof", pprofHandler(pprof.Index))
	DebugRouters.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
	DebugRouters.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
	DebugRouters.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
	DebugRouters.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
	DebugRouters.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
	DebugRouters.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	DebugRouters.GET("/cmdline", pprofHandler(pprof.Cmdline))
	DebugRouters.GET("/profile", pprofHandler(pprof.Profile))
	DebugRouters.GET("/symbol", pprofHandler(pprof.Symbol))
	DebugRouters.GET("/trace", pprofHandler(pprof.Trace))
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
