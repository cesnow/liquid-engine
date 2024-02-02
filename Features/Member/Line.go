package Member

import (
	"errors"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
)

type LineMemberSystem struct {
}

func (fb *LineMemberSystem) Register(fromType, account, password, platform string, extraArgs interface{}) (status int, error string) {
	err := errors.New("line account can't register by server, visit: https://www.line.com/")
	return 0, err.Error()
}

func NewLineMemberSystem() LiquidSDK.IMemberSystem {
	lineMember := new(LineMemberSystem)
	return lineMember
}

func (fb *LineMemberSystem) Validate(
	fromId, fromToken, platform string,
	args interface{},
) (valid bool, msg string, overrideFromId string) {
	Logger.SysLog.Infof("[Member|Validate|Line] FromId: %s, FromToken: %s", fromId, fromToken)
	valid = false
	msg = ""
	overrideFromId = ""
	return
}
