package settings

type AMQPConf struct {
	ProtocolVer string `envField:"amqp:ProtocolVer" default:"1"`
	Host        string `envField:"amqp:Host" default:"127.0.0.1"`
	Port        int    `envField:"amqp:Port" default:"5672"`
	Username    string `envField:"amqp:Username" default:""`
	Password    string `envField:"amqp:Password" default:""`
	Locale      string `envField:"amqp:Locale" default:""`
	FrameMax    int    `envField:"amqp:FrameMax" default:"10"`
	Heartbeat   int    `envField:"amqp:Heartbeat" default:"1000"`
	VisualHost  string `envField:"amqp:VisualHost" default:"/"`
}
