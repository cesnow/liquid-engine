package LiquidDB

import (
	"github.com/cesnow/liquid-engine/internal/database"
	"github.com/cesnow/liquid-engine/internal/msg-queue"
	"github.com/cesnow/liquid-engine/logger"
	"github.com/cesnow/liquid-engine/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func (db *LiquidDB) SetCodeName(codename string) {
	if codename == "" {
		logger.SysLog.Errorf("[Engine] Codename is empty !!!")
		os.Exit(100)
	}
	logger.SysLog.Infof("[Engine] Codename -> %s", codename)
	db.CodeName = codename
}

func (db *LiquidDB) ConnectRDbService(config *settings.RDBConf) {
	var err error
	client, err := database.ConnectWithRDB(config)
	if err != nil {
		logger.SysLog.Errorf("[RelationDB] Try To Connect Relation Database Failed -> (%s)", err)
	}
	db.liquidRelationDb = client
}

func (db *LiquidDB) ConnectDocDbService(config *settings.DocDbConf) {
	client, err := database.ConnectWithDocDB(config)
	if err != nil {
		logger.SysLog.Errorf("[DocumentDB] Try To Connect Document Database Failed -> (%s)", err)
		os.Exit(102)
	}
	db.liquidDocDb = client
}

func (db *LiquidDB) InitializeSystemDocIndexes() {
	logger.SysLog.Infof("[DocumentDB] Ensure System Collection Indexes .. (may take some moment)")
	db.EnsureLiquidMemberIndexes()
	db.EnsureLiquidUserIndexes()
	db.EnsureLiquidAdminIndexes()
}

func (db *LiquidDB) ConnectCacheDbService(config *settings.CacheDbConf) {
	client, err := database.ConnectWithCacheDB(config)
	if err != nil {
		logger.SysLog.Errorf("[CacheDb] Try To Connect Cache Database Failed -> (%s)", err)
		os.Exit(101)
	}
	db.liquidCacheDb = client
}

func (db *LiquidDB) ConnectMsgQueueService(config *settings.AMQPConf) {
	mqClient, err := MsgQueue.Connect(config)
	if err != nil {
		logger.SysLog.Errorf("[MsgQueue] Try To Connect Message Queue Failed -> (%s)", err)
	}
	db.liquidMsgQueue = mqClient
}

func (db *LiquidDB) GetRdb() *database.RDB {
	if db.liquidRelationDb == nil {
		return nil
	}
	return db.liquidRelationDb
}

func (db *LiquidDB) GetDocDb() *database.DocDB {
	if db.liquidDocDb == nil {
		return nil
	}
	return db.liquidDocDb
}

func (db *LiquidDB) GetDocColl(collection string) *mongo.Collection {
	if db.liquidDocDb == nil {
		return nil
	}
	return db.liquidDocDb.Database(db.CodeName).Collection(collection)
}

func (db *LiquidDB) GetDocCollWithDb(dbName, collection string) *mongo.Collection {
	if db.liquidDocDb == nil {
		return nil
	}
	return db.liquidDocDb.Database(dbName).Collection(collection)
}

func (db *LiquidDB) GetCacheDb() *database.CacheDB {
	if db.liquidCacheDb == nil {
		return nil
	}
	return db.liquidCacheDb
}

func (db *LiquidDB) GetMsgQueueV1() *MsgQueue.AMQPv1 {
	if db.liquidMsgQueue == nil || db.liquidMsgQueue.GetProtocolVersion() == 0 {
		return nil
	}
	return db.liquidMsgQueue.(*MsgQueue.AMQPv1)
}

func (db *LiquidDB) GetMsgQueueV0() *MsgQueue.AMQPv0 {
	if db.liquidMsgQueue == nil || db.liquidMsgQueue.GetProtocolVersion() == 1 {
		return nil
	}
	return db.liquidMsgQueue.(*MsgQueue.AMQPv0)
}
