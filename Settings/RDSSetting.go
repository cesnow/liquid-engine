package Settings

type RDBConf struct {
	Dialects    string   `envField:"rds:Dialects"`
	Host        string   `envField:"rds:Host"`
	User        string   `envField:"rds:User"`
	Pass        string   `envField:"rds:Pass"`
	DbNames     []string `envField:"rds:DbNames"`
	TablePrefix string   `envField:"rds:TablePrefix"`
	MaxIdleConn int      `envFields:"rds:MaxIdleConn"`
	MaxOpenConn int      `envFields:"rds:MaxOpenConn"`
}
