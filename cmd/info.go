package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a topic",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("shortname").Value.String() == "" {
			fmt.Println("--shortname parameter required")
			return
		}
		topic, err := getManifest(cmd.Flag("shortname").Value.String())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(topic)
	},
}

func init() {
	RootCmd.AddCommand(infoCmd)

	infoCmd.PersistentFlags().String("shortname", "", "Topic shortname")
}
