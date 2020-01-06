package LiquidEngine

import (
	"errors"
	"fmt"
	"github.com/cesnow/LiquidEngine/Options"
	"github.com/cesnow/LiquidEngine/Settings"
	"github.com/koding/multiconfig"
)

type IConfig interface {
}

type Config struct {
	App     *Settings.AppConf
	Gin     *Settings.GinConf
	AMQP    *Settings.AMQPConf
	CacheDB *Settings.CacheDbConf
	DocDB   *Settings.DocDbConf
	RDB     *Settings.RDBConf
	custom  map[string]interface{}
	raw     map[string]string
	engine  *Engine
}

var _ IConfig = &Config{}

func (config *Config) LoadExternalEnv(envPrefix string, conf interface{}, opts ...*Options.LoadEnvOptions) {
	envOpt := Options.MergeLoadEnvOptions(opts...)
	config.loadEnv(envPrefix, conf, envOpt)
	config.custom[envPrefix] = conf
}

func (config *Config) GetEnv(prefix string) (interface{}, error) {
	if val, ok := config.custom[prefix]; ok {
		return val, nil
	}
	fmt.Println(fmt.Errorf("[ConfigSystem] Config Not Found in Prefix `%s`, Please Check", prefix))
	return nil, errors.New("settings not found")
}

func (config *Config) systemExternalEnv(envPrefix string, conf interface{}, opts ...*Options.LoadEnvOptions) {
	envOpt := Options.MergeLoadEnvOptions(opts...)
	config.loadEnv(envPrefix, conf, envOpt)
}

func (config *Config) loadEnv(envPrefix string, conf interface{}, opts *Options.LoadEnvOptions) {
	InstantiateLoader := &multiconfig.EnvironmentLoader{
		Prefix:    envPrefix,
		CamelCase: *opts.CamelCase,
	}
	err := InstantiateLoader.Load(conf)
	if err != nil {
		panic(err)
	}
}
