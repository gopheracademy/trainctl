package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
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

func createReadme(cmd *cobra.Command, path string) error {
	readme := filepath.Join(getSrcPath(), "github.com", "gophertrain", "trainctl", "templates", "info.tmpl")
	rt, err := template.ParseFiles(readme)
	if err != nil {
		return errors.Wrap(err, "reading readme template")
	}

	rm, err := os.Create(filepath.Join(path, "README.txt"))
	if err != nil {
		fmt.Println("create readme: ", err)
		return err
	}
	defer rm.Close()
	err = rt.Execute(rm, nil)
	if err != nil {
		fmt.Print("execute course template: ", err)
		return err
	}
	return nil
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
	err = createReadme(cmd, path)
	if err != nil {
		return errors.Wrap(err, "making module exercises directory readme")
	}
	path = filepath.Join(ProjectPath(), "../", cmd.Flag("module").Value.String(), "solutions")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module solutions directory")
	}
	err = createReadme(cmd, path)
	if err != nil {
		return errors.Wrap(err, "making module solutions directory readme")
	}
	path = filepath.Join(ProjectPath(), "../", cmd.Flag("module").Value.String(), "demos")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module demos directory")
	}
	err = createReadme(cmd, path)
	if err != nil {
		return errors.Wrap(err, "making module demos directory readme")
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
