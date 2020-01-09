package GameFoundation

import (
	"context"
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
)

func GRpcCommand(command *LiquidSDK.CmdCommand) ([]byte, error) {
	c := LiquidSDK.GetServer().GetGameRpcConnection()
	marshalCmdData, _ := json.Marshal(command.CmdData)
	r, err := c.Command(context.Background(), &LiquidRpc.RpcCmdCommand{
		UserID:   *command.LiquidId,
		Platform: *command.Platform,
		CmdId:    *command.CmdId,
		CmdName:  *command.CmdName,
		CmdData:  marshalCmdData,
	})
	if err != nil {
		Logger.SysLog.Warnf("[Engine] Game Rpc Traffic Failed, %+v", err)
		return nil, err
	}
	return r.CmdData, nil
}
