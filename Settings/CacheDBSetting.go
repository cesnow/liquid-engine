package Settings

type CacheDbConf struct {
	Host        string `envField:"cachedb:Host" default:"127.0.0.1"`
	Port        int    `envField:"cachedb:Port" default:"6379"`
	Password    string `envField:"cachedb:Password" default:""`
	MaxIdle     int    `envField:"cachedb:MaxIdle" default:"100"`
	MaxActive   int    `envField:"cachedb:MaxActive" default:"4000"`
	IdleTimeout int    `envField:"cachedb:IdleTimeout" default:"180"`
	Wait        bool   `envField:"cachedb:Wait" default:"true"`
	Database    int    `envField:"cachedb:Database" default:"0"`
}
