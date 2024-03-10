package MsgQueue

import "github.com/cesnow/liquid-engine/settings"

type AMQPv0 struct {
}

func NewConnectWithAMQPv0(conf *settings.AMQPConf) (IAMQP, error) {
	return &AMQPv0{}, nil
}

// --------------------------  IAMQP Implement --------------------------

func (a0 *AMQPv0) Name() string {
	return "amqp >=0.9.x, <1.0"
}

func (a0 *AMQPv0) SendMessage(queue, msg string) {

}

func (a0 *AMQPv0) ReceiveMessage(queue string, f func(interface{})) {
}
func (a0 *AMQPv0) GetProtocolVersion() int8 {
	return 0
}
