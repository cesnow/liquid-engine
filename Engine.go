package LiquidEngine

import (
	"context"
	"fmt"
	"github.com/cesnow/LiquidEngine/Features/Foundation"
	"github.com/cesnow/LiquidEngine/Features/GameFoundation"
	"github.com/cesnow/LiquidEngine/Features/HealthFoundation"
	_ "github.com/cesnow/LiquidEngine/Features/HealthFoundation"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Middlewares"
	"github.com/cesnow/LiquidEngine/Modules/LiquidModule"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/cesnow/LiquidEngine/Options"
	"github.com/cesnow/LiquidEngine/Settings"
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
	Logger.SysLog.Info("[Engine] Starting")
	engine := &Engine{
		Config: &Config{
			App:     &Settings.AppConf{},
			Gin:     &Settings.GinConf{},
			AMQP:    &Settings.AMQPConf{},
			CacheDB: &Settings.CacheDbConf{},
			DocDB:   &Settings.DocDbConf{},
			RDB:     &Settings.RDBConf{},
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
	Logger.SysLog.Info("[Engine] Environment Loaded")
	LiquidSDK.GetServer().SetCodeName(engine.Config.App.Codename)
	LiquidSDK.GetServer().ConnectCacheDbService(engine.Config.CacheDB)
	LiquidSDK.GetServer().ConnectDocDbService(engine.Config.DocDB)
	LiquidSDK.GetServer().InitializeSystemDocIndexes()
	LiquidSDK.GetServer().InitCodenameKey()
	engine.initializeGinEngine()
	engine.initializeFeatures()
	return engine
}

func (engine *Engine) initializeGinEngine() {
	gin.SetMode(engine.Config.Gin.RunMode)
	engine.ginEngine = gin.New()
	engine.ginEngine.Use(Middlewares.GinLogger(1 * time.Second))
	engine.ginEngine.Use(gin.Recovery())
	engine.ginEngine.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	engine.RegisterGin("Foundation", Foundation.Routers)
	engine.RegisterGin("GameFoundation", GameFoundation.Routers)
	engine.RegisterGin("HealthFoundation", HealthFoundation.Routers)
}

func (engine *Engine) initializeFeatures() {
	for name, mod := range LiquidModule.GetModuleList() {
		engine.RegisterGame(name, mod)
	}
}

func (engine *Engine) GetGin() *gin.Engine {
	return engine.ginEngine
}

func (engine *Engine) Serve(opts ...*Options.ServeOptions) {
	serveOptions := Options.MergeServeOptions(opts...)

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
		Logger.SysLog.Infof("[Engine] Serving HTTP(%s) in %dms", endPoint, serveTime.Sub(engine.StartTime).Milliseconds())
		if err := server.ListenAndServe(); err != nil {
			Logger.SysLog.Warnf("[Engine] Stop Serving (%s)", err)
		}
	}()

	signalChan := make(chan os.Signal)
	exitChan := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		Logger.SysLog.Warnf("[Engine] Caught Signal(%03d)", sig)

		if err := server.Shutdown(context.Background()); err != nil {
			Logger.SysLog.Warnf("[Engine] Shutdown Server with Error, %s", err)
		}

		exitChan <- true
	}()

	<-exitChan
	Logger.SysLog.Warn("[Engine] Shutdown Server")
}
