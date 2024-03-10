package settings

type GinConf struct {
	RunMode      string `envField:"gin:RunMode" default:"debug"`
	HttpPort     int    `envField:"gin:HttpPort" default:"8080"`
	ReadTimeout  int64  `envField:"gin:ReadTimeout" default:"2000"`
	WriteTimeout int64  `envField:"gin.WriteTimeout" default:"15000"`
}
