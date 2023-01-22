package cmd

import (
	"fmt"

	"github.com/Shresht7/gh-gitignore/api"
	"github.com/spf13/cobra"
)

//	============
//	VIEW COMMAND
//	============

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:     "view",
	Aliases: []string{"show", "get"},
	Short:   "View a gitignore file",
	Long:    `View a gitignore file in the terminal. Use the "list" command to view all available gitignore templates.`,
	Args:    cobra.ExactArgs(1),
	Example: `gh gitignore view node`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get the gitignore template name
		name := args[0]

		//	Get the gitignore template
		gitignore, err := api.GetGitignoreTemplate(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print the gitignore template
		fmt.Println(gitignore.Source)

	},
}

func init() {
	//	Add view command
	rootCmd.AddCommand(viewCmd)
}
