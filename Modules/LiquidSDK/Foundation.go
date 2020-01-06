package LiquidSDK

import (
	"github.com/cesnow/LiquidEngine/Modules/Database"
	"github.com/cesnow/LiquidEngine/Modules/MsgQueue"
	"sync"
)

type ILiquidServer interface {
}

type LiquidServer struct {
	CodeName        string
	LiquidKey       string
	TokenExpireTime int

	systemGameDict map[string]IGameSystem
	socketGameDict map[string]IGameSystem

	liquidDocDb      *Database.DocDB
	liquidCacheDb    *Database.CacheDB
	liquidRelationDb *Database.RDB
	liquidMsgQueue   MsgQueue.IAMQP
}

var liquidInstance *LiquidServer
var once sync.Once

func GetServer() *LiquidServer {
	once.Do(func() {
		liquidInstance = &LiquidServer{
			TokenExpireTime: 1800,
			systemGameDict:  make(map[string]IGameSystem),
			socketGameDict:  make(map[string]IGameSystem),
		}
	})
	return liquidInstance
}
