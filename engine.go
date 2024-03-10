package LiquidEngine

import (
	"context"
	"fmt"
	"github.com/cesnow/liquid-engine/internal/feature-register"
	LiquidDb "github.com/cesnow/liquid-engine/liquid-db"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"

	"github.com/cesnow/liquid-engine/features/foundation"
	"github.com/cesnow/liquid-engine/features/health"
	_ "github.com/cesnow/liquid-engine/features/health"
	"github.com/cesnow/liquid-engine/internal/middlewares"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/cesnow/liquid-engine/options"
	"github.com/cesnow/liquid-engine/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "google.golang.org/grpc/encoding/gzip"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type IEngine interface {
}

type Engine struct {
	Config    *Config
	ginEngine *gin.Engine
	StartTime time.Time
}

var _ IEngine = &Engine{}

func New() *Engine {
	startTime := time.Now()
	logger.SysLog.Info("[Engine] Starting")
	engine := &Engine{
		Config: &Config{
			App:     &settings.AppConf{},
			Gin:     &settings.GinConf{},
			AMQP:    &settings.AMQPConf{},
			CacheDB: &settings.CacheDbConf{},
			DocDB:   &settings.DocDbConf{},
			RDB:     &settings.RDBConf{},
			custom:  make(map[string]interface{}),
		},
	}
	engine.StartTime = startTime
	engine.Config.engine = engine
	engine.Config.raw, _ = godotenv.Read()
	engine.Config.systemExternalEnv("app", engine.Config.App)
	engine.Config.systemExternalEnv("gin", engine.Config.Gin)
	engine.Config.systemExternalEnv("amqp", engine.Config.AMQP)
	engine.Config.systemExternalEnv("cachedb", engine.Config.CacheDB)
	engine.Config.systemExternalEnv("docdb", engine.Config.DocDB)
	engine.Config.systemExternalEnv("rdb", engine.Config.RDB)
	logger.SysLog.Info("[Engine] Environment Loaded")
	LiquidSDK.GetServer().SetCodeName(engine.Config.App.Codename)
	LiquidDb.GetInstance().SetCodeName(engine.Config.App.Codename)
	LiquidDb.GetInstance().ConnectCacheDbService(engine.Config.CacheDB)
	LiquidDb.GetInstance().ConnectDocDbService(engine.Config.DocDB)
	LiquidDb.GetInstance().InitializeSystemDocIndexes()
	LiquidSDK.GetServer().InitCodenameKey()
	engine.initializeGinEngine()
	engine.initializeFeatures()
	return engine
}

func (engine *Engine) initializeGinEngine() {
	gin.SetMode(engine.Config.Gin.RunMode)
	engine.ginEngine = gin.New()
	engine.ginEngine.Use(middlewares.GinLogger(1 * time.Second))
	engine.ginEngine.Use(gin.Recovery())
	engine.ginEngine.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	engine.RegisterGin("Foundation", foundation.Routers)
	engine.RegisterGin("HealthFoundation", health.Routers)
}

func (engine *Engine) initializeFeatures() {
	for name, mod := range feature_register.GetModuleList() {
		engine.RegisterFeature(name, mod)
	}
}

func (engine *Engine) GetGin() *gin.Engine {
	return engine.ginEngine
}

func (engine *Engine) Serve(opts ...*options.ServeOptions) {
	serveOptions := options.MergeServeOptions(opts...)

	if serveOptions.ServePort != nil {
		engine.Config.Gin.HttpPort = *serveOptions.ServePort
	}

	maxHeaderBytes := 1 << 20
	endPoint := fmt.Sprintf(":%d", engine.Config.Gin.HttpPort)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine.ginEngine,
		ReadTimeout:    time.Duration(engine.Config.Gin.ReadTimeout) * time.Millisecond,
		WriteTimeout:   time.Duration(engine.Config.Gin.WriteTimeout) * time.Millisecond,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		serveTime := time.Now()
		logger.SysLog.Infof("[Engine] Serving HTTP(%s) in %dms", endPoint, serveTime.Sub(engine.StartTime).Milliseconds())
		if err := server.ListenAndServe(); err != nil {
			logger.SysLog.Warnf("[Engine] Stop Serving (%s)", err)
		}
	}()

	signalChan := make(chan os.Signal)
	exitChan := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		logger.SysLog.Warnf("[Engine] Caught Signal(%03d)", sig)

		if err := server.Shutdown(context.Background()); err != nil {
			logger.SysLog.Warnf("[Engine] Shutdown Server with Error, %s", err)
		}

		exitChan <- true
	}()

	<-exitChan
	logger.SysLog.Warn("[Engine] Shutdown Server")
}
