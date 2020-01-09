package GameFoundation

import (
	"context"
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
)

func GRpcCommand(command *LiquidSDK.CmdCommand) ([]byte, error) {

	c, err := LiquidSDK.GetServer().GetGameRpcConnection()
	if err != nil {
		return nil, err
	}
	marshalCmdData, _ := json.Marshal(command.CmdData)
	r, err := c.Command(context.Background(), &LiquidRpc.CmdCommand{
		UserID:    *command.LiquidId,
		UserToken: *command.LiquidToken,
		Platform:  *command.Platform,
		CmdId:     *command.CmdId,
		CmdSn:     uint64(*command.CmdSn),
		CmdName:   *command.CmdName,
		CmdData:   marshalCmdData,
	})
	if err != nil {
		return nil, err
	}
	return r.CmdData, nil
}
