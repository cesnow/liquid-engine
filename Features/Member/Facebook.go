package Member

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
)

type FacebookMemberSystem struct {
}

func NewFacebookMemberSystem() LiquidSDK.IMemberSystem {
	fbMember := new(FacebookMemberSystem)
	return fbMember
}

func (fb *FacebookMemberSystem) Validate(
	fromId, fromToken, platform string,
	args interface{},
) (valid bool, msg string, overrideFromId string) {
	Logger.SysLog.Infof("[Member|Validate|Facebook] FromId: %s, FromToken: %s", fromId, fromToken)
	valid = false
	msg = ""
	overrideFromId = ""
	return
}
