package LiquidEngine

import (
	"github.com/cesnow/liquid-engine/features/member"
	"testing"
)

const MemberLine = "line"
const MemberFacebook = "facebook"

func Test_EngineServe(t *testing.T) {
	engine := New()
	engine.RegisterMember(MemberFacebook, member.NewFacebookMemberSystem())
	engine.RegisterMember(MemberLine, member.NewLineMemberSystem())
	//WorkQueue.Sample()
	engine.Serve()
}
