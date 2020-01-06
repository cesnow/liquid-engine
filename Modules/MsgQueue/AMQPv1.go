package MsgQueue

import (
	"context"
	"fmt"
	"github.com/cesnow/LiquidEngine/Settings"
	"os"
	"pack.ag/amqp"
	"time"
)

type AMQPv1 struct {
	client  *amqp.Client
	session *amqp.Session
}

func NewConnectWithAMQPv1(conf *Settings.AMQPConf) (IAMQP, error) {

	client, err := amqp.Dial(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		amqp.ConnSASLPlain(conf.Username, conf.Password),
	)

	if err != nil {
		return nil, err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	return &AMQPv1{
		client:  client,
		session: session,
	}, nil
}

// --------------------------  IAMQP Implement --------------------------

func (a1 *AMQPv1) GetClient() *amqp.Client {
	return a1.client
}

func (a1 *AMQPv1) GetSession() *amqp.Session {
	return a1.session
}

func (a1 *AMQPv1) Name() string {
	return "amqp 1.0"
}

func (a1 *AMQPv1) GetProtocolVersion() int8 {
	return 1
}

func (a1 *AMQPv1) SendMessage(queue, msg string) {
	sender, err := a1.session.NewSender(
		amqp.LinkTargetAddress(queue),
	)
	if err != nil {
		fmt.Println("Creating sender link:", err)
	}

	err = sender.Send(context.Background(), amqp.NewMessage([]byte(msg)))
	if err != nil {
		fmt.Println("Sending message:", err)
	}
	sender.Close(context.Background())
}

func (a1 *AMQPv1) ReceiveMessage(queue string) {

	session, err := a1.client.NewSession()
	if err != nil {
		fmt.Println("Creating AMQP session:", err)
	}

	// Create a receiver
	receiver, err := session.NewReceiver(
		amqp.LinkSourceAddress(queue),
		amqp.LinkCredit(10),
	)
	if err != nil {
		fmt.Println("Creating receiver link:", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
		receiver.Close(ctx)
		cancel()
	}()

	stopChan := make(chan bool)
	go func() {
		fmt.Printf("consumer ready, pid: %d", os.Getpid())
		for {
			msg, _ := receiver.Receive(context.TODO())
			_ = msg.Accept()
			xx(msg.GetData())
		}
	}()
	<-stopChan

}

func xx(x []byte) {
	fmt.Printf("Message received: %s\n", x)
}
