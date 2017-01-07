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

// lessonCmd represents the lesson command
var lessonCmd = &cobra.Command{
	Use:   "create-lesson",
	Short: "Create a new Lesson for a Module.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("module").Value.String() == "" {
			fmt.Println("missing required argument: --module [module name]")
			return
		}
		if cmd.Flag("description").Value.String() == "" {
			fmt.Println("missing required argument: --description [lesson description]")
			return

		}
		if cmd.Flag("name").Value.String() == "" {
			fmt.Println("missing required argument: --name [lesson name]")
			return

		}
		module, err := getManifest(cmd.Flag("module").Value.String())
		if err != nil {
			fmt.Println("unable to find module", cmd.Flag("module").Value.String())
			return
		}

		lesson := models.NewLesson(cmd,
			cmd.Flag("name").Value.String(),
			cmd.Flag("description").Value.String(),
			ProjectPath(),
		)
		module.Lessons = append(module.Lessons, &lesson)
		err = createManifest(cmd, module)
		if err != nil {
			fmt.Println("unable to write manifest", err)
			return
		}

		err = createSlide(cmd, lesson)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	moduleCmd.AddCommand(lessonCmd)
	lessonCmd.Flags().String("name", "lesson", "Lesson Name")
	lessonCmd.Flags().String("topic", "Go", "Lesson Topic {Go,Kubernetes}")
	lessonCmd.Flags().String("level", "beginner", "{beginner,intermediate,advanced,expert}")
	lessonCmd.Flags().String("description", "Lesson Description", "Lesson Description")

}
