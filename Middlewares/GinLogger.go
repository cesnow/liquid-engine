package Middlewares

import (
	"fmt"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

func setupLogging(duration time.Duration) {
	go func() {
		for range time.Tick(duration) {
			Logger.SysLog.Sync()
		}
	}()
}

func ErrorLogger() gin.HandlerFunc {
	return ErrorLoggerT(gin.ErrorTypeAny)
}

func ErrorLoggerT(t gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !c.Writer.Written() {
			json := c.Errors.ByType(t).JSON()
			if json != nil {
				c.JSON(-1, json)
			}
		}
	}
}

func GinLogger(duration time.Duration) gin.HandlerFunc {
	setupLogging(duration)

	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)

		// clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)
		methodColor := colorForMethod(method)
		path := c.Request.URL.Path
		latencyDiff := float64(latency.Microseconds())
		latencyUnit := "µs"

		if latencyDiff > 1000 {
			latencyDiff = latencyDiff / float64(1000)
			latencyUnit = "ms"
		}

		msg := ""
		ginMsg, msgExists := c.Get("GIN_MSG")
		if msgExists {
			msg = ginMsg.(string)
		}

		message := fmt.Sprintf(
			"%s[%d]%s%s[%s]%s %s %s (%.3f %s)",
			statusColor,
			statusCode,
			reset,
			methodColor,
			method,
			reset,
			path,
			msg,
			latencyDiff,
			latencyUnit,
		)

		switch {
		case statusCode >= 400 && statusCode <= 499:
			Logger.SysLog.Warn("[GIN]", message)
		case statusCode >= 500:
			Logger.SysLog.Error("[GIN]", message)
		default:
			Logger.SysLog.Info("[GIN]", message)
		}
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code <= 299:
		return green
	case code >= 300 && code <= 399:
		return white
	case code >= 400 && code <= 499:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch {
	case method == "GET":
		return blue
	case method == "POST":
		return cyan
	case method == "PUT":
		return yellow
	case method == "DELETE":
		return red
	case method == "PATCH":
		return green
	case method == "HEAD":
		return magenta
	case method == "OPTIONS":
		return white
	default:
		return reset
	}
}
