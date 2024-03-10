package LiquidDB

import (
	"github.com/cesnow/liquid-engine/internal/database"
	"github.com/cesnow/liquid-engine/internal/msg-queue"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLiquidMemberCol() *mongo.Collection {
	return GetInstance().GetLiquidMemberCol()
}

func GetLiquidUserCol() *mongo.Collection {
	return GetInstance().GetLiquidUserCol()
}

func GetLiquidUserDataCol() *mongo.Collection {
	return GetInstance().GetLiquidUserDataCol()
}

func GetLiquidAdminCol() *mongo.Collection {
	return GetInstance().GetLiquidAdminCol()
}

func EnsureLiquidMemberIndexes() {
	GetInstance().EnsureLiquidMemberIndexes()
}
func EnsureLiquidUserIndexes() {
	GetInstance().EnsureLiquidUserIndexes()
}
func EnsureLiquidAdminIndexes() {
	GetInstance().EnsureLiquidAdminIndexes()
}

func GetRdb() *database.RDB {
	return GetInstance().GetRdb()
}

func GetDocDb() *database.DocDB {
	return GetInstance().GetDocDb()
}

func GetDocColl(collection string) *mongo.Collection {
	return GetInstance().GetDocColl(collection)
}

func GetDocCollWithDb(dbName, collection string) *mongo.Collection {
	return GetInstance().GetDocCollWithDb(dbName, collection)
}

func GetCacheDb() *database.CacheDB {
	return GetInstance().GetCacheDb()
}

func GetMsgQueueV1() *MsgQueue.AMQPv1 {
	return GetInstance().GetMsgQueueV1()
}

func GetMsgQueueV0() *MsgQueue.AMQPv0 {
	return GetInstance().GetMsgQueueV0()
}
