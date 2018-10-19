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
	"log"
	"os"

	"github.com/nicwestvold/i18n/parser"
	"github.com/nicwestvold/i18n/parser/real"
	"github.com/spf13/cobra"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Adds a new key to the default (en-US) i18n file.",
	// Args:  cobra.ExactArgs(2),
	Run: run,
}

func run(cmd *cobra.Command, args []string) {

	var p parser.Parser
	var err error
	switch os.Getenv("WHICH_PARSER") {
	case "real":
		p, err = real.New("translations.json", "output.json")
	default:
		// p, err = fake.New()
	}

	if err != nil {
		log.Fatalf("failed to create parser: %v", err)
	}

	data, err := p.ReadFile()
	if err != nil {
		log.Fatalf("failed to read data from file: %v", err)
	}

	out, err := p.Convert(data)
	if err != nil {
		log.Fatalf("failed to convert intput: %v", err)
	}

	// check if the key already exists in the i18n file
	if err := p.WriteFile(out); err != nil {
		log.Fatalf("failed to write JSON file: %v", err)
	}
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
