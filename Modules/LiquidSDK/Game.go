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

func (e *RpcFeature) Login(ctx context.Context, req *LiquidRpc.ReqLogin) (resp *LiquidRpc.RespLogin, err error) {
	Logger.SysLog.Infof("[RPC|Recv|Login] From: %s, Id: %s, Token: %s",
		req.FromType,
		req.FromId,
		req.FromToken,
	)
	validResp := &LiquidRpc.RespLogin{Valid: false, Msg: "", OverrideFromId: ""}
	member := GetServer().GetMemberSystem(req.FromType)
	if member == nil {
		Logger.SysLog.Warnf("[RPC|Recv|Login] Member system not found `%s`", req.FromType)
		return validResp, nil
	}
	var extraArgs interface{}
	_ = json.Unmarshal(req.ExtraArgs, &extraArgs)
	resultValidate, resultMsg, resultOverrideFromId := member.Validate(
		req.FromId,
		req.FromToken,
		req.Platform,
		extraArgs,
	)
	return &LiquidRpc.RespLogin{
		Valid:          resultValidate,
		Msg:            resultMsg,
		OverrideFromId: resultOverrideFromId,
	}, nil
}
