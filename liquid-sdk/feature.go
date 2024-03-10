package LiquidSDK

func (server *LiquidServer) RegisterFeature(Name string, Instance CommandSystem) bool {
	if _, find := server.featureDict[Name]; find {
		return false
	}
	server.featureDict[Name] = Instance
	return true
}

func (server *LiquidServer) GetFeature(Name string) CommandSystem {
	if feature, find := server.featureDict[Name]; find {
		return feature
	}
	return nil
}
