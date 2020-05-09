package LiquidSDK

import (
	"github.com/cesnow/LiquidEngine/Modules/Database"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"github.com/cesnow/LiquidEngine/Modules/MsgQueue"
	"sync"
	"time"
)

type ILiquidServer interface {
}

type LiquidServer struct {
	CodeName        string
	LiquidKey       string
	LiquidKeyUpdate time.Time
	TokenExpireTime int

	systemGameDict map[string]IGameSystem
	socketGameDict map[string]IGameSystem
	memberDict     map[string]IMemberSystem

	enableRpcTraffic  bool
	gameRpcConnection LiquidRpc.GameAdapterClient
	liquidDocDb       *Database.DocDB
	liquidCacheDb     *Database.CacheDB
	liquidRelationDb  *Database.RDB
	liquidMsgQueue    MsgQueue.IAMQP
}

var liquidInstance *LiquidServer
var once sync.Once

func GetServer() *LiquidServer {
	once.Do(func() {
		liquidInstance = &LiquidServer{
			TokenExpireTime:  int(time.Hour.Seconds()),
			enableRpcTraffic: false,
			systemGameDict:   make(map[string]IGameSystem),
			socketGameDict:   make(map[string]IGameSystem),
			memberDict:       make(map[string]IMemberSystem),
		}
	})
	return liquidInstance
}
