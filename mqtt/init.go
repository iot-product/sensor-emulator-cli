package mqttService

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

var broker = "broker.hivemq.com"
var port = 1883

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Printf("Connected to %s with port %d\n", broker, port)
}

var disconnectHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	if err != nil {
		log.Panicln(err.Error())
	}
	fmt.Println("Connection lost")
}

func Init() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = disconnectHandler
	return mqtt.NewClient(opts)
}
