package LiquidDB

import (
	"github.com/cesnow/liquid-engine/internal/database"
	"github.com/cesnow/liquid-engine/internal/msg-queue"
	"sync"
)

type LiquidDB struct {
	CodeName string

	liquidDocDb      *database.DocDB
	liquidCacheDb    *database.CacheDB
	liquidRelationDb *database.RDB
	liquidMsgQueue   MsgQueue.IAMQP
}

var dbInstance *LiquidDB
var once sync.Once

func GetInstance() *LiquidDB {
	once.Do(func() {
		dbInstance = &LiquidDB{}
	})
	return dbInstance
}
