package Member

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
)

type LineMemberSystem struct {

}

func NewLineMemberSystem() LiquidSDK.IMemberSystem {
	lineMember := new(LineMemberSystem)
	return lineMember
}

func (fb *LineMemberSystem) Validate(fromId, fromToken string) bool {
	Logger.SysLog.Infof("[Member|Validate|Line] FromId: %s, FromToken: %s", fromId, fromToken)
	return false
}