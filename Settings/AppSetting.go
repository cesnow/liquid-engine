package Settings

type AppConf struct {
	Codename       string `envField:"app:Codename" default:""`
	LogLevel       string `envField:"app:LogLevel" default:"info"`
	RpcCommandMode bool   `envField:"app:RpcCommandMode" default:"false"`
	RpcRemoteIp    string `envField:"app:RpcRemoteIp" default:"0.0.0.0:9999"`
}
