package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gophertrain/trainctl/models"
	"github.com/gophertrain/trainctl/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func createSlide(cmd *cobra.Command, lesson models.Lesson) error {
	name := cmd.Flag("name").Value.String() + ".md"
	path := filepath.Join(ProjectPath(), cmd.Flag("module").Value.String())

	return writeTemplateToFile(path, name, templates.Slide, lesson)
}

func createManifest(cmd *cobra.Command, module models.Module) error {
	name := cmd.Flag("module").Value.String() + ".json"

	js, err := json.Marshal(module)
	if err != nil {
		return errors.Wrap(err, "encoding module manifest")
	}
	return writeStringToFile(ProjectPath(), name, string(js))
}

func createSubdirectories(cmd *cobra.Command) error {
	path := filepath.Join(ProjectPath(), "../", cmd.Flag("module").Value.String())
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module src directory")
	}
	path = filepath.Join(ProjectPath(), "../", cmd.Flag("module").Value.String(), "exercises")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module exercises directory")
	}
	path = filepath.Join(ProjectPath(), "../", cmd.Flag("module").Value.String(), "solutions")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module solutions directory")
	}
	path = filepath.Join(ProjectPath(), "../", cmd.Flag("module").Value.String(), "demos")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module demos directory")
	}
	/*	for _, dir := range subdirs {
			path := filepath.Join(ProjectPath(), cmd.Flag("name").Value.String(), dir)
			err := os.MkdirAll(path, 0755)
			if err != nil {
				return errors.Wrap(err, "making module slide directories")
			}
		}
	*/
	return nil
}
