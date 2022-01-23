package mqttService

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iot-product/sensor-emulator-cli/payload"
)

func Publish(client mqtt.Client, topic string, payload *payload.Payload) bool {
	out, err := json.Marshal(payload)
	if err != nil {
		panic(err.Error())
	}
	token := client.Publish(topic, 0, false, string(out))
	return token.Wait()
}
