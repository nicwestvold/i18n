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
	"github.com/spf13/cobra"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize i18n files for a directory",
	Long:  `A "en-US.json" file is created in the root directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		str := []byte("{\n}\n")

		if _, err := os.Stat("en-US.json"); !os.IsNotExist(err) {
			fmt.Println("file exists. skipping init")
		} else {
			f, err := os.Create("en-US.json")
			check(err)

			_, err = f.Write(str)
			check(err)

			defer f.Close()
			fmt.Println("i18n initialized")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
