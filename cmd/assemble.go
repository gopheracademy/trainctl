package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	Short: "Assemble topicss into a course",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkParams(cmd); err != nil {
			fmt.Println(err)
			return
		}
		course := templates.NewCourse(cmd)
		topics, err := cmd.PersistentFlags().GetStringSlice("topics")
		if err != nil {
			fmt.Println("Error parsing topics.")
			return
		}
		var manifests []*templates.Topic
		for _, m := range topics {
			man, err := getManifest(m)
			if err != nil {
				fmt.Println("Error getting topic manifest.", err)
				return
			}
			manifests = append(manifests, &man)
		}
		course.Topics = manifests
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
	for i, topic := range course.Topics {
		topicDir := filepath.Join(ProjectPath(), topic.ShortName)
		newTopicDir := filepath.Join(course.OutputDirectory, topic.ShortName)
		err := os.Symlink(topicDir, newTopicDir)
		if err != nil {
			return errors.Wrap(err, "symlink topic directory")
		}

		// slide
		slide := filepath.Join(ProjectPath(), topic.ShortName+".slide")
		newSlide := filepath.Join(course.OutputDirectory, topic.NumberedPath(i+1)+".slide")
		err = os.Symlink(slide, newSlide)
		if err != nil {
			return errors.Wrap(err, "symlink slide")
		}
		// manifest

		// source code
		srcDir := filepath.Join(ProjectPath(), "src", topic.ShortName)
		newsrcDir := filepath.Join(course.OutputDirectory, "src", topic.ShortName)
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

	readme := filepath.Join(getSrcPath(), "github.com", "gophertrain", "trainctl", "templates", "readme.tmpl")
	rt, err := template.ParseFiles(readme)
	if err != nil {
		return errors.Wrap(err, "reading readme template")
	}

	rm, err := os.Create(filepath.Join(course.OutputDirectory, "README.md"))
	if err != nil {
		fmt.Println("create readme: ", err)
		return err
	}
	defer rm.Close()
	err = rt.Execute(rm, course)
	if err != nil {
		fmt.Print("execute course template: ", err)
		return err
	}

	envT := filepath.Join(getSrcPath(), "github.com", "gophertrain", "trainctl", "templates", "environment.tmpl")
	et, err := template.ParseFiles(envT)
	if err != nil {
		return errors.Wrap(err, "reading environment template")
	}

	envf, err := os.Create(filepath.Join(course.OutputDirectory, "environment.sh"))
	if err != nil {
		fmt.Println("create environment.sh: ", err)
		return err
	}
	defer envf.Close()
	err = et.Execute(envf, nil)
	if err != nil {
		fmt.Print("execute environment template: ", err)
		return err
	}
	err = os.Chmod(filepath.Join(course.OutputDirectory, "environment.sh"), 0755)
	if err != nil {
		fmt.Print("chmod environment template: ", err)
		return err
	}

	enrcT := filepath.Join(getSrcPath(), "github.com", "gophertrain", "trainctl", "templates", "envrc.tmpl")
	envrct, err := template.ParseFiles(enrcT)
	if err != nil {
		return errors.Wrap(err, "reading envrc template")
	}

	envrcf, err := os.Create(filepath.Join(course.OutputDirectory, ".envrc"))
	if err != nil {
		fmt.Println("create .envrc: ", err)
		return err
	}
	defer envrcf.Close()
	err = envrct.Execute(envrcf, nil)
	if err != nil {
		fmt.Print("execute envrc template: ", err)
		return err
	}
	return nil

}

func init() {
	RootCmd.AddCommand(assembleCmd)

	assembleCmd.PersistentFlags().StringSlice("topics", []string{}, "List of topics to assemble, comma separated 'topic1,topic2'")
	assembleCmd.PersistentFlags().String("course", "", "Course Name e.g: 'Go for the Future'")
	assembleCmd.PersistentFlags().String("shortname", "", "Course Short Name e.g: 'goforfuture'")
}

func checkParams(cmd *cobra.Command) error {
	topics, err := cmd.PersistentFlags().GetStringSlice("topics")
	if err != nil {
		return errors.Wrap(err, "Check parameters: topics")
	}
	if len(topics) < 1 {
		return errors.New("At least one topic is required")
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
	newTopicDir := filepath.Join(viper.GetString("coursedir"), shortname)

	b, err := dirExists(newTopicDir)
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
