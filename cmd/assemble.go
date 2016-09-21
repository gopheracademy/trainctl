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
)

// assembleCmd represents the assemble command
var assembleCmd = &cobra.Command{
	Use:   "assemble",
	Short: "Assemble modules into a course",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkParams(cmd); err != nil {
			fmt.Println(err)
			return
		}
		course := templates.NewCourse(cmd)
		modules, err := cmd.PersistentFlags().GetStringSlice("modules")
		if err != nil {
			fmt.Println("Error parsing modules.")
			return
		}
		var manifests []*templates.Module
		for _, m := range modules {
			man, err := getManifest(m)
			if err != nil {
				fmt.Println("Error getting module manifest.", err)
				return
			}
			manifests = append(manifests, &man)
		}
		course.Modules = manifests
		err = assembleCourse(course)
		if err != nil {
			fmt.Println("Error assembling course", err)
			return
		}
	},
}

func assembleCourse(course templates.Course) error {

	err := os.MkdirAll(course.OutputDirectory, 0755)
	if err != nil {
		return errors.Wrap(err, "create output directory")
	}
	for _, dir := range outputdirs {
		path := filepath.Join(course.OutputDirectory, dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return errors.Wrap(err, "making output directories")
		}
	}
	for i, module := range course.Modules {
		moduleDir := filepath.Join(ProjectPath(), module.ShortName)
		newModuleDir := filepath.Join(course.OutputDirectory, "slides", module.NumberedPath(i+1))
		err := os.Symlink(moduleDir, newModuleDir)
		if err != nil {
			return errors.Wrap(err, "symlink module directory")
		}

		srcDir := filepath.Join(ProjectPath(), "src", module.ShortName)
		newsrcDir := filepath.Join(course.OutputDirectory, "src", module.ShortName)
		err = os.Symlink(srcDir, newsrcDir)
		if err != nil {
			return errors.Wrap(err, "symlink source directory")
		}
	}

	return err
}

func init() {
	RootCmd.AddCommand(assembleCmd)

	assembleCmd.PersistentFlags().StringSliceP("modules", "m", []string{}, "List of modules to assemble, comma separated 'module1,module2'")
	assembleCmd.PersistentFlags().StringP("course", "c", "", "Course Name e.g: 'Go for the Future'")
	assembleCmd.PersistentFlags().StringP("output", "o", "", "Output Directory e.g. /Users/you/goforthefuture")

}

func checkParams(cmd *cobra.Command) error {
	modules, err := cmd.PersistentFlags().GetStringSlice("modules")
	if err != nil {
		return errors.Wrap(err, "Check parameters: modules")
	}
	if len(modules) < 1 {
		return errors.New("At least one module is required")
	}

	course, err := cmd.PersistentFlags().GetString("course")
	if err != nil {
		return errors.Wrap(err, "Check parameters: course")
	}
	if course == "" {
		return errors.New("Course name is required")
	}
	output, err := cmd.PersistentFlags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "Check parameters: output")
	}
	if output == "" {
		return errors.New("Output Directory is required")
	}
	b, err := dirExists(output)
	if err != nil {
		return errors.Wrap(err, "Check output directory")
	}
	if b {
		return errors.New("Output directory exists. Cowardly refusing to overwrite.")
	}
	return nil
}
