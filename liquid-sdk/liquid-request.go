package LiquidSDK

import (
	"github.com/cesnow/liquid-engine/logger"
	"github.com/mitchellh/mapstructure"
)

type CommandRequest interface {
	ToStruct(interface{}, string) error
	Raw() interface{}
	Get(key string) interface{}
	GetLiquidId() string
	GetPlatform() string
	GetCmdId() string
	GetCmdSn() string
	GetCmdName() string
}

type LiquidRequest struct {
	LiquidId *string
	Platform *string
	CmdId    *string
	CmdSn    *string
	CmdName  *string
	CmdData  interface{}
}

func (request *LiquidRequest) ToStruct(target interface{}, tag string) error {
	config := &mapstructure.DecoderConfig{
		TagName: tag,
	}
	config.Result = &target
	decoder, _ := mapstructure.NewDecoder(config)
	decodeErr := decoder.Decode(request.CmdData)
	if decodeErr != nil {
		logger.SysLog.Warnf("[Utils][ConvertStruct] Convert Failed, %s", decodeErr)
		return decodeErr
	}
	return nil
}

func (request *LiquidRequest) Raw() interface{} {
	return request.CmdData
}

func (request *LiquidRequest) Get(key string) interface{} {
	m := request.CmdData.(map[string]interface{})
	if r, ok := m[key]; ok {
		return r
	}
	return nil
}

func (request *LiquidRequest) GetLiquidId() string {
	if request.LiquidId == nil {
		return ""
	}
	return *request.LiquidId
}

func (request *LiquidRequest) GetPlatform() string {
	if request.Platform == nil {
		return ""
	}
	return *request.Platform
}

func (request *LiquidRequest) GetCmdId() string {
	if request.CmdId == nil {
		return ""
	}
	return *request.CmdId
}

func (request *LiquidRequest) GetCmdSn() string {
	if request.CmdSn == nil {
		return ""
	}
	return *request.CmdSn
}

func (request *LiquidRequest) GetCmdName() string {
	if request.CmdName == nil {
		return ""
	}
	return *request.CmdName
}
