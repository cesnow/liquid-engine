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

type RpcCmdCommand struct{}

func (e *RpcCmdCommand) Command(ctx context.Context, req *LiquidRpc.RpcCmdCommand) (resp *LiquidRpc.RpcCmdCommandReply, err error) {
	Logger.SysLog.Infof("[RPC][Received Game Command] ID: %s Name: %s", req.CmdId, req.CmdName)
	gameFeature := GetServer().GetGameFeature(req.CmdId)
	if gameFeature == nil {
		Logger.SysLog.Warnf("[RPC][Game Feature] Can't Find Feature `%s`", req.CmdId)
		return &LiquidRpc.RpcCmdCommandReply{
			CmdData: nil,
		}, nil
	}
	command := &CmdCommand{
		LiquidId:    &req.UserID,
		LiquidToken: nil,
		Platform:    &req.Platform,
		CmdId:       &req.CmdId,
		CmdSn:       nil,
		CmdName:     &req.CmdName,
		CmdData:     req.CmdData,
	}
	runCommandData := gameFeature.RunCommand(command)
	marshalCommandData, _ := json.Marshal(runCommandData)
	return &LiquidRpc.RpcCmdCommandReply{
		CmdData: marshalCommandData,
	}, nil
}
