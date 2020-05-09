package LiquidEngine

import (
	"github.com/cesnow/LiquidEngine/Features/Game"
	"github.com/cesnow/LiquidEngine/Features/Member"
	"testing"
)

const CmdEcho = "echo"
const MemberLine = "line"
const MemberFacebook = "facebook"

func TestEngine_Serve(t *testing.T) {
	engine := New()
	engine.RegisterGame(CmdEcho, Game.NewEchoSystem())
	engine.RegisterMember(MemberFacebook, Member.NewFacebookMemberSystem())
	engine.RegisterMember(MemberLine, Member.NewLineMemberSystem())
	engine.Serve()
}
