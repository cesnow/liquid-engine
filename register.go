package LiquidEngine

import (
	LiquidDB "github.com/cesnow/liquid-engine/liquid-db"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

func (engine *Engine) RegisterGin(featureName string, router func(*gin.Engine)) {
	initTime := time.Now()
	router(engine.ginEngine)
	diffTime := time.Since(initTime).Microseconds()
	logger.SysLog.Infof("[Engine] Register HTTP Feature (%s) in %dµs", featureName, diffTime)
}

func (engine *Engine) RegisterFeature(Cmd string, Instance LiquidSDK.CommandSystem) {
	// Register Router
	if Instance.GetRouterFunc() != nil {
		cmdGroup := engine.ginEngine.Group("/api/" + Cmd)
		Instance.GetRouterFunc()(cmdGroup)
	}
	// Register Command
	registered := LiquidSDK.GetServer().RegisterFeature(Cmd, Instance)
	if !registered {
		logger.SysLog.Warnf(
			"[Engine] Register Feature (Name: `%s`, Cmd: `%s`) Failed, Command Duplicate!",
			reflect.TypeOf(Instance),
			Cmd,
		)
		return
	}
	logger.SysLog.Infof(
		"[Engine] Feature Registered (Name: `%s`, Cmd: `%s`)",
		reflect.TypeOf(Instance),
		Cmd,
	)
}

func (engine *Engine) RegisterMember(MemberType string, MemberInstance LiquidSDK.MemberSystem) {
	registered := LiquidSDK.GetServer().RegisterMember(MemberType, MemberInstance)
	if !registered {
		logger.SysLog.Warnf(
			"[Engine] Register Member (Name: `%s`, Cmd: `%s`) Failed, Member Type Duplicate!",
			reflect.TypeOf(MemberInstance),
			MemberType,
		)
		return
	}
	logger.SysLog.Infof(
		"[Engine] Member Registered (Name: `%s`, From: `%s`)",
		reflect.TypeOf(MemberInstance),
		MemberType,
	)
}

func (engine *Engine) UsingRDBService() {
	LiquidDB.GetInstance().ConnectRDbService(engine.Config.RDB)
}

func (engine *Engine) UsingDocumentDBService() {
	LiquidDB.GetInstance().ConnectDocDbService(engine.Config.DocDB)
}

func (engine *Engine) UsingCacheDBService() {
	LiquidDB.GetInstance().ConnectCacheDbService(engine.Config.CacheDB)
}

func (engine *Engine) UsingMsgQueueService() {
	LiquidDB.GetInstance().ConnectMsgQueueService(engine.Config.AMQP)
}
