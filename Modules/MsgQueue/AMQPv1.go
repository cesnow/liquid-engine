package MsgQueue

import (
	"context"
	"fmt"
	"github.com/cesnow/LiquidEngine/Logger"
	"github.com/cesnow/LiquidEngine/Settings"
	"os"
	"pack.ag/amqp"
	"time"
)

type AMQPv1 struct {
	config *Settings.AMQPConf
}

func connectToAMQP(conf *Settings.AMQPConf) (*amqp.Client, error) {
	client, err := amqp.Dial(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		amqp.ConnSASLPlain(conf.Username, conf.Password),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewConnectWithAMQPv1(conf *Settings.AMQPConf) (IAMQP, error) {
	client, err := connectToAMQP(conf)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	a := &AMQPv1{
		config: conf,
	}
	return a, nil
}

// --------------------------  IAMQP Implement --------------------------

func (a1 *AMQPv1) Name() string {
	return "amqp 1.0"
}

func (a1 *AMQPv1) GetProtocolVersion() int8 {
	return 1
}

func (a1 *AMQPv1) SendMessage(queue, msg string) {

	client, err := connectToAMQP(a1.config)
	if err != nil {
		Logger.SysLog.Warnf("[MsgQueue][AMQPv1] Dial New Connection Failed, %s", err)
		return
	}

	session, err := client.NewSession()
	if err != nil {
		Logger.SysLog.Warnf("[MsgQueue][AMQPv1] Creating AMQPv1 session Failed:, %s", err)
		return
	}

	ctx := context.Background()
	{
		sender, err := session.NewSender(
			amqp.LinkTargetAddress(queue),
		)
		if err != nil {
			Logger.SysLog.Warnf("[MsgQueue][AMQPv1] Creating AMQPv1 SenderLink Failed:, %s", err)
		}
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		err = sender.Send(ctx, amqp.NewMessage([]byte(msg)))

		if err != nil {
			Logger.SysLog.Warnf("[MsgQueue][AMQPv1] Sending Message Failed:, %s", err)
		}
		sender.Close(ctx)
		session.Close(ctx)
		cancel()
	}
}

func (a1 *AMQPv1) ReceiveMessage(queue string, callback func(interface{})) {
	client, err := connectToAMQP(a1.config)
	if err != nil {
		Logger.SysLog.Warnf("[MsgQueue][AMQPv1] Dial New Connection Failed, %s", err)
		return
	}
	session, err := client.NewSession()
	if err != nil {
		Logger.SysLog.Warnf("[MsgQueue][AMQPv1] Creating AMQPv1 session Failed:, %s", err)
		return
	}

	receiver, err := session.NewReceiver(
		amqp.LinkSourceAddress(queue),
		amqp.LinkCredit(10),
	)
	if err != nil {
		Logger.SysLog.Warnf("[MsgQueue][AMQPv1] Creating AMQPv1 ReceiverLink Failed:, %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
		_ = receiver.Close(ctx)
		cancel()
	}()
	stopChan := make(chan bool)
	go func() {
		Logger.SysLog.Debugf("[MsgQueue][AMQPv1] Consumer Ready, PID: %d", os.Getpid())
		for {
			msg, _ := receiver.Receive(context.TODO())
			_ = msg.Accept()
			callback(msg)
		}
	}()
	<-stopChan
	session.Close(context.Background())
}
