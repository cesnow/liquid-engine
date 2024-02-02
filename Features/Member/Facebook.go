package Member

import (
	"errors"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
)

type FacebookMemberSystem struct {
}

func (fb *FacebookMemberSystem) Register(fromType, account, password, platform string, extraArgs interface{}) (status int, error string) {
	err := errors.New("facebook account can't register by server, visit: https://www.facebook.com/")
	return 0, err.Error()
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
