package LiquidSDK

import (
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
	"time"
)

func GameRpcConnection(remoteIp string) (LiquidRpc.GameAdapterClient, error) {

	keepAlive := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             3 * time.Second,
		PermitWithoutStream: true,
	}

	conn, err := grpc.Dial(
		remoteIp,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepAlive),
		grpc.WithDefaultCallOptions(
			grpc.UseCompressor("gzip"),
			grpc.MaxCallSendMsgSize(50*1024*1024),
			grpc.MaxCallRecvMsgSize(50*1024*1024),
		),
	)
	if err != nil {
		return nil, err
	}

	c := LiquidRpc.NewGameAdapterClient(conn)
	return c, nil
}

func (server *LiquidServer) GetRpcTrafficEnabled() bool {
	return server.enableRpcTraffic
}

func (server *LiquidServer) GetGameRpcConnection() LiquidRpc.GameAdapterClient {
	return server.gameRpcConnection
}
