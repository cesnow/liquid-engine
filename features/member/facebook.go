package member

import (
	"errors"
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"

	"github.com/cesnow/liquid-engine/logger"
)

type FacebookMemberSystem struct {
}

func (fb *FacebookMemberSystem) Register(fromType, account, password, platform string, extraArgs interface{}) (status int, error string) {
	err := errors.New("REGISTER_NOT_SUPPORT:facebook account can't register by server, visit: https://www.facebook.com/")
	return 0, err.Error()
}

func NewFacebookMemberSystem() LiquidSDK.MemberSystem {
	fbMember := new(FacebookMemberSystem)
	return fbMember
}

func (fb *FacebookMemberSystem) Validate(
	fromId, fromToken, platform string,
	args interface{},
) (valid bool, msg string, overrideFromId string) {
	logger.SysLog.Infof("[Member|Validate|Facebook] FromId: %s, FromToken: %s", fromId, fromToken)
	valid = false
	msg = ""
	overrideFromId = ""
	return
}
