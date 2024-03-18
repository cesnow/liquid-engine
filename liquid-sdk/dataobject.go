package LiquidSDK

type CmdSignedBody struct {
	LiSign string `json:"LiSign"`
	LiData string `json:"LiData"`
}

type CmdRegister struct {
	FromType  string      `json:"from_type"`
	Account   string      `json:"account"`
	Password  string      `json:"password"`
	ExtraData interface{} `json:"extra_data" default:""`
}

type CmdAccount struct {
	FromType  string      `json:"from_type" default:""`
	FromId    string      `json:"from_id" default:""`
	FromToken string      `json:"from_token" default:""`
	Platform  string      `json:"platform" default:"main"`
	ExtraData interface{} `json:"extra_data" default:""`
}

type CmdAccountResponse struct {
	AutoId     *string `json:"auto_id" bson:"auto_id"`
	InviteCode *string `json:"invite_code" bson:"invite_code"`
}

type CmdBind struct {
	AutoId     string `json:"auto_id" bson:"auto_id"`
	InviteCode string `json:"invite_code" bson:"invite_code"`
	FromId     string `json:"from_id" bson:"from_id"`       //(str)third_party id or device id
	FromToken  string `json:"from_token" bson:"from_token"` // (str)third_party token
	FromType   string `json:"from_type" bson:"from_type"`   // (str)third_party name
}

type CmdAuth struct {
	AutoId     *string `json:"auto_id" bson:"auto_id"`
	InviteCode *string `json:"invite_code" bson:"invite_code"`
	Platform   *string `json:"platform" bson:"platform"`
}

type CmdAuthResponse struct {
	LiquidId    *string `json:"liquid_id" bson:"liquid_id"`
	LiquidToken *string `json:"liquid_token" bson:"liquid_token"`
}

type CmdCommand struct {
	LiquidId    *string     `json:"liquid_id" bson:"liquid_id"`
	LiquidToken *string     `json:"liquid_token" bson:"liquid_token"`
	Platform    *string     `json:"platform" bson:"platform"`
	CmdId       *string     `json:"cmd_id"`
	CmdSn       *string     `json:"cmd_sn"`
	CmdName     *string     `json:"cmd_name"`
	CmdData     interface{} `json:"cmd_data"`
}

type CmdCommandResponse struct {
	CmdData interface{} `json:"cmd_data"`
	CmdSn   *string     `json:"cmd_sn"`
}

type CmdErrorResponse struct {
	Error string `json:"error"`
}
