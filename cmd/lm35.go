/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"time"

	"github.com/iot-product/sensor-emulator-cli/payload"
	"github.com/spf13/cobra"
)

// lm35Cmd represents the lm35 command
var lm35Cmd = &cobra.Command{
	Use:   "lm35",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := generateTick(1)
		data := -55
		id := 0
		for range ticker.C {
			if data > 150 {
				data = -55
			}
			payload := buildPayload(id, data, time.RFC3339)
			fmt.Println(payload)
			data++
			id++
		}
	},
}

func buildPayload(id int, data int, timestamp string) *payload.Payload {
	return &payload.Payload{
		Id:        id,
		Data:      data,
		Timestamp: timestamp,
	}
}

func generateTick(interval int) *time.Ticker {
	return time.NewTicker(time.Duration(interval) * time.Second)
}

func init() {
	rootCmd.AddCommand(lm35Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lm35Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lm35Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
