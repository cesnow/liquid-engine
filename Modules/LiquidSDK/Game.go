package LiquidSDK

import (
	"context"
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
)

func (server *LiquidServer) RegisterGameFeature(GameName string, GameInstance IGameSystem) bool {
	if _, find := server.systemGameDict[GameName]; find {
		return false
	}
	server.systemGameDict[GameName] = GameInstance
	return true
}

func (server *LiquidServer) GetGameFeature(GameName string) IGameSystem {
	if game, find := server.systemGameDict[GameName]; find {
		return game
	}
	return nil
}

type RpcFeature struct{}

func (e *RpcFeature) Command(ctx context.Context, req *LiquidRpc.ReqCmd) (resp *LiquidRpc.RespCmd, err error) {
	Logger.SysLog.Infof("[RPC][Received Game Command] ID: %s Name: %s", req.CmdId, req.CmdName)
	gameFeature := GetServer().GetGameFeature(req.CmdId)
	if gameFeature == nil {
		Logger.SysLog.Warnf("[RPC][Game Feature] Can't Find Feature `%s`", req.CmdId)
		return &LiquidRpc.RespCmd{
			CmdData: nil,
		}, nil
	}
	runCommandData := gameFeature.RunRpcCommand(req)
	marshalCommandData, _ := json.Marshal(runCommandData)
	return &LiquidRpc.RespCmd{
		CmdData: marshalCommandData,
	}, nil
}
