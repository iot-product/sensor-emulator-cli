/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/iot-product/sensor-emulator-cli/helper"
	mqttService "github.com/iot-product/sensor-emulator-cli/mqtt"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// sr04Cmd represents the sr04 command
var sr04Cmd = &cobra.Command{
	Use:   "sr04",
	Short: "HC-SR04 is an Ultrasonic Sensor",
	Long: `Ultrasonic ranging module HC - SR04 provides 2cm - 400cm non-contact
measurement function, the ranging accuracy can reach to 3mm. The modules
includes ultrasonic transmitters, receiver and control circuit`,
	Run: func(cmd *cobra.Command, args []string) {
		interval, err := cmd.Flags().GetString("interval")
		topic, err := cmd.Flags().GetString("topic")
		client := mqttService.Init()
		ticker := helper.GenerateTick(interval)
		id := 0
		data := 2
		if err != nil {
			log.Fatalln("Flags Not Found")
		}
		for range ticker.C {
			t := time.Now()
			if data > 400 {
				data = 2
			}
			payload := helper.BuildPayload(id, data, t.String())
			if mqttService.Publish(client, topic, payload) {
				log.Printf("Data with id=%d is published", id)
			}
			id++
			data++
		}

	},
}

func init() {
	rootCmd.AddCommand(sr04Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sr04Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sr04Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sr04Cmd.PersistentFlags().String("interval", "1", "Set interval to generate")
	sr04Cmd.PersistentFlags().String("topic", "", "Set topic to subscribe")
}
