package LiquidSDK

type IMemberSystem interface {
	Validate(
		fromId,
		fromToken,
		platform string,
		extraArgs interface{}) (valid bool, msg string, overrideFromId string)
}

func (server *LiquidServer) RegisterMember(memberType string, MemberInstance IMemberSystem) bool {
	if _, find := server.memberDict[memberType]; find {
		return false
	}
	server.memberDict[memberType] = MemberInstance
	return true
}

func (server *LiquidServer) GetMemberSystem(memberType string) IMemberSystem {
	if member, find := server.memberDict[memberType]; find {
		return member
	}
	return nil
}
