package LiquidEngine

import (
	"github.com/cesnow/liquid-engine/features/member"
	"github.com/cesnow/liquid-engine/features/system"
	"testing"
)

const MemberLine = "line"
const MemberFacebook = "facebook"

func Test_EngineServe(t *testing.T) {
	engine := New()

	engine.RegisterMember(MemberFacebook, member.NewFacebookMemberSystem())
	engine.RegisterMember(MemberLine, member.NewLineMemberSystem())

	engine.RegisterFeature("echo", system.NewEchoSystem())

	//WorkQueue.Sample()
	engine.Serve()
}
