package MsgQueue

import (
	"errors"
	"fmt"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Settings"
	"log"
	"strings"
)

type IAMQP interface {
	Name() string
	SendMessage(string, string)
	ReceiveMessage(string)
	GetProtocolVersion() int8
}

type AMQPFactory func(conf *Settings.AMQPConf) (IAMQP, error)

var amqpFactories = make(map[string]AMQPFactory)

func Register(name string, factory AMQPFactory) {
	if factory == nil {
		log.Panicf("AMQP factory %s does not exist.", name)
	}
	_, registered := amqpFactories[name]
	if registered {
		fmt.Printf("AMQP factory %s already registered. Ignoring.", name)
	}
	amqpFactories[name] = factory
}

func init() {
	Register("0", NewConnectWithAMQPv0)
	Register("1", NewConnectWithAMQPv1)
}

func Connect(conf *Settings.AMQPConf) (IAMQP, error) {
	Logger.SysLog.Infof("[MsgQueue] Connecting to Message Queue Service (AMQPv%s)", conf.ProtocolVer)
	engineName := conf.ProtocolVer
	engineFactory, ok := amqpFactories[engineName]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available amqp factories for logging.
		availableAMQPs := make([]string, len(amqpFactories))
		for k, _ := range amqpFactories {
			availableAMQPs = append(availableAMQPs, k)
		}
		return nil, errors.New(fmt.Sprintf("Invalid AMQP name. Must be one of: %s", strings.Join(availableAMQPs, ", ")))
	}
	// Run the factory with the configuration.
	engine, err := engineFactory(conf)
	if err != nil {
		return nil, err
	}
	Logger.SysLog.Info("[MsgQueue] Connected to the Message Queue Service")
	return engine, nil
}
