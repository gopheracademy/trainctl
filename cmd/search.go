package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gophertrain/trainctl/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search topics by metadata",
	Long:  `Search performs a search of the repository's topics using a logical OR of all the supplied search conditions.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := checkSearchParams(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		mods, err := search(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, m := range mods {
			fmt.Println(m)
		}
		return
	},
}

func checkSearchParams(cmd *cobra.Command) error {

	shortname, err := cmd.PersistentFlags().GetString("shortname")
	if err != nil {
		return errors.Wrap(err, "Check parameters: shortname")
	}
	topic, err := cmd.PersistentFlags().GetString("topic")
	if err != nil {
		return errors.Wrap(err, "Check parameters: topic")
	}

	level, err := cmd.PersistentFlags().GetString("level")
	if err != nil {
		return errors.Wrap(err, "Check parameters: level")
	}

	description, err := cmd.PersistentFlags().GetString("description")
	if err != nil {
		return errors.Wrap(err, "Check parameters: description")
	}
	var found bool
	if shortname != "" {
		found = true
	}
	if topic != "" {
		found = true
	}

	if description != "" {
		found = true
	}
	if level != "" {
		found = true
	}
	if !found {

		return errors.New("At least one search parameter required: shortname, topic, level")
	}
	return nil
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().String("shortname", "", "Topic shortname")
	searchCmd.PersistentFlags().String("subject", "", "Topic subject {Go,Kubernetes}")
	searchCmd.PersistentFlags().String("level", "", "{beginner,intermediate,advanced,expert}")
	searchCmd.PersistentFlags().String("description", "", "Topic description")
}

func search(cmd *cobra.Command) ([]templates.Topic, error) {
	var results []templates.Topic

	dir, err := os.Open(ProjectPath())
	if err != nil {
		return results, errors.Wrap(err, "open project directory")
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		return results, errors.Wrap(err, "list project directory")
	}

	shortname, err := cmd.PersistentFlags().GetString("shortname")
	if err != nil {
		return results, errors.Wrap(err, "Check parameters: shortname")
	}
	subject, err := cmd.PersistentFlags().GetString("subject")
	if err != nil {
		return results, errors.Wrap(err, "Check parameters: topic")
	}

	description, err := cmd.PersistentFlags().GetString("description")
	if err != nil {
		return results, errors.Wrap(err, "Check parameters: description")
	}
	level, err := cmd.PersistentFlags().GetString("level")
	if err != nil {
		return results, errors.Wrap(err, "Check parameters: level")
	}
	for _, f := range files {
		if f.IsDir() {
			var hit bool
			manifestPath := filepath.Join(ProjectPath(), f.Name()+".json")
			_, err := os.Stat(manifestPath)
			if err != nil {
				continue
			}

			topic, err := getManifest(f.Name())
			if err != nil {
				return results, errors.Wrap(err, "get topic")
			}
			if shortname != "" {
				if strings.Contains(topic.ShortName, shortname) {
					hit = true
				}
			}

			if subject != "" {
				if strings.Contains(topic.Subject, subject) {
					hit = true
				}
			}
			if level != "" {
				if strings.Contains(string(topic.Level), level) {
					hit = true
				}
			}

			if description != "" {
				if strings.Contains(topic.Description, description) {
					hit = true
				}
			}
			if hit {
				results = append(results, topic)
			}

		}
	}

	return results, nil

}
