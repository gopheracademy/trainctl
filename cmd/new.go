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

	"github.com/gophertrain/trainctl/models"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Module",
	Long:  `Use trainctl new module to create a new module.`,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flag("module").Value.String() == "" {
			fmt.Println("missing required argument: --name [module name]")
			return
		}
		module := models.NewModule(cmd,
			cmd.Flag("description").Value.String(),
			ProjectPath(),
		)
		err := createSubdirectories(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		createManifest(cmd, module)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	moduleCmd.AddCommand(newCmd)
}
