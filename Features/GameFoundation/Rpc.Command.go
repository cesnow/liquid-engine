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
		Logger.SysLog.Warnf("[Engine] Game Rpc Traffic Failed, %+v", err)
		return nil, err
	}
	return r.CmdData, nil

}
