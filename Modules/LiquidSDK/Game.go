package LiquidSDK

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
