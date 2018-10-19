// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Adds a new key to the default (en-US) i18n file.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var data map[string]interface{}

		// read the contents of the default file
		content, err := ioutil.ReadFile("en-US.json")
		if err != nil {
			log.Fatalf("error reading file: %v", err)
		}

		// unmarshal the JSON
		if err := json.Unmarshal(content, &data); err != nil {
			log.Fatalf("error unmarshalling JSON: %v", err)
		}

		// check if the key already exists in the i18n file
		if val, ok := data[args[0]]; ok {
			fmt.Println("key already exists with value: \"" + val.(string) + "\"")
		} else {
			fmt.Printf("added key - %s: \"%s\"\n", args[0], args[1])
			data[args[0]] = args[1]
			jsonStr, err := json.MarshalIndent(&data, "", "  ")
			if err != nil {
				log.Fatalf("error marshalling JSON: %v", err)
			}
			// write the new JSON to the file
			err = ioutil.WriteFile("en-US.json", jsonStr, os.ModeAppend)
			if err != nil {
				log.Fatalf("error writing JSON file: %v", err)
			}
		}
	},
}

func init() {
	addCmd.AddCommand(keyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
