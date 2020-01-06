package Settings

type CacheDbConf struct {
	Host        string `envField:"cachedb:Host" default:"127.0.0.1"`
	Port        int    `envField:"cachedb:Port" default:"6379"`
	Password    string `envField:"cachedb:Password" default:""`
	MaxIdle     int    `envField:"cachedb:MaxIdle" default:"30"`
	MaxActive   int    `envField:"cachedb:MaxActive" default:"20"`
	IdleTimeout int    `envField:"cachedb:IdleTimeout" default:"200"`
	Wait        bool   `envField:"cachedb:Wait" default:"true"`
}
