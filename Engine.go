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
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/cesnow/LiquidEngine/Options"
	"github.com/cesnow/LiquidEngine/Settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
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
	LiquidSDK.GetServer().InitRpcTraffic(engine.Config.App)
	engine.initGinEngine()
	return engine
}

func (engine *Engine) initGinEngine() {
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

func (engine *Engine) RpcModeServe() {

	var keepAliveEP = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	var keepAliveSP = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}

	var gRpcServer *grpc.Server
	go func() {
		serveTime := time.Now()
		apiListener, err := net.Listen("tcp", fmt.Sprintf(":%d", engine.Config.App.RpcBindPort))
		if err != nil {
			Logger.SysLog.Warnf("[Engine] gRPC Mode Serve Failed (%s)", err)
			return
		}
		gRpcServer = grpc.NewServer(
			grpc.InitialWindowSize(64*1024*2),
			grpc.InitialConnWindowSize(64*1024*2),
			grpc.KeepaliveEnforcementPolicy(keepAliveEP),
			grpc.KeepaliveParams(keepAliveSP),
			grpc.MaxRecvMsgSize(50*1024*1024),
			grpc.MaxSendMsgSize(50*1024*1024),
		)
		LiquidRpc.RegisterGameAdapterServer(gRpcServer, &LiquidSDK.RpcFeature{})
		reflection.Register(gRpcServer)
		Logger.SysLog.Infof("[Engine] Serving gRpc(:%d) in %dms",
			engine.Config.App.RpcBindPort,
			serveTime.Sub(engine.StartTime).Milliseconds(),
		)
		if err := gRpcServer.Serve(apiListener); err != nil {
			Logger.SysLog.Warnf("[Engine] gRPC Mode Serve Failed (%s)", err)
			return
		}
	}()

	signalChan := make(chan os.Signal)
	exitChan := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		Logger.SysLog.Warnf("[Engine] Caught Signal(%03d)", sig)
		gRpcServer.Stop()
		exitChan <- true
	}()

	<-exitChan
	Logger.SysLog.Warn("[Engine] Shutdown Server")
}
