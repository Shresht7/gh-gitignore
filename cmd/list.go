package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-gitignore/api"
)

//	============
//	LIST COMMAND
//	============

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all gitignore templates",
	Long:  `List all the available gitignore templates`,
	Run: func(cmd *cobra.Command, args []string) {

		//	Get list of gitignore templates
		gitignoreList, err := api.ListGitignoreTemplates()
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print the list of gitignore templates
		printList(gitignoreList)

	},
}

// Print the list of licenses to the console
func printList(gitignoreTemplates []string) {
	for _, name := range gitignoreTemplates {
		fmt.Println(name)
	}
}

func init() {
	//	Add list command
	rootCmd.AddCommand(listCmd)
}
