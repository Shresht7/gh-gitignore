package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-gitignore/api"
	"github.com/Shresht7/gh-gitignore/helpers"
)

//	============
//	LIST COMMAND
//	============

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all gitignore templates",
	Long:  `List all the available gitignore templates`,
	Example: helpers.ListExamples([]string{
		"gh gitignore list",
	}),
	Run: func(cmd *cobra.Command, args []string) {

		//	Get list of gitignore templates
		gitignoreList, err := api.ListGitignoreTemplates()
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print the list of gitignore templates
		for _, name := range gitignoreList {
			fmt.Println(name)
		}

	},
}

func init() {
	//	Add list command
	RootCmd.AddCommand(listCmd)
}
