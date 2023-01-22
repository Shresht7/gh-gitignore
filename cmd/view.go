package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-gitignore/api"
	"github.com/Shresht7/gh-gitignore/helpers"
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

		//	Get the gitignore templates
		templates, err := api.ListGitignoreTemplates()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Iterate over all provided gitignore template names
		for _, name := range args {

			// Check if the provided gitignore template name is invalid
			if !helpers.Contains(templates, name) {
				helpers.HandleInvalidTemplate(name, templates)
				continue // Skip the iteration
			}

			//	Get the gitignore template
			gitignore, err := api.GetGitignoreTemplate(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			//	Print the gitignore template
			fmt.Println(gitignore.Source + "\n")
		}

	},
}

func init() {
	//	Add view command
	rootCmd.AddCommand(viewCmd)
}
