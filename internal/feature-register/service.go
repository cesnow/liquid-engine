package feature_register

import (
	LiquidSDK "github.com/cesnow/liquid-engine/liquid-sdk"
	"github.com/cesnow/liquid-engine/logger"
	"reflect"
	"sync"
)

type TypeFeatures map[string]LiquidSDK.CommandSystem

var (
	Features     TypeFeatures
	featuresOnce sync.Once
)

func AddFeature(name string, module LiquidSDK.CommandSystem) {
	if _, find := GetModuleList()[name]; find {
		logger.SysLog.Warnf(
			"[Engine] Pre-Define Feature (Name: `%s`, Cmd: `%s`) Failed, Duplicated!",
			reflect.TypeOf(module),
			name,
		)
		return
	}
	GetModuleList()[name] = module
}

func GetModuleList() TypeFeatures {
	featuresOnce.Do(func() {
		Features = make(TypeFeatures)
	})
	return Features
}
