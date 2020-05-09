package LiquidEngine

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

func (engine *Engine) RegisterGin(featureName string, router func(*gin.Engine)) {
	initTime := time.Now()
	router(engine.ginEngine)
	diffTime := time.Since(initTime).Microseconds()
	Logger.SysLog.Infof("[Engine] Register HTTP Feature (%s) in %dµs", featureName, diffTime)
}

func (engine *Engine) RegisterGame(GameCmd string, GameInstance LiquidSDK.IGameSystem) {
	registered := LiquidSDK.GetServer().RegisterGameFeature(GameCmd, GameInstance)
	if !registered {
		Logger.SysLog.Warnf(
			"[Engine] Register Game Feature (Name: `%s`, GameCmd: `%s`) Failed, Game Command Duplicate!",
			reflect.TypeOf(GameInstance),
			GameCmd,
		)
		return
	}
	Logger.SysLog.Infof(
		"[Engine] Game Feature Registered (Name: `%s`, GameCmd: `%s`)",
		reflect.TypeOf(GameInstance),
		GameCmd,
	)
}

func (engine *Engine) RegisterMember(MemberType string, MemberInstance LiquidSDK.IMemberSystem){
	registered := LiquidSDK.GetServer().RegisterMember(MemberType, MemberInstance)
	if !registered {
		Logger.SysLog.Warnf(
			"[Engine] Register Member (Name: `%s`, GameCmd: `%s`) Failed, Member Type Duplicate!",
			reflect.TypeOf(MemberInstance),
			MemberType,
		)
		return
	}
	Logger.SysLog.Infof(
		"[Engine] Member Registered (Name: `%s`, From: `%s`)",
		reflect.TypeOf(MemberInstance),
		MemberType,
	)
}

func (engine *Engine) UsingRDBService() {
	LiquidSDK.GetServer().ConnectRDbService(engine.Config.RDB)
}

func (engine *Engine) UsingDocumentDBService() {
	LiquidSDK.GetServer().ConnectDocDbService(engine.Config.DocDB)
}

func (engine *Engine) UsingCacheDBService() {
	LiquidSDK.GetServer().ConnectCacheDbService(engine.Config.CacheDB)
}

func (engine *Engine) UsingMsgQueueService() {
	LiquidSDK.GetServer().ConnectMsgQueueService(engine.Config.AMQP)
}
