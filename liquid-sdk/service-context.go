package LiquidSDK

import (
	"github.com/cesnow/liquid-engine/internal/database"
	"github.com/cesnow/liquid-engine/internal/msg-queue"
	"sync"
	"time"
)

type LiquidServer struct {
	CodeName        string
	LiquidKey       string
	LiquidKeyUpdate time.Time
	TokenExpireTime int

	featureDict       map[string]CommandSystem
	socketFeatureDict map[string]CommandSystem
	memberDict        map[string]MemberSystem

	liquidDocDb      *database.DocDB
	liquidCacheDb    *database.CacheDB
	liquidRelationDb *database.RDB
	liquidMsgQueue   MsgQueue.IAMQP
}

var liquidInstance *LiquidServer
var once sync.Once

func GetServer() *LiquidServer {
	once.Do(func() {
		liquidInstance = &LiquidServer{
			TokenExpireTime:   int(time.Hour.Seconds()),
			featureDict:       make(map[string]CommandSystem),
			socketFeatureDict: make(map[string]CommandSystem),
			memberDict:        make(map[string]MemberSystem),
		}
	})
	return liquidInstance
}
