package helper

import (
	"github.com/iot-product/sensor-emulator-cli/payload"
	"log"
	"strconv"
	"time"
)

func BuildPayload(id int, data int, timestamp string) *payload.Payload {
	return &payload.Payload{
		Id:        id,
		Data:      data,
		Timestamp: timestamp,
	}
}

func GenerateTick(interval string) *time.Ticker {
	duration, err := strconv.Atoi(interval)
	if err != nil && duration < 0 {
		log.Panic("Invalid interval value, please use integer positive value for interval")
	}
	return time.NewTicker(time.Duration(duration) * time.Second)
}
