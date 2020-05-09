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

func (fb *FacebookMemberSystem) Validate(fromId, fromToken string) bool {
	Logger.SysLog.Infof("[Member|Validate|Facebook] FromId: %s, FromToken: %s", fromId, fromToken)
	return false
}