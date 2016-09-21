// Copyright © 2016 Brian Ketelsen <me@brianketelsen.com>
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
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// removeCmd represents the create command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a new module skeleton",
	Long:  `Removes the module directory and all of its contents, as well as the directory`,
	Run: func(cmd *cobra.Command, args []string) {
		really, _ := cmd.PersistentFlags().GetBool("really")
		if really {
			err := removeSubdirectories(cmd)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
		fmt.Printf("Cowardly refusing to remove module %s without --really flag", cmd.Flag("name").Value.String())
		return
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	removeCmd.PersistentFlags().String("name", "", "Module name")

	removeCmd.PersistentFlags().Bool("really", false, "Really destroy this module? It's permanent!")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func removeSubdirectories(cmd *cobra.Command) error {

	path := filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String())
	err := os.RemoveAll(path)
	if err != nil {
		return errors.Wrap(err, "removing module src directory")
	}
	name := cmd.Flag("name").Value.String()
	path = getPath(name, "")
	err = os.RemoveAll(path)
	if err != nil {
		return errors.Wrap(err, "removing module directory")
	}

	return nil

}
