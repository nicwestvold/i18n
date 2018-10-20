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
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key [OPTIONS] <key> <value>",
	Short: "Adds a new key to the default (en-US) i18n file.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		forceIt := cmd.Flag("force").Value.String() == "true"

		p, err := real.New("en-US.json")
		if err != nil {
			log.Fatal(err)
		}
		data, err := p.ReadFile()
		if err != nil {
			log.Fatal(err)
		}

		// check if the key already exists in the i18n file
		_, ok := data[args[0]]
		// if the "--force" flag is set, then replace, otherwise do nothing
		if ok && !forceIt {
			fmt.Println("key already exists with value: \"" + data[args[0]].(string) + "\"")
		} else {
			if forceIt && ok {
				fmt.Printf("updating \"%s\" with value \"%s\" (previous: \"%s\")\n", args[0], args[1], data[args[0]].(string))
			} else {
				fmt.Printf("added key - %s: \"%s\"\n", args[0], args[1])
			}
			data[args[0]] = args[1]

			err = p.WriteFile(p.Filename, data)
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

	keyCmd.Flags().BoolP("force", "f", false, "Overwrite any existing value")
}
