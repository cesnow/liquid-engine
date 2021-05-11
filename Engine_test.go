package LiquidEngine

import (
	_ "github.com/cesnow/LiquidEngine/Features/Game"
	"github.com/cesnow/LiquidEngine/Features/Member"
	"testing"
)

const MemberLine = "line"
const MemberFacebook = "facebook"

func TestEngine_Serve(t *testing.T) {
	engine := New()
	// engine.RegisterGame("echo", Game.NewEchoSystem())
	engine.RegisterMember(MemberFacebook, Member.NewFacebookMemberSystem())
	engine.RegisterMember(MemberLine, Member.NewLineMemberSystem())
	engine.Serve()
}
