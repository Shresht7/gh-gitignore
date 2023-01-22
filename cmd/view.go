package cmd

import (
	"fmt"

	"github.com/Shresht7/gh-gitignore/api"
	"github.com/Shresht7/gh-gitignore/helpers"
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
	Example: `gh gitignore view Node`,
	Run: func(cmd *cobra.Command, args []string) {

		// Iterate over all provided gitignore template names
		for _, name := range args {

			// Check if the provided gitignore template name is valid
			if !helpers.IsValidTemplate(name) {
				fmt.Printf("Invalid gitignore template name: %s\n", name)
				fmt.Println("\nUse the \"list\" command to view all available gitignore templates.")
				continue
			}

			//	Get the gitignore template
			gitignore, err := api.GetGitignoreTemplate(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			//	Print the gitignore template
			fmt.Println(gitignore.Source)
		}

	},
}

func init() {
	//	Add view command
	rootCmd.AddCommand(viewCmd)
}
