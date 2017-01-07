// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
)

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Work with Modules",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("subcommand required")
	},
}

func init() {
	RootCmd.AddCommand(moduleCmd)
	moduleCmd.PersistentFlags().String("module", "", "Module name")
	moduleCmd.PersistentFlags().String("topic", "Go", "Module Topic {Go,Kubernetes}")
	moduleCmd.PersistentFlags().String("level", "beginner", "{beginner,intermediate,advanced,expert}")
	moduleCmd.PersistentFlags().String("description", "Module Description", "Module description")

}
