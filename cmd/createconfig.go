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

	"github.com/gophertrain/trainctl/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createconfigCmd represents the createconfig command
var createconfigCmd = &cobra.Command{
	Use:   "createconfig",
	Short: "Create a config file in your home directory",
	Long:  `Creates a .trainctl.yaml file in your home directory, which you can edit to provide default values for the create commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := createTheConfig(cmd)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(createconfigCmd)
	createconfigCmd.Flags().BoolP("force", "f", false, "Overwrite existing")
}

func createTheConfig(cmd *cobra.Command) error {
	name := ".trainctl.yaml"
	path, err := homeDir()
	if err != nil {
		return errors.Wrap(err, "get home directory")
	}
	e, err := exists(filepath.Join(path, name))
	if err != nil {
		return errors.Wrap(err, "check existing config")
	}
	f, err := cmd.Flags().GetBool("force")
	if err != nil {
		return errors.Wrap(err, "getting force flag")
	}
	if e && !f {
		return errors.New("File exists, use --force")
	}
	if e {
		err := os.Remove(filepath.Join(path, name))

		if err != nil {
			return errors.Wrap(err, "removing existing configuration file")
		}
	}
	conf := templates.Conf{
		Author:  viper.GetString("author"),
		Twitter: viper.GetString("twitter"),
		Email:   viper.GetString("email"),
	}
	return writeTemplateToFile(path, name, templates.Config, conf)
}
