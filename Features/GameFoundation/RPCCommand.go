package GameFoundation

import (
	"context"
	"encoding/json"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"io"
)

func GRpcCommand(command *LiquidSDK.CmdCommand, direct bool) (respCmd []byte, respErr error) {
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

	stream, err := c.Command(context.Background())
	if err != nil {
		Logger.SysLog.Warnf("[Engine] Game Rpc Traffic Failed, %+v", err)
		respCmd = nil
		respErr = err
		return
	}

	var replyCmdData []byte
	// ctx := stream.Context()
	done := make(chan bool)

	// first goroutine sends random increasing numbers to stream
	// and closes it after 10 iterations
	go func() {
		if err := stream.Send(&LiquidRpc.ReqCmd{
			UserID:   UserID,
			Platform: Platform,
			CmdId:    *command.CmdId,
			CmdName:  *command.CmdName,
			CmdData:  marshalCmdData,
			Direct:   direct,
		}); err != nil {
			Logger.SysLog.Fatalf("[Engine] Can't send, %+v", err)
		}
		if err := stream.CloseSend(); err != nil {
			Logger.SysLog.Warnf("[Engine] Close Send Failed, %+v", err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				Logger.SysLog.Errorf("[Engine] Failed to receive reply, %+v", err)
			}
			replyCmdData = resp.CmdData
		}
	}()

	// third goroutine closes done channel if context is done
	//go func() {
	//	Logger.SysLog.Error("1")
	//	<-ctx.Done()
	//	Logger.SysLog.Error("2")
	//	if err := ctx.Err(); err != nil {
	//		Logger.SysLog.Warnf("[Engine] %+v", err)
	//	}
	//	close(done)
	//}()

	<-done
	respCmd = replyCmdData
	respErr = nil
	return
}
