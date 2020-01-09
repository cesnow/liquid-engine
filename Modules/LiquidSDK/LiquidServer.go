package LiquidSDK

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Settings"
	"strconv"
	"time"
)

var liquidKeyTemplate = "LiquidServer_%s"

func (server *LiquidServer) SetCodeName(codename string) {
	Logger.SysLog.Infof("[Engine] Codename -> %s", codename)
	server.CodeName = codename
}

func (server *LiquidServer) GetKeyStatic() string {
	if time.Now().Sub(server.LiquidKeyUpdate).Minutes() > 60 {
		return server.GetKey()
	}
	return server.LiquidKey
}

func (server *LiquidServer) GetKey() string {
	RedisLiquidKeyName := fmt.Sprintf(liquidKeyTemplate, server.CodeName)
	LiquidKey, GetKeyErr := server.GetCacheDb().Get(RedisLiquidKeyName)
	if GetKeyErr != nil {
		server.GenerateKey()
		server.LiquidKeyUpdate = time.Now()
		return server.LiquidKey
	}
	ReceivedLiquidKey := string(LiquidKey)
	if ReceivedLiquidKey != server.LiquidKey {
		server.LiquidKey = ReceivedLiquidKey
	}
	server.LiquidKeyUpdate = time.Now()
	return server.LiquidKey
}

func (server *LiquidServer) InitCodenameKey() {
	RedisLiquidKeyName := fmt.Sprintf(liquidKeyTemplate, server.CodeName)
	LiquidKey, GetKeyErr := server.GetCacheDb().Get(RedisLiquidKeyName)
	if GetKeyErr != nil {
		server.GenerateKey()
	} else {
		server.LiquidKey = string(LiquidKey)
	}
	Logger.SysLog.Infof("[Engine] System Key -> %s", server.LiquidKey)
}

func (server *LiquidServer) InitRpcTraffic(conf *Settings.AppConf) {
	if !conf.RpcCommandMode {
		return
	}
	Logger.SysLog.Infof("[Engine] RPC Enabled, Wait For Game RPC Ready ...")
	rpcClient, err := GameRpcConnection()
	if err == nil {
		server.gameRpcConnection = rpcClient
		server.enableRpcTraffic = true
		Logger.SysLog.Infof("[Engine] RPC Command Traffic Available")
	}
}

func (server *LiquidServer) GenerateKey() {
	conJunctions := "-LiquidSDK-"
	md5Generate := md5.New()
	var keyOriginConcat bytes.Buffer
	keyOriginConcat.Write([]byte(server.CodeName))
	keyOriginConcat.Write([]byte(conJunctions))
	keyOriginConcat.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
	md5Generate.Write(keyOriginConcat.Bytes())
	RedisLiquidKeyName := fmt.Sprintf(liquidKeyTemplate, server.CodeName)
	LiquidKey := hex.EncodeToString(md5Generate.Sum(nil))
	SaveKey2RedisErr := server.GetCacheDb().SetString(RedisLiquidKeyName, LiquidKey, -1)
	if SaveKey2RedisErr != nil {
		Logger.SysLog.Errorf("[System] Save System Key To Redis Error, %s", SaveKey2RedisErr)
	}
	server.LiquidKey = LiquidKey
}
