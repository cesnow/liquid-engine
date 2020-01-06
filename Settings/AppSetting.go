package Settings

type AppConf struct {
	Codename  string `envField:"app:Codename" default:""`
	JwtSecret string `envField:"app:JwtSecret" default:""`
}
