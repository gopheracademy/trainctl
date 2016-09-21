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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gophertrain/trainctl/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new module skeleton",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flag("name").Value.String() == "" {
			fmt.Println("missing required argument: --name [module name]")
			return
		}
		module := templates.NewModule(cmd,
			cmd.Flag("description").Value.String(),
			guessImportPath(),
		)
		err := createSubdirectories(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		createSlide(cmd, module)
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
	RootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().String("name", "", "Module name")
	createCmd.PersistentFlags().String("topic", "Go", "Module Topic {Go,Kubernetes}")
	createCmd.PersistentFlags().String("level", "beginner", "{beginner,intermediate,advanced,expert}")
	createCmd.PersistentFlags().String("description", "Module Description", "Module description")

}

func createSlide(cmd *cobra.Command, module templates.Module) error {
	name := cmd.Flag("name").Value.String() + ".slide"
	path := getPath(ProjectPath(), cmd.Flag("name").Value.String())

	return writeTemplateToFile(path, name, templates.Slide, module)
}

func createManifest(cmd *cobra.Command, module templates.Module) error {
	name := cmd.Flag("name").Value.String() + ".json"
	path := getPath(ProjectPath(), cmd.Flag("name").Value.String())

	js, err := json.Marshal(module)
	if err != nil {
		return errors.Wrap(err, "encoding module manifest")
	}
	return writeStringToFile(path, name, string(js))
}

func createSubdirectories(cmd *cobra.Command) error {
	path := filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String())
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module src directory")
	}
	path = filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String(), "exercises")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module src directory")
	}
	path = filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String(), "solutions")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module src directory")
	}
	for _, dir := range subdirs {
		path := getPath(cmd.Flag("name").Value.String(), dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return errors.Wrap(err, "making module slide directories")
		}
	}
	return nil
}
