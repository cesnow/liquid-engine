package settings

type AppConf struct {
	Codename       string `envField:"app:Codename" default:""`
	LogLevel       string `envField:"app:LogLevel" default:"info"`
	RpcCommandMode bool   `envField:"app:RpcCommandMode" default:"false"`
	RpcBindPort    int    `envField:"app:RpcBindPort" default:"9999"`
	RpcEndpoint    string `envField:"app:RpcEndpoint" default:"0.0.0.0"`
}
