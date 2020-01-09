package LiquidSDK

import (
	"context"
	"errors"
	"github.com/cesnow/LiquidEngine/Modules/LiquidRpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

func GameRpcConnection() (LiquidRpc.GameAdapterClient, error){
	keepAlive := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	conn, err := grpc.DialContext(
		ctx, "localhost:9999",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepAlive),
	)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := LiquidRpc.NewGameAdapterClient(conn)
	return c, nil
}

func (server *LiquidServer) SetGameRpcConnection(client LiquidRpc.GameAdapterClient){
	server.gameRpcConnection = client
	server.enableRpcTraffic = true
}

func (server *LiquidServer) GetGameRpcConnection() (LiquidRpc.GameAdapterClient, error){
	if server.enableRpcTraffic {
		return server.gameRpcConnection, nil
	}
	return nil, errors.New("")
}