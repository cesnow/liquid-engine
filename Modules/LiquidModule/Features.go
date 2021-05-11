package LiquidModule

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Modules/LiquidSDK"
	"reflect"
	"sync"
)

type TypeFeatures map[string]LiquidSDK.IGameSystem

var (
	Features     TypeFeatures
	featuresOnce sync.Once
)

func AddFeature(name string, module LiquidSDK.IGameSystem) {
	if _, find := GetModuleList()[name]; find {
		Logger.SysLog.Warnf(
			"[Engine] Pre-Define Game Feature (Name: `%s`, GameCmd: `%s`) Failed, Duplicated!",
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
