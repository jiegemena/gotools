package zmqtools

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"time"
)

func Publish(subscribe string, tcpbind string, f func() string) {
	//  Prepare our context and publisher socket
	publisher, _ := zmq.NewSocket(zmq.PUB)
	publisher.Bind(tcpbind)
	time.Sleep(time.Second)
	sequence := int64(1)
	for {
		msgs := f()
		_, err := publisher.SendMessage(fmt.Sprintf("%s", subscribe), msgs)
		if err != nil {
			break
		}
	}
	fmt.Printf("Interrupted\n%d messages out\n", sequence)
}

func Subscribe(subscribe string, tcpaddr string, f func(data string)) {

	subscriber, _ := zmq.NewSocket(zmq.SUB)
	subscriber.SetRcvhwm(100000)

	subscriber.SetSubscribe(subscribe)
	subscriber.Connect(tcpaddr)

	time.Sleep(time.Second)

	for {
		info, err := subscriber.RecvMessage(0)
		if err != nil {
			fmt.Println("sub recv err", err)
			break
		}
		for _, v := range info {
			f(v)
		}
	}
}
