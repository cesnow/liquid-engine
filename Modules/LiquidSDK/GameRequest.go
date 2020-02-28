package LiquidSDK

import (
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/mitchellh/mapstructure"
)

type IGameRequest interface {
	ToStruct(interface{}, string) error
	Raw() interface{}
	Get(key string) interface{}
}

type GameRequest struct {
	CmdData interface{}
}

func (gameRequest *GameRequest) ToStruct(target interface{}, tag string) error {
	config := &mapstructure.DecoderConfig{
		TagName: tag,
	}
	config.Result = &target
	decoder, _ := mapstructure.NewDecoder(config)
	decodeErr := decoder.Decode(gameRequest.CmdData)
	if decodeErr != nil {
		Logger.SysLog.Warnf("[Utils][ConvertStruct] Convert Failed, %s", decodeErr)
		return decodeErr
	}
	return nil
}

func (gameRequest *GameRequest) Raw() interface{} {
	return gameRequest.CmdData
}

func (gameRequest *GameRequest) Get(key string) interface{} {
	m := gameRequest.CmdData.(map[string]interface{})
	if r, ok := m[key]; ok {
		return r
	}
	return nil
}
