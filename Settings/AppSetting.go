package Settings

type AppConf struct {
	Codename  string `envField:"app:Codename" default:""`
	LogLevel string `envField:"app:LogLevel" default:"info"`
}
