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
	Short: "Create a new topic skeleton",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flag("name").Value.String() == "" {
			fmt.Println("missing required argument: --name [topic name]")
			return
		}
		topic := templates.NewTopic(cmd,
			cmd.Flag("description").Value.String(),
			ProjectPath(),
		)
		err := createSubdirectories(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		createSlide(cmd, topic)
		if err != nil {
			fmt.Println(err)
			return
		}

		createManifest(cmd, topic)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().String("shortname", "", "Topic short name")
	createCmd.PersistentFlags().String("name", "", "Topic name")
	createCmd.PersistentFlags().String("subject", "Go", "Subject {Go,Kubernetes}")
	createCmd.PersistentFlags().String("level", "beginner", "{beginner,intermediate,advanced,expert}")
	createCmd.PersistentFlags().String("description", "Topic Description", "Topic description")

}

func createSlide(cmd *cobra.Command, topic templates.Topic) error {
	name := cmd.Flag("shortname").Value.String() + ".slide"

	return writeTemplateToFile(ProjectPath(), name, templates.Slide, topic)
}

func createManifest(cmd *cobra.Command, topic templates.Topic) error {
	name := cmd.Flag("shortname").Value.String() + ".json"

	js, err := json.Marshal(topic)
	if err != nil {
		return errors.Wrap(err, "encoding topic manifest")
	}
	return writeStringToFile(ProjectPath(), name, string(js))
}

func createSubdirectories(cmd *cobra.Command) error {
	path := filepath.Join(ProjectPath(), "src", cmd.Flag("shortname").Value.String())
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making topic src directory")
	}
	path = filepath.Join(ProjectPath(), "src", cmd.Flag("shortname").Value.String(), "exercises")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making topic exercises directory")
	}
	path = filepath.Join(ProjectPath(), "src", cmd.Flag("shortname").Value.String(), "solutions")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making topic solutions directory")
	}
	path = filepath.Join(ProjectPath(), "src", cmd.Flag("shortname").Value.String(), "demos")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making topic demos directory")
	}
	for _, dir := range subdirs {
		path := filepath.Join(ProjectPath(), cmd.Flag("shortname").Value.String(), dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return errors.Wrap(err, "making topic slide directories")
		}
	}
	return nil
}
