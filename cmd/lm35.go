/*
Copyright © 2022 Ekky Kharismadhany <backendprogrammer43@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	mqttService "github.com/iot-product/sensor-emulator-cli/mqtt"
	"log"
	"strconv"
	"time"

	"github.com/iot-product/sensor-emulator-cli/payload"
	"github.com/spf13/cobra"
)

// lm35Cmd represents the lm35 command
var lm35 = &cobra.Command{
	Use:   "lm35",
	Short: "LM 35 Simulator",
	Long:  `LM 35 is a temperature sensor which has -55 to 150 degress celcius range of measurement`,
	Run: func(cmd *cobra.Command, args []string) {
		interval, err := cmd.Flags().GetString("interval")
		topic, err := cmd.Flags().GetString("topic")
		client := mqttInit()
		if err != nil {
			log.Fatal("Flags not found")
		}
		ticker := generateTick(interval)
		data := -55
		id := 0
		for range ticker.C {
			t := time.Now()
			if data > 150 {
				data = -55
			}
			payload := buildPayload(id, data, t.String())
			isPublished := mqttPublish(client, topic, payload)
			if isPublished {
				log.Printf("Data with id=%d is published", id)
			}
			data++
			id++
		}
	},
}

func mqttPublish(client mqtt.Client, topic string, payload *payload.Payload) bool {
	out, err := json.Marshal(payload)
	if err != nil {
		panic(err.Error())
	}
	token := client.Publish(topic, 0, false, string(out))
	return token.Wait()
}

func mqttInit() mqtt.Client {
	client := mqttService.Init()
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func buildPayload(id int, data int, timestamp string) *payload.Payload {
	return &payload.Payload{
		Id:        id,
		Data:      data,
		Timestamp: timestamp,
	}
}

func generateTick(interval string) *time.Ticker {
	duration, err := strconv.Atoi(interval)
	if err != nil && duration < 0 {
		log.Panic("Invalid interval value, please use integer positive value for interval")
	}
	return time.NewTicker(time.Duration(duration) * time.Second)
}

func init() {
	rootCmd.AddCommand(lm35)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lm35Cmd.PersistentFlags().String("foo", "", "A help for foo")
	lm35.PersistentFlags().String("interval", "1", "Set interval to generate")
	lm35.PersistentFlags().String("topic", "", "Set topic to subscribe")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lm35Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
