package GameFoundation

import (
	"context"
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
)

func GRpcCommand(command *LiquidSDK.CmdCommand, direct bool) ([]byte, error) {
	c := LiquidSDK.GetServer().GetGameRpcConnection()
	marshalCmdData, _ := json.Marshal(command.CmdData)
	UserID := ""
	Platform := "main"
	if command.LiquidId != nil {
		UserID = *command.LiquidId
	}
	if command.Platform != nil {
		Platform = *command.Platform
	}

	r, err := c.Command(context.Background(), &LiquidRpc.ReqCmd{
		UserID:   UserID,
		Platform: Platform,
		CmdId:    *command.CmdId,
		CmdName:  *command.CmdName,
		CmdData:  marshalCmdData,
		Direct:   direct,
	})

	if err != nil {
		Logger.SysLog.Warnf("[RPC|Engine] Game Traffic Failed, %+v", err)
		return nil, err
	}
	return r.CmdData, nil

}

func GRpcLogin(command *LiquidSDK.CmdAccount) (*LiquidRpc.RespLogin, error) {
	c := LiquidSDK.GetServer().GetGameRpcConnection()

	marshalExtraArgs, _ := json.Marshal(command.ExtraArgs)
	r, err := c.Login(context.Background(), &LiquidRpc.ReqLogin{
		FromType:  command.FromType,
		FromId:    command.FromId,
		FromToken: command.FromToken,
		Platform:  command.Platform,
		ExtraArgs: marshalExtraArgs,
	})

	if err != nil {
		Logger.SysLog.Warnf("[RPC|Engine] Login Traffic Failed, %+v", err)
		return nil, err
	}
	return r, nil

}
