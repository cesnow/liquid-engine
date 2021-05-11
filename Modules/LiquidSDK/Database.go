package LiquidSDK

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/Database"
	"github.com/cesnow/LiquidEngine/Modules/MsgQueue"
	"github.com/cesnow/LiquidEngine/Settings"
	"os"
)

func (server *LiquidServer) ConnectRDbService(config *Settings.RDBConf) {
	var err error
	client, err := Database.ConnectWithRDB(config)
	if err != nil {
		Logger.SysLog.Errorf("[RelationDB] Try To Connect Relation Database Failed -> (%s)", err)
	}
	server.liquidRelationDb = client
}

func (server *LiquidServer) ConnectDocDbService(config *Settings.DocDbConf) {
	client, err := Database.ConnectWithDocDB(config)
	if err != nil {
		Logger.SysLog.Errorf("[DocumentDB] Try To Connect Document Database Failed -> (%s)", err)
		os.Exit(102)
	}
	server.liquidDocDb = client
}

func (server *LiquidServer) InitializeSystemDocIndexes() {
	Logger.SysLog.Infof("[DocumentDB] Ensure System Collection Indexes .. (may take some moment)")
	server.EnsureLiquidMemberIndexes()
	server.EnsureLiquidUserIndexes()
	server.EnsureLiquidAdminIndexes()
}

func (server *LiquidServer) ConnectCacheDbService(config *Settings.CacheDbConf) {
	client, err := Database.ConnectWithCacheDB(config)
	if err != nil {
		Logger.SysLog.Errorf("[CacheDb] Try To Connect Cache Database Failed -> (%s)", err)
		os.Exit(101)
	}
	server.liquidCacheDb = client
}

func (server *LiquidServer) ConnectMsgQueueService(config *Settings.AMQPConf) {
	mqClient, err := MsgQueue.Connect(config)
	if err != nil {
		Logger.SysLog.Errorf("[MsgQueue] Try To Connect Message Queue Failed -> (%s)", err)
	}
	server.liquidMsgQueue = mqClient
}

func (server *LiquidServer) GetRdb() *Database.RDB {
	if server.liquidRelationDb == nil {
		return nil
	}
	return server.liquidRelationDb
}

func (server *LiquidServer) GetDocDb() *Database.DocDB {
	if server.liquidDocDb == nil {
		return nil
	}
	return server.liquidDocDb
}

func (server *LiquidServer) GetCacheDb() *Database.CacheDB {
	if server.liquidCacheDb == nil {
		return nil
	}
	return server.liquidCacheDb
}

func (server *LiquidServer) GetMsgQueueV1() *MsgQueue.AMQPv1 {
	if server.liquidMsgQueue == nil || server.liquidMsgQueue.GetProtocolVersion() == 0 {
		return nil
	}
	return server.liquidMsgQueue.(*MsgQueue.AMQPv1)
}

func (server *LiquidServer) GetMsgQueueV0() *MsgQueue.AMQPv0 {
	if server.liquidMsgQueue == nil || server.liquidMsgQueue.GetProtocolVersion() == 1 {
		return nil
	}
	return server.liquidMsgQueue.(*MsgQueue.AMQPv0)
}
