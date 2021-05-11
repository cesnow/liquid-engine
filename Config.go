package LiquidEngine

import (
	"errors"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Options"
	"github.com/cesnow/LiquidEngine/Settings"
	"github.com/koding/multiconfig"
	"github.com/xxtea/xxtea-go/xxtea"
	"os"
	"reflect"
	"strings"
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
	v := reflect.ValueOf(conf)
	if v.IsValid() == false {
		panic("not valid")
	}
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	tv := v
	if tv.Kind() == reflect.Ptr && tv.CanSet() {
		tv.Set(reflect.New(tv.Type().Elem()))
		tvi := tv.Interface()
		elem := tv.Elem()
		if elem.Kind() != reflect.Struct {
			panic("[Config] not struct")
		}

		config.loadEnv(envPrefix, tvi, envOpt)
		config.custom[envPrefix] = tvi
		for i := 0; i < elem.NumField(); i++ {
			valueField := elem.Field(i)
			typeField := elem.Type().Field(i)
			tag := typeField.Tag
			decodedTag := tag.Get("decode")
			decodedKey := tag.Get("key")
			fieldRequired := tag.Get("required")
			if fieldRequired == "true" && valueField.String() == "" {
				Logger.SysLog.Errorf(
					"[Config] Please check field required `%s -> %s`",
					envPrefix,
					elem.Type().Field(i).Name,
				)
				os.Exit(97)
			}

			if decodedKey != "" && decodedTag != "" {
				newDecodedValue, deErr := xxtea.DecryptString(valueField.String(), decodedKey)
				if deErr == nil {
					if decodedTag == "pem" {
						newDecodedValue = strings.Replace(newDecodedValue, `\n`, "\n", -1)
					}
					valueField.SetString(newDecodedValue)
				} else {
					Logger.SysLog.Warnf("[ConfigConvertFailed] %s -> %s -> %+v", typeField, valueField, deErr)
				}
			}
		}
	}
}

func (config *Config) GetEnv(prefix string) (interface{}, error) {
	if val, ok := config.custom[prefix]; ok {
		return val, nil
	}
	Logger.SysLog.Errorf("[ConfigSystem] Config Not Found in Prefix `%s`, Please Check", prefix)
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
