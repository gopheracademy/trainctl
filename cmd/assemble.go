package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gophertrain/trainctl/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		err = assembleCourse(cmd, course)
		if err != nil {
			fmt.Println("Error assembling course", err)
			return
		}
	},
}

func assembleCourse(cmd *cobra.Command, course templates.Course) error {

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
		newModuleDir := filepath.Join(course.OutputDirectory, module.ShortName)
		err := os.Symlink(moduleDir, newModuleDir)
		if err != nil {
			return errors.Wrap(err, "symlink module directory")
		}

		// slide
		slide := filepath.Join(ProjectPath(), module.ShortName+".slide")
		newSlide := filepath.Join(course.OutputDirectory, module.NumberedPath(i+1)+".slide")
		err = os.Symlink(slide, newSlide)
		if err != nil {
			return errors.Wrap(err, "symlink slide")
		}
		// manifest

		// source code
		srcDir := filepath.Join(ProjectPath(), "src", module.ShortName)
		newsrcDir := filepath.Join(course.OutputDirectory, "src", module.ShortName)
		err = os.Symlink(srcDir, newsrcDir)
		if err != nil {
			return errors.Wrap(err, "symlink source directory")
		}

	}
	// only once
	err = writeStringToFile(course.OutputDirectory, "Vagrantfile", templates.Vagrantfile)
	if err != nil {
		return errors.Wrap(err, "Create Vagrantfile")
	}
	err = writeStringToFile(course.OutputDirectory, "bootstrap-vagrant.sh", templates.Bootstrap)
	if err != nil {
		return errors.Wrap(err, "Create bootstrap script")
	}
	err = os.Chmod(filepath.Join(course.OutputDirectory, "bootstrap-vagrant.sh"), 0755)
	if err != nil {
		return errors.Wrap(err, "make bootstrap script execuatable")
	}

	err = createCourseManifest(cmd, course)
	if err != nil {
		return errors.Wrap(err, "create course manifest")
	}
	return err
}

func init() {
	RootCmd.AddCommand(assembleCmd)

	assembleCmd.PersistentFlags().StringSlice("modules", []string{}, "List of modules to assemble, comma separated 'module1,module2'")
	assembleCmd.PersistentFlags().String("course", "", "Course Name e.g: 'Go for the Future'")
	assembleCmd.PersistentFlags().String("shortname", "", "Course Short Name e.g: 'goforfuture'")
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

	shortname, err := cmd.PersistentFlags().GetString("shortname")
	if err != nil {
		return errors.Wrap(err, "Check parameters: shortname")
	}
	if shortname == "" {
		return errors.New("Course shortname is required")
	}
	newModuleDir := filepath.Join(viper.GetString("coursedir"), shortname)

	b, err := dirExists(newModuleDir)
	if err != nil {
		return errors.Wrap(err, "Check output directory")
	}
	if b {
		return errors.New("Output directory exists. Cowardly refusing to overwrite.")
	}
	return nil
}

func createCourseManifest(cmd *cobra.Command, course templates.Course) error {
	name := cmd.Flag("shortname").Value.String() + ".json"

	js, err := json.Marshal(course)
	if err != nil {
		return errors.Wrap(err, "encoding course manifest")
	}
	return writeStringToFile(viper.GetString("coursedir"), name, string(js))
}
