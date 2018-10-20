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
	"fmt"
	"github.com/nicwestvold/i18n/parser/real"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// valuesCmd represents the values command
var valuesCmd = &cobra.Command{
	Use:   "values",
	Short: "List current key/values",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := real.New("en-US.json")
		if err != nil {
			log.Fatal(err)
		}
		data, err := p.ReadFile()
		if err != nil {
			log.Fatal(err)
		}

		// iterate over values to get the max key length
		maxLen := 0
		for key := range data {
			if len(key) > maxLen {
				maxLen = len(key)
			}
		}

		// iterate over the values to display all key/values
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
