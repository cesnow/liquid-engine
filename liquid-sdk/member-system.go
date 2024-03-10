package LiquidSDK

type MemberSystem interface {
	Register(
		fromType,
		account,
		password,
		platform string,
		extraArgs interface{}) (status int, error string)
	Validate(
		fromId,
		fromToken,
		platform string,
		extraArgs interface{}) (valid bool, msg string, overrideFromId string)
}

func (server *LiquidServer) RegisterMember(memberType string, MemberInstance MemberSystem) bool {
	if _, find := server.memberDict[memberType]; find {
		return false
	}
	server.memberDict[memberType] = MemberInstance
	return true
}

func (server *LiquidServer) GetMemberSystem(memberType string) MemberSystem {
	if member, find := server.memberDict[memberType]; find {
		return member
	}
	return nil
}
