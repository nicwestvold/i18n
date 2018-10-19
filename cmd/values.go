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
	"strings"
)

// valuesCmd represents the values command
var valuesCmd = &cobra.Command{
	Use:   "values",
	Short: "List current key/values",
	Run: func(cmd *cobra.Command, args []string) {
		var data map[string]interface{}

		// open file
		content, err := ioutil.ReadFile("en-US.json")
		if err != nil {
			log.Fatalf("%v", err)
		}

		// get JSON
		if err := json.Unmarshal(content, &data); err != nil {
			log.Fatalf("error unmarshalling JSON: %v", err)
		}

		// iterate over values, display
		maxLen := 0
		for key := range data {
			if len(key) > maxLen {
				maxLen = len(key)
			}
		}

		for key, val := range data {
			k := key + ": "
			if len(key) < maxLen {
				k = key + ": " + strings.Repeat(" ", maxLen-len(key))
			}
			fmt.Printf("%s \"%s\"\n", k, val)
		}
	},
}

func init() {
	listCmd.AddCommand(valuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// valuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// valuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
