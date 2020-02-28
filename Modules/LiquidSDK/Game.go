package LiquidSDK

import (
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"io"
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

func (e *RpcFeature) Command(in LiquidRpc.GameAdapter_CommandServer) error {

	ctx := in.Context()

	for {
		// exit if context is done or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := in.Recv()
		if err == io.EOF {
			// return will close stream from server side
			return nil
		}
		if err != nil {
			Logger.SysLog.Warnf("[RPC|Cmd] Failed, %+v", err)
			continue
		}

		Logger.SysLog.Infof("[RPC|Cmd] ID: %s Name: %s", req.CmdId, req.CmdName)
		gameFeature := GetServer().GetGameFeature(req.CmdId)
		if gameFeature == nil {
			Logger.SysLog.Warnf("[RPC|Cmd] Can't Find Feature `%s`", req.CmdId)
			if err := in.Send(&LiquidRpc.RespCmd{CmdData: nil,}); err != nil {
				Logger.SysLog.Warnf("[RPC|Cmd] Reply Failed, %s", err)
			}
		}
		runCommandData := gameFeature.RunRpcCommand(req)
		marshalCommandData, _ := json.Marshal(runCommandData)
		result := &LiquidRpc.RespCmd{CmdData: marshalCommandData,}
		if err := in.Send(result); err != nil {
			Logger.SysLog.Warnf("[RPC|Cmd] Reply Failed, %s", err)
		}
	}
}
